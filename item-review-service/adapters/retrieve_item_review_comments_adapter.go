package adapters

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases/events"
)

type RetrieveItemReviewComment struct {
	Result RetrieveItemReviewCommentResultSchema
}

func (rirc *RetrieveItemReviewComment) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.RetrieveItemReviewComments
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (rirc *RetrieveItemReviewComment) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.ItemReviewCommentsRetrieved)

	outByte, err := proto.Marshal(out)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type RetrieveItemReviewCommentResultSchema struct {
	ItemReview *usecases.ItemReviewTemp      `json:"item_review"`
	Comments   []*usecases.ReviewCommentTemp `json:"comments"`
}

func (rirc *RetrieveItemReviewComment) SetResultSources(i ...interface{}) usecases.ResultSchemer {
	rirc.Result = RetrieveItemReviewCommentResultSchema{
		ItemReview: i[0].(*usecases.ItemReviewTemp),
		Comments:   i[1].([]*usecases.ReviewCommentTemp)}

	return rirc
}
func (rirc *RetrieveItemReviewComment) ParseToResult() []byte {
	resultJSON, err := json.Marshal(rirc.Result)
	errorkit.ErrorHandled(err)

	return resultJSON
}
