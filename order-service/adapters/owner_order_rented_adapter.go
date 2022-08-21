package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type OwnerOrderRented struct{}

func (oor *OwnerOrderRented) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.OwnerOrderRented
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (oor *OwnerOrderRented) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.OwnerOrderRentedOut)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
