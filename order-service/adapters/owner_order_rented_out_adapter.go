package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type OwnerOrderRentedOut struct{}

func (ooro *OwnerOrderRentedOut) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.OwnerOrderRentedOut
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (ooro *OwnerOrderRentedOut) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.OrderRentedOut)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
