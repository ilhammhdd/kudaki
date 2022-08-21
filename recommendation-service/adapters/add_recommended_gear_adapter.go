package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-recommendation-service/usecases/events"
)

type AddRecommendationGear struct{}

func (arg *AddRecommendationGear) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.AddRecommendedGear
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (arg *AddRecommendationGear) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.RecommendedGearAdded)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
