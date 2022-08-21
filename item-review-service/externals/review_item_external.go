package externals

import (
	"github.com/golang/protobuf/proto"
)

type ReviewItem struct{}

func (ri *ReviewItem) Work() interface{} {
	// usecase := &usecases.ReviewItem{DBO: mysql.NewDBOperation(mysql.QueryDB)}

	// ede := EventDrivenExternal{
	// 	PostUsecaseExecutor: ri,
	// 	eventDrivenAdapter:  new(adapters.ReviewItem),
	// 	eventDrivenUsecase:  usecase,
	// 	eventName:           events.ItemReviewServiceCommandTopic_REVIEW_ITEM.String(),
	// 	inTopics:            []string{events.ItemReviewServiceCommandTopic_REVIEW_ITEM.String()},
	// 	outTopic:            events.ItemReviewServiceEventTopic_ITEM_REVIEWED.String()}

	// ede.handle()
	return nil
}

func (ri *ReviewItem) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	// out := outEvent.(*events.ItemReviewed)

	// if out.EventStatus.HttpCode != http.StatusOK {
	// 	return
	// }

	// dboItemReview := mysql.NewDBOperation(mysql.CommandDB)

	// _, err := dboItemReview.Command("INSERT INTO kudaki_store.item_reviews(uuid,user_uuid,item_uuid,review,rating,created_at) VALUES (?,?,?,?,?,UNIX_TIMESTAMP());",
	// 	out.ItemReview.Uuid, out.ItemReview.UserUuid, out.ItemReview.Item.Uuid, out.ItemReview.Review, out.ItemReview.Rating)
	// errorkit.ErrorHandled(err)

	// dboItem := mysql.NewDBOperation(mysql.CommandDB)
	// _, err = dboItem.Command("UPDATE kudaki_store.items SET rating=?,total_raw_rating=? WHERE uuid=?;", out.ItemReview.Item.Rating, out.ItemReview.Item.TotalRawRating, out.ItemReview.Item.Uuid)
	// errorkit.ErrorHandled(err)
}
