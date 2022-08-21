package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-recommendation-service/usecases/events"
)

type RecommendedGearItemsRetrieved struct{}

func (rgir *RecommendedGearItemsRetrieved) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.RecommendedGearItemsRetrieved
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}
