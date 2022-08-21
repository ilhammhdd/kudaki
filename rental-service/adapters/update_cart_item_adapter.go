package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-rental-service/usecases/events"
)

type UpdateCartItem struct{}

func (uci *UpdateCartItem) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.UpdateCartItem
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (uci *UpdateCartItem) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.CartItemsUpdated)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
