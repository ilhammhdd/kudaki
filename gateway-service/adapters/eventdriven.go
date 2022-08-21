package adapters

import (
	"net/http"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/jwtkit"
	"github.com/ilhammhdd/kudaki-gateway-service/entities/aggregates/user"

	"github.com/ilhammhdd/kudaki-gateway-service/usecases"

	"github.com/golang/protobuf/proto"
)

type EventDrivenHandler interface {
	ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte)
	ParseEventToResponse(in proto.Message) *Response
	initUsecaseHandler(outKey string) usecases.EventDrivenHandler
}

func HandleEventDriven(r *http.Request, edha EventDrivenHandler) *Response {
	outKey, outMsg := edha.ParseRequestToKafkaMessage(r)
	usecaseHandler := edha.initUsecaseHandler(outKey)
	inEvent := usecaseHandler.Handle(outKey, outMsg)
	return edha.ParseEventToResponse(inEvent)
}

func GetUserFromKudakiToken(kudakiToken string) *user.User {
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
