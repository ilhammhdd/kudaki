package adapters

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-recommendation-service/entities/aggregates/mountain"
	"github.com/ilhammhdd/kudaki-recommendation-service/usecases"
	"github.com/ilhammhdd/kudaki-recommendation-service/usecases/events"
)

type RetrieveRecommendedGearItems struct {
	Result RetrieveRecommendedGearItemsResult
}

func (rrgi *RetrieveRecommendedGearItems) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.RetrieveRecommendedGearItems
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (rrgi *RetrieveRecommendedGearItems) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.RecommendedGearItemsRetrieved)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type RetrieveRecommendedGearItemsResult struct {
	RecommendedGear      *usecases.RecommendedGearTemp   `json:"recommended_gear"`
	RecommendedGearItems []*mountain.RecommendedGearItem `json:"recommended_gear_items"`
}

func (rrgi *RetrieveRecommendedGearItems) SetResultSources(i ...interface{}) usecases.ResultSchemer {
	rrgi.Result = RetrieveRecommendedGearItemsResult{
		RecommendedGear:      i[0].(*usecases.RecommendedGearTemp),
		RecommendedGearItems: i[1].([]*mountain.RecommendedGearItem)}

	return rrgi
}

func (rrgi *RetrieveRecommendedGearItems) ParseToResult() []byte {
	resultJSON, err := json.Marshal(rrgi.Result)
	errorkit.ErrorHandled(err)

	return resultJSON
}
