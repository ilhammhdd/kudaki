package adapters

import (
	"github.com/golang/protobuf/proto"
)

type ReviewItem struct{}

func (ri *ReviewItem) ParseIn(msg []byte) (proto.Message, bool) {
	// var inEvent events.ReviewItem
	// if proto.Unmarshal(msg, &inEvent) == nil {
	// 	return &inEvent, true
	// }
	return nil, false
}

func (ri *ReviewItem) ParseOut(out proto.Message) (key string, message []byte) {
	// outEvent := out.(*events.ItemReviewed)

	// outByte, err := proto.Marshal(out)
	// errorkit.ErrorHandled(err)

	// return outEvent.Uid, outByte
	return "", nil
}
