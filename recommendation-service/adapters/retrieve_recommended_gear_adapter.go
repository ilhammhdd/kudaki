package adapters

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-recommendation-service/usecases"
	"github.com/ilhammhdd/kudaki-recommendation-service/usecases/events"
)

type RetrieveRecommendationGear struct {
	Result RetrieveRecommendationGearResult
}

func (rrg *RetrieveRecommendationGear) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.RetrieveRecommendedGears
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (rrg *RetrieveRecommendationGear) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.RecommendedGearsRetrieved)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type RetrieveRecommendationGearResult struct {
	RecommendedGears []*usecases.RecommendedGearTemp `json:"recommended_gears"`
}

func (rrg *RetrieveRecommendationGear) SetResultSources(i ...interface{}) usecases.ResultSchemer {
	rrg.Result = RetrieveRecommendationGearResult{
		RecommendedGears: i[0].([]*usecases.RecommendedGearTemp)}

	return rrg
}

func (rrg *RetrieveRecommendationGear) ParseToResult() []byte {
	resutltJSON, err := json.Marshal(rrg.Result)
	errorkit.ErrorHandled(err)

	return resutltJSON
}
