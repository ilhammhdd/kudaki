package externals

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os"

	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/jwtkit"
	"github.com/ilhammhdd/kudaki-user-auth-service/entities/aggregates/user"
	kudakigrpc "github.com/ilhammhdd/kudaki-user-auth-service/externals/grpc"
	"github.com/ilhammhdd/kudaki-user-auth-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases/events"
)

type UsergRPCService struct{}

func (u UsergRPCService) UserAuthentication(ctx context.Context, uar *kudakigrpc.AuthenticateUser) (*kudakigrpc.UserAuthenticated, error) {
	ua := kudakigrpc.UserAuthenticated{
		Uid:         uar.Uid,
		EventStatus: new(events.Status)}

	userFromKudakiToken, err := u.getUserFromKudakiToken(uar.Jwt)
	if err != nil {
		ua.EventStatus.HttpCode = http.StatusBadRequest
		ua.EventStatus.Timestamp = ptypes.TimestampNow()
		ua.EventStatus.Errors = []string{err.Error()}

		return &ua, nil
	}
	if _, ok := u.userExists(userFromKudakiToken); !ok {
		ua.EventStatus.HttpCode = http.StatusNotFound
		ua.EventStatus.Timestamp = ptypes.TimestampNow()
		ua.EventStatus.Errors = []string{"user not found"}

		return &ua, nil
	}

	e := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}

	ok, err := jwtkit.VerifyJWTString(e, jwtkit.JWTString(uar.Jwt))

	if !ok {
		ua.EventStatus.HttpCode = http.StatusUnauthorized
		ua.EventStatus.Timestamp = ptypes.TimestampNow()
		ua.EventStatus.Errors = []string{"jwt not verified"}
	} else {
		ua.EventStatus.HttpCode = http.StatusOK
		ua.EventStatus.Timestamp = ptypes.TimestampNow()
		ua.EventStatus.Errors = []string{"jwt verified"}
	}

	return &ua, err
}

func (u UsergRPCService) getUserFromKudakiToken(kudakiToken string) (*user.User, error) {
	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(kudakiToken))
	if errorkit.ErrorHandled(err) {
		return nil, errors.New("not a valid Kudaki-Token")
	}

	userClaim := jwt.Payload.Claims["user"].(map[string]interface{})
	usr := &user.User{
		AccountType: user.AccountType(user.AccountType_value[userClaim["account_type"].(string)]),
		Email:       userClaim["email"].(string),
		PhoneNumber: userClaim["phone_number"].(string),
		Role:        user.UserRole(user.UserRole_value[userClaim["role"].(string)]),
		Uuid:        userClaim["uuid"].(string),
	}

	return usr, nil
}

func (u UsergRPCService) userExists(usr *user.User) (*user.User, bool) {
	dbo := mysql.NewDBOperation(mysql.QueryDB)
	row, err := dbo.QueryRow("SELECT uuid,email,password,token,role,phone_number,account_type FROM kudaki_user.users WHERE uuid = ?;", usr.Uuid)
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

func (u UsergRPCService) UserAuthorization(ctx context.Context, uar *kudakigrpc.AuthorizeUser) (*kudakigrpc.UserAuthorized, error) {

	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(uar.Jwt))
	errorkit.ErrorHandled(err)

	var argsBuffer bytes.Buffer
	var argsVal []interface{}

	argsVal = append(argsVal, jwt.Payload.Claims["user"].(map[string]interface{})["uuid"])

	for i := 0; i < len(uar.UserRoles); i++ {
		argsVal = append(argsVal, uar.UserRoles[i].String())

		if len(uar.UserRoles) == 1 {
			argsBuffer.WriteString("(?)")
			break
		}

		if i == 0 {
			argsBuffer.WriteString("(?,")
		} else if i == len(uar.UserRoles)-1 {
			argsBuffer.WriteString("?)")
		} else {
			argsBuffer.WriteString("?,")
		}
	}

	query := "SELECT id FROM kudaki_user.users WHERE uuid=? AND role IN" + argsBuffer.String() + ";"

	dbo := mysql.NewDBOperation(mysql.QueryDB)
	row, err := dbo.QueryRow(query, argsVal...)
	errorkit.ErrorHandled(err)

	var totalUserId int

	if scanErr := row.Scan(&totalUserId); scanErr == sql.ErrNoRows {
		grpcErr := "user's role not authorized"

		return &kudakigrpc.UserAuthorized{
			EventStatus: &events.Status{
				Errors:    []string{grpcErr},
				HttpCode:  http.StatusUnauthorized,
				Timestamp: ptypes.TimestampNow()},
			Uid: uar.Uid,
		}, nil
	} else {
		return &kudakigrpc.UserAuthorized{
			EventStatus: &events.Status{
				HttpCode:  http.StatusOK,
				Timestamp: ptypes.TimestampNow()},
			Uid: uar.Uid,
		}, nil
	}
}
