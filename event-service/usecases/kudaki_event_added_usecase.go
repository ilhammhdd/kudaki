package usecases

import (
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-event-service/usecases/events"
)

type KudakiEventAdded struct {
	DBO DBOperator
}

func (ae *KudakiEventAdded) Handle(in proto.Message) *UsecaseHandlerResponse {
	inEvent := in.(*events.KudakiEventAdded)
	log.Println(inEvent)

	return nil 
}
