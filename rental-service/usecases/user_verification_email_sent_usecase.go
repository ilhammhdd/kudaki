package usecases

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-rental-service/entities/aggregates/rental"
	"github.com/ilhammhdd/kudaki-rental-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-rental-service/usecases/events"
)

type UserVerificationEmailSent struct{}

func (uves *UserVerificationEmailSent) Handle(in proto.Message) *UsecaseHandlerResponse {
	inEvent := in.(*events.UserVerificationEmailSent)

	cart := uves.initEmptyCart(inEvent.User)

	return &UsecaseHandlerResponse{
		Data: cart,
		Ok:   true}
}

func (uves *UserVerificationEmailSent) initEmptyCart(usr *user.User) *rental.Cart {
	return &rental.Cart{
		CreatedAt:  ptypes.TimestampNow(),
		Open:       true,
		TotalItems: 0,
		TotalPrice: 0,
		UserUuid:   usr.Uuid,
		Uuid:       uuid.New().String()}
}
