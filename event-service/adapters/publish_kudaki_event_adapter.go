package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-event-service/usecases/events"
)

type PublishKudakiEvent struct{}

func (pke *PublishKudakiEvent) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.PublishKudakiEvent
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (pke *PublishKudakiEvent) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.KudakiEventPublished)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
