package externals

import (
	"github.com/ilhammhdd/kudaki-item-review-service/adapters"
	"github.com/ilhammhdd/kudaki-item-review-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases/events"
)

type RetrieveItemReviewComment struct{}

func (rirc *RetrieveItemReviewComment) Work() interface{} {
	adapter := adapters.RetrieveItemReviewComment{}
	usecase := usecases.RetrieveItemReviewComment{
		DBO:           mysql.NewDBOperation(mysql.QueryDB),
		ResultSchemer: &adapter}

	ede := EventDrivenExternal{
		eventDrivenAdapter: &adapter,
		eventDrivenUsecase: &usecase,
		eventName:          events.ItemReviewServiceCommandTopic_RETRIEVE_ITEM_REVIEW_COMMENTS.String(),
		inTopics:           []string{events.ItemReviewServiceCommandTopic_RETRIEVE_ITEM_REVIEW_COMMENTS.String()},
		outTopic:           events.ItemReviewServiceEventTopic_ITEM_REVIEW_COMMENTS_RETRIEVED.String()}

	ede.handle()
	return nil
}
