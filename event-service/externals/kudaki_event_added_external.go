package externals

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-event-service/usecases"
)

type KudakiEventAdded struct{}

func (ae *KudakiEventAdded) Work() interface{} {

	return nil
}

func (ae *KudakiEventAdded) ExecutePostDownstreamUsecase(inEvent proto.Message, usecaseRes *usecases.UsecaseHandlerResponse) {
	if !usecaseRes.Ok {
		return
	}

}
