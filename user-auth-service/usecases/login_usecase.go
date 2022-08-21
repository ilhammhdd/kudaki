package usecases

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/jwtkit"
	"github.com/ilhammhdd/kudaki-user-auth-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases/events"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	DBO DBOperator
}

func (l *Login) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := l.initInOutEvent(in)

	l.initInOutEvent(in)
	existedUser, ok := l.userExists(inEvent)
	if !ok {
		outEvent.EventStatus.Errors = []string{"user with the given email doesn't exists"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	if !l.userVerified(existedUser.Uuid) {
		outEvent.EventStatus.Errors = []string{"user wasn't verified"}
		outEvent.EventStatus.HttpCode = http.StatusUnauthorized
		return outEvent
	}

	existedProfile, ok := l.profileExists(inEvent, existedUser)
	if !ok {
		outEvent.EventStatus.Errors = []string{"profile corresponds to user doesn't exists"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(inEvent.Password)); err != nil {
		outEvent.EventStatus.Errors = []string{"wrong password"}
		outEvent.EventStatus.HttpCode = http.StatusUnauthorized
		return outEvent
	}

	existedUser.Token = l.generateAuthToken(existedUser, existedProfile)
	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.User = existedUser

	return outEvent
}

func (l *Login) initInOutEvent(in proto.Message) (inEvent *events.Login, outEvent *events.Loggedin) {
	inEvent = in.(*events.Login)

	outEvent = new(events.Loggedin)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Login = inEvent
	outEvent.Uid = inEvent.Uid

	return
}

func (l *Login) userExists(inEvent *events.Login) (*user.User, bool) {
	row, err := l.DBO.QueryRow("SELECT uuid,email,password,token,role,phone_number,account_type FROM kudaki_user.users WHERE email = ?;", inEvent.Email)
	errorkit.ErrorHandled(err)

	var existedUser user.User
	var role, accountType string
	if row.Scan(&existedUser.Uuid, &existedUser.Email, &existedUser.Password, &existedUser.Token, &role, &existedUser.PhoneNumber, &accountType) == sql.ErrNoRows {
		return nil, false
	}
	existedUser.Role = user.UserRole(user.UserRole_value[role])
	existedUser.AccountType = user.AccountType(user.AccountType_value[accountType])

	return &existedUser, true
}

func (l *Login) profileExists(inEvent *events.Login, existedUser *user.User) (*user.Profile, bool) {
	row, err := l.DBO.QueryRow("SELECT uuid,full_name,photo FROM kudaki_user.profiles WHERE user_uuid = ?;", existedUser.Uuid)
	errorkit.ErrorHandled(err)

	var existedProfile user.Profile
	if row.Scan(&existedProfile.Uuid, &existedProfile.FullName, &existedProfile.Photo) == sql.ErrNoRows {
		return nil, false
	}
	existedProfile.User = existedUser

	return &existedProfile, true
}

func (l *Login) userVerified(userUUID string) bool {
	row, err := l.DBO.QueryRow("SELECT id FROM kudaki_user.unverified_users WHERE user_uuid = ?;", userUUID)
	errorkit.ErrorHandled(err)

	var unverifiedUserID uint64
	if row.Scan(&unverifiedUserID) == sql.ErrNoRows {
		return true
	}
	return false
}

func (l *Login) generateAuthToken(usr *user.User, profile *user.Profile) string {
	authECDSA := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}

	jwtString, err := jwtkit.JWTExpiration(5.256e+9).GenerateSignedJWTString(
		authECDSA,
		"verified Kudaki.id user",
		"Kudaki.id user service",
		&map[string]interface{}{
			"user": map[string]interface{}{
				"account_type": usr.AccountType.String(),
				"email":        usr.Email,
				"phone_number": usr.PhoneNumber,
				"role":         usr.Role.String(),
				"uuid":         usr.Uuid,
			},
			"profile": map[string]interface{}{
				"user_uuid": profile.User.Uuid,
				"uuid":      profile.Uuid,
				"full_name": profile.FullName,
				"photo":     profile.Photo,
			},
		})
	errorkit.ErrorHandled(err)

	return string(jwtString)
}
