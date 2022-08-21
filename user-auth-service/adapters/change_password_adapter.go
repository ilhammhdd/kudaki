package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases/events"
)

type ChangePassword struct{}

func (cp *ChangePassword) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.ChangePassword
	if err := proto.Unmarshal(msg, &inEvent); err == nil {
		return &inEvent, true
	}
	return nil, false
}

func (cp *ChangePassword) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.PasswordChanged)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
