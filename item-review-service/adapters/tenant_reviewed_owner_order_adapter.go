package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases/events"
)

type TenantReviewedOwnerOrder struct{}

func (troo *TenantReviewedOwnerOrder) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.TenantReviewedOwnerOrder
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (troo *TenantReviewedOwnerOrder) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.ItemsReviewed)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
