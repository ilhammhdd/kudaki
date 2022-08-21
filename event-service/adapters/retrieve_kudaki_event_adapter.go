package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-event-service/usecases/events"
)

type RetrieveKudakiEvent struct{}

func (rke *RetrieveKudakiEvent) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.RetrieveKudakiEvent
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (rke *RetrieveKudakiEvent) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.KudakiEventRetrieved)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
