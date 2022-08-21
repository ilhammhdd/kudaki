package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type TenantReviewsOwnerOrder struct{}

func (tro *TenantReviewsOwnerOrder) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.TenantReviewOwnerOrder
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (tro *TenantReviewsOwnerOrder) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.TenantReviewedOwnerOrder)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
