package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-rental-service/usecases/events"
)

type AddCartItem struct{}

func (aci *AddCartItem) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.AddCartItem
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (aci *AddCartItem) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.CartItemAdded)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
