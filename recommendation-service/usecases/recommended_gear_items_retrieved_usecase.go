package usecases

import (
	"github.com/golang/protobuf/proto"
)

type RecommendedGearItemsRetrieved struct {
	DBO DBOperator
}

func (rgir *RecommendedGearItemsRetrieved) Handle(in proto.Message) *UsecaseHandlerResponse {

	return nil
}
