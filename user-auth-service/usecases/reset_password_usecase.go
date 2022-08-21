package usecases

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/mail"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/jwtkit"
	"github.com/ilhammhdd/kudaki-user-auth-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases/events"
	"golang.org/x/crypto/bcrypt"
)

type ResetPasswordSendEmail struct {
	DBO DBOperator
}

func (rpse *ResetPasswordSendEmail) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := rpse.initInOutEvent(in)

	existedUser, ok := rpse.userExists(inEvent)
	if !ok {
		outEvent.EventStatus.Errors = []string{"user with the given email doesn't exists"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	existedProfile, ok := rpse.profileExists(existedUser)
	if !ok {
		outEvent.EventStatus.Errors = []string{"profile that corresponds with user doesn't exists"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	resetToken := rpse.generateResetToken(existedUser, existedProfile)
	err := rpse.sendEmail(resetToken, existedUser, existedProfile)
	errorkit.ErrorHandled(err)

	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.User = existedUser

	return outEvent
}

func (rpse *ResetPasswordSendEmail) initInOutEvent(in proto.Message) (inEvent *events.SendResetPasswordEmail, outEvent *events.ResetPasswordEmailSent) {
	inEvent = in.(*events.SendResetPasswordEmail)

	outEvent = new(events.ResetPasswordEmailSent)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.SendResetPasswordEmail = inEvent
	outEvent.Uid = inEvent.Uid

	return
}

func (rpse *ResetPasswordSendEmail) userExists(inEvent *events.SendResetPasswordEmail) (*user.User, bool) {
	row, err := rpse.DBO.QueryRow("SELECT uuid,email,password,token,role,phone_number,account_type FROM kudaki_user.users WHERE email = ?;", inEvent.Email)
	errorkit.ErrorHandled(err)

	var retrievedUser user.User
	var role string
	var accountType string
	if row.Scan(&retrievedUser.Uuid, &retrievedUser.Email, &retrievedUser.Password, &retrievedUser.Token, &role, &retrievedUser.PhoneNumber, &accountType) == sql.ErrNoRows {
		return nil, false
	}
	retrievedUser.Role = user.UserRole(user.UserRole_value[role])
	retrievedUser.AccountType = user.AccountType(user.AccountType_value[accountType])

	return &retrievedUser, true
}

func (rpse *ResetPasswordSendEmail) profileExists(existedUser *user.User) (*user.Profile, bool) {
	row, err := rpse.DBO.QueryRow("SELECT uuid,full_name,photo FROM kudaki_user.profiles WHERE user_uuid = ?;", existedUser.Uuid)
	errorkit.ErrorHandled(err)

	var existedProfile user.Profile
	if row.Scan(&existedProfile.Uuid, &existedProfile.FullName, &existedProfile.Photo) == sql.ErrNoRows {
		return nil, false
	}
	existedProfile.User = existedUser

	return &existedProfile, true
}

func (rpse *ResetPasswordSendEmail) generateResetToken(usr *user.User, profile *user.Profile) string {
	e := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("RESET_PASSWORD_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("RESET_PASSWORD_PUBLIC_KEY")}

	je := jwtkit.JWTExpiration(86400000)
	resetToken, err := je.GenerateSignedJWTString(
		e,
		"Kudaki.id user resetting password",
		"Kudaki.id user service",
		&map[string]interface{}{
			"user_uuid": usr.Uuid,
			"full_name": profile.FullName,
		})
	errorkit.ErrorHandled(err)

	return string(resetToken)
}

func (rpse *ResetPasswordSendEmail) sendEmail(resetToken string, usr *user.User, profile *user.Profile) error {
	resetPasswordLink := fmt.Sprintf("%s/user/password/reset?reset_token=%s", os.Getenv("GATEWAY_HOST"), resetToken)

	mail := Mail{
		From: mail.Address{
			Address: os.Getenv("MAIL"),
			Name:    "Kudaki.id account management",
		},
		Body:    []byte(resetPasswordLink),
		Subject: "Reset password user",
		To: mail.Address{
			Address: usr.Email,
			Name:    profile.FullName,
		},
	}

	err := mail.SendWithTLS()
	if errorkit.ErrorHandled(err) {
		return err
	}
	return nil
}

type ResetPassword struct {
	DBO DBOperator
}

func (rp *ResetPassword) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := rp.initInOutEvent(in)

	if ok := rp.verifyResetToken(inEvent.ResetToken); !ok {
		outEvent.EventStatus.Errors = []string{"reset token not verified"}
		outEvent.EventStatus.HttpCode = http.StatusUnauthorized
		return outEvent
	}

	valid, err := jwtkit.ValidateExpired(jwtkit.JWTString(inEvent.ResetToken))
	errorkit.ErrorHandled(err)
	if !valid {
		outEvent.EventStatus.Errors = []string{"reset token expired"}
		outEvent.EventStatus.HttpCode = http.StatusUnauthorized
		return outEvent
	}

	existedUser, ok := rp.userExists(inEvent.ResetToken)
	if !ok {
		outEvent.EventStatus.Errors = []string{"user doesn't exists"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	rp.replacePassword(existedUser, inEvent.NewPassword)

	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.User = existedUser

	return outEvent
}

func (rp *ResetPassword) initInOutEvent(in proto.Message) (inEvent *events.ResetPassword, outEvent *events.PasswordReseted) {
	inEvent = in.(*events.ResetPassword)

	outEvent = new(events.PasswordReseted)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.ResetPassword = inEvent
	outEvent.Uid = inEvent.Uid

	return
}

func (rp *ResetPassword) userExists(resetToken string) (*user.User, bool) {
	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(resetToken))
	errorkit.ErrorHandled(err)

	userUUID := jwt.Payload.Claims["user_uuid"].(string)

	row, err := rp.DBO.QueryRow("SELECT email,password,token,role,phone_number,account_type FROM kudaki_user.users WHERE uuid = ?", userUUID)
	errorkit.ErrorHandled(err)

	var existedUser user.User
	var role, accountType string
	if row.Scan(&existedUser.Email, &existedUser.Password, &existedUser.Token, &role, &existedUser.PhoneNumber, &accountType) == sql.ErrNoRows {
		return nil, false
	}
	existedUser.Uuid = userUUID
	existedUser.Role = user.UserRole(user.UserRole_value[role])
	existedUser.AccountType = user.AccountType(user.AccountType_value[accountType])

	return &existedUser, true
}

func (rp *ResetPassword) verifyResetToken(resetToken string) bool {
	ecdsaPair := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("RESET_PASSWORD_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("RESET_PASSWORD_PUBLIC_KEY")}

	ok, err := jwtkit.VerifyJWTString(ecdsaPair, jwtkit.JWTString(resetToken))
	errorkit.ErrorHandled(err)

	return ok
}

func (rp *ResetPassword) replacePassword(usr *user.User, newPassword string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
	errorkit.ErrorHandled(err)

	usr.Password = string(hashedPassword)
}
