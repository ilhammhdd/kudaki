package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases/events"
)

type ResetPasswordSendEmail struct{}

func (rpse *ResetPasswordSendEmail) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.SendResetPasswordEmail
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}

	return nil, false
}

func (rpse *ResetPasswordSendEmail) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.ResetPasswordEmailSent)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type ResetPassword struct{}

func (rp *ResetPassword) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.ResetPassword

	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (rp *ResetPassword) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.PasswordReseted)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
