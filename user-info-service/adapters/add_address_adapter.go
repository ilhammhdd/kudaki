package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases/events"
)

type AddAddress struct{}

func (ad *AddAddress) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.AddAddress

	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (ad *AddAddress) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.AddressAdded)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
