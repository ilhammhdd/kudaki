package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type DisapproveOwnerOrder struct{}

func (ao *DisapproveOwnerOrder) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.DisapproveOwnerOrder
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (ao *DisapproveOwnerOrder) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.OwnerOrderDisapproved)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
