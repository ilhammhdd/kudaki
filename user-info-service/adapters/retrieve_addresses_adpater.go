package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases/events"
)

type RetrieveAddresses struct{}

func (ra *RetrieveAddresses) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.RetrieveAddresses
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (ra *RetrieveAddresses) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.AddressesRetrieved)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
