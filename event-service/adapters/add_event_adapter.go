package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-event-service/usecases/events"
)

type AddKudakiEvent struct{}

func (ae *AddKudakiEvent) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.AddKudakiEvent
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (ae *AddKudakiEvent) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.KudakiEventAdded)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
