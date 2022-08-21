package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-event-payment-service/usecases/events"
)

type PaymentRequestDoku struct{}

func (prd *PaymentRequestDoku) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.PaymentRequestDoku
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (prd *PaymentRequestDoku) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.PaymentRequestedDoku)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
