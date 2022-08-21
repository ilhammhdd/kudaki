package usecases

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/jwtkit"
	"github.com/ilhammhdd/kudaki-user-auth-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases/events"
)

type VerifyUser struct {
	DBO DBOperator
}

func (vu *VerifyUser) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := vu.initInOutEvent(in)

	ecdsaPair := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}

	if !vu.verifyToken(ecdsaPair, inEvent) {
		outEvent.EventStatus.Errors = []string{"token not verified"}
		outEvent.EventStatus.HttpCode = http.StatusUnauthorized
		return outEvent
	}

	if !vu.validateToken(inEvent) {
		outEvent.EventStatus.Errors = []string{"token not valid"}
		outEvent.EventStatus.HttpCode = http.StatusUnauthorized
		return outEvent
	}

	user := vu.retrieveUser(inEvent)
	if user == nil {
		outEvent.EventStatus.Errors = []string{"user not found"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.User = user

	return outEvent
}

func (vu *VerifyUser) initInOutEvent(in proto.Message) (inEvent *events.VerifyUser, outEvent *events.UserVerified) {
	inEvent = in.(*events.VerifyUser)

	outEvent = new(events.UserVerified)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Uid = inEvent.Uid
	outEvent.VerifyUser = inEvent
	return
}

func (vu *VerifyUser) verifyToken(ecdsaPair *jwtkit.ECDSA, inEvent *events.VerifyUser) bool {
	verified, err := jwtkit.VerifyJWTString(ecdsaPair, jwtkit.JWTString(inEvent.VerifyToken))
	errorkit.ErrorHandled(err)
	return verified
}

func (vu *VerifyUser) validateToken(inEvent *events.VerifyUser) bool {
	validated, err := jwtkit.ValidateExpired(jwtkit.JWTString(inEvent.VerifyToken))
	errorkit.ErrorHandled(err)
	return validated
}

func (vu *VerifyUser) retrieveUser(inEvent *events.VerifyUser) *user.User {
	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(inEvent.VerifyToken))
	errorkit.ErrorHandled(err)

	userUUID := jwt.Payload.Claims["user_uuid"].(string)
	log.Println(jwt.Payload.Claims)
	row, err := vu.DBO.QueryRow("SELECT uuid,email,password,token,role,phone_number,account_type FROM kudaki_user.users WHERE uuid=?;",
		userUUID)
	errorkit.ErrorHandled(err)

	var usr user.User
	var role string
	var accountType string
	err = row.Scan(&usr.Uuid, &usr.Email, &usr.Password, &usr.Token, &role, &usr.PhoneNumber, &accountType)
	if err == sql.ErrNoRows {
		return nil
	}

	usr.Role = user.UserRole(user.UserRole_value[role])
	usr.AccountType = user.AccountType(user.AccountType_value[accountType])

	return &usr
}
