package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type CheckOut struct{}

func (co *CheckOut) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.CheckOut
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (co *CheckOut) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.CheckedOut)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
