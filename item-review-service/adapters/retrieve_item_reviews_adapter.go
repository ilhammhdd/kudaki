package adapters

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases/events"
)

type RetrieveItemReviews struct {
	Result RetrieveItemReviewsResultSchema
}

func (rir *RetrieveItemReviews) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.RetrieveItemReviews
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}
func (rir *RetrieveItemReviews) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.ItemReviewsRetrieved)

	outByte, err := proto.Marshal(out)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type RetrieveItemReviewsResultSchema struct {
	ItemReviews []*usecases.ItemReviewTemp `json:"item_reviews"`
}

func (rir *RetrieveItemReviews) SetResultSources(i ...interface{}) usecases.ResultSchemer {
	rir.Result = RetrieveItemReviewsResultSchema{
		ItemReviews: i[0].([]*usecases.ItemReviewTemp)}

	return rir
}

func (rir *RetrieveItemReviews) ParseToResult() []byte {
	resultJSON, err := json.Marshal(rir.Result)
	errorkit.ErrorHandled(err)

	return resultJSON
}
