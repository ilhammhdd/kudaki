package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases/events"
)

type CommentReviewItem struct{}

func (cri *CommentReviewItem) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.CommentItemReview
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (cri *CommentReviewItem) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.ItemReviewCommented)

	outByte, err := proto.Marshal(out)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
