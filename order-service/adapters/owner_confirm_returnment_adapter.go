package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type OwnerConfirmReturnment struct{}

func (ocr *OwnerConfirmReturnment) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.OwnerConfirmReturnment
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (ocr *OwnerConfirmReturnment) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.OwnerConfirmedReturnment)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
