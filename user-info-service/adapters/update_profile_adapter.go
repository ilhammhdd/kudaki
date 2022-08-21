package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases/events"
)

type UpdateProfile struct{}

func (up *UpdateProfile) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.UpdateProfile
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (up *UpdateProfile) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.ProfileUpdated)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
