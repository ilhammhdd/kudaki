package externals

import (
	"github.com/ilhammhdd/kudaki-item-review-service/adapters"
	"github.com/ilhammhdd/kudaki-item-review-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases/events"
)

type RetrieveItemReviews struct{}

func (rir *RetrieveItemReviews) Work() interface{} {
	adapter := adapters.RetrieveItemReviews{}
	usecase := usecases.RetrieveItemReviews{
		DBO:           mysql.NewDBOperation(mysql.QueryDB),
		ResultSchemer: &adapter}

	ede := EventDrivenExternal{
		eventDrivenAdapter: &adapter,
		eventDrivenUsecase: &usecase,
		eventName:          events.ItemReviewServiceCommandTopic_RETRIEVE_ITEM_REVIEWS.String(),
		inTopics:           []string{events.ItemReviewServiceCommandTopic_RETRIEVE_ITEM_REVIEWS.String()},
		outTopic:           events.ItemReviewServiceEventTopic_ITEM_REVIEWS_RETRIEVED.String()}

	ede.handle()
	return nil
}
