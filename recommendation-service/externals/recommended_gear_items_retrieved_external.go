package externals

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-recommendation-service/usecases"
)

type RecommendedGearItemsRetrieved struct{}

func (rgir *RecommendedGearItemsRetrieved) Work() interface{} {

	return nil
}

func (rgir *RecommendedGearItemsRetrieved) ExecutePostDownstreamUsecase(inEvent proto.Message, usecaseRes *usecases.UsecaseHandlerResponse) {

}
