package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-rental-service/usecases/events"
)

type UserVerificationEmailSent struct{}

func (uves *UserVerificationEmailSent) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.UserVerificationEmailSent
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}
