package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-event-service/usecases/events"
)

type DeleteKudakiEvent struct{}

func (dke *DeleteKudakiEvent) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.DeleteKudakiEvent
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (dke *DeleteKudakiEvent) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.KudakiEventDeleted)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
