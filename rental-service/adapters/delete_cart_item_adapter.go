package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-rental-service/usecases/events"
)

type DeleteCartItem struct{}

func (dci *DeleteCartItem) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.DeleteCartItem
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (dci *DeleteCartItem) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.CartItemDeleted)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
