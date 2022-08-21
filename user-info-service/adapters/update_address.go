package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases/events"
)

type UpdateAddress struct{}

func (ua *UpdateAddress) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.UpdateAddress

	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (ua *UpdateAddress) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.AddressUpdated)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
