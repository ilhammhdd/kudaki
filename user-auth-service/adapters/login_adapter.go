package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases/events"
)

type Login struct{}

func (l *Login) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.Login

	if err := proto.Unmarshal(msg, &inEvent); err == nil {
		return &inEvent, true
	}

	return nil, false
}

func (l *Login) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.Loggedin)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
