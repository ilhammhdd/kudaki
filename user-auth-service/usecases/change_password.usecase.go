package usecases

import (
	"database/sql"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/jwtkit"
	"github.com/ilhammhdd/kudaki-user-auth-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases/events"
	"golang.org/x/crypto/bcrypt"
)

type ChangePassword struct {
	DBO DBOperator
}

func (cp *ChangePassword) initInOutEvent(in proto.Message) (inEvent *events.ChangePassword, outEvent *events.PasswordChanged) {
	inEvent = in.(*events.ChangePassword)

	outEvent = new(events.PasswordChanged)
	outEvent.ChangePassword = inEvent
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Uid = inEvent.Uid

	return
}

func (cp *ChangePassword) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := cp.initInOutEvent(in)

	userFromJWT := cp.getUserFromKudakiToken(inEvent.KudakiToken)
	existedUser, ok := cp.userExists(userFromJWT)

	if !ok {
		outEvent.EventStatus.Errors = []string{"user doesn't exists"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	if bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(inEvent.OldPassword)) != nil {
		outEvent.EventStatus.Errors = []string{"wrong old password"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(inEvent.NewPassword), bcrypt.MinCost)
	errorkit.ErrorHandled(err)
	existedUser.Password = string(newPassword)

	outEvent.User = existedUser
	outEvent.EventStatus.HttpCode = http.StatusOK
	return outEvent
}

func (cp *ChangePassword) userExists(usr *user.User) (*user.User, bool) {
	row, err := cp.DBO.QueryRow("SELECT uuid,email,password,token,role,phone_number,account_type FROM kudaki_user.users WHERE uuid = ?;", usr.Uuid)
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

func (cp *ChangePassword) getUserFromKudakiToken(kudakiToken string) *user.User {
	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(kudakiToken))
	errorkit.ErrorHandled(err)

	userClaim := jwt.Payload.Claims["user"].(map[string]interface{})
	usr := &user.User{
		AccountType: user.AccountType(user.AccountType_value[userClaim["account_type"].(string)]),
		Email:       userClaim["email"].(string),
		PhoneNumber: userClaim["phone_number"].(string),
		Role:        user.UserRole(user.UserRole_value[userClaim["role"].(string)]),
		Uuid:        userClaim["uuid"].(string),
	}

	return usr
}
