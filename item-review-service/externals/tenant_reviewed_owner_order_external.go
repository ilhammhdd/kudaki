package externals

import (
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-item-review-service/adapters"
	"github.com/ilhammhdd/kudaki-item-review-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases/events"
)

type TenantReviewedOwnerOrder struct{}

func (troo *TenantReviewedOwnerOrder) Work() interface{} {
	usecase := usecases.TenantReviewedOwnerOrder{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}
	ede := EventDrivenExternal{
		PostUsecaseExecutor: troo,
		eventDrivenAdapter:  new(adapters.TenantReviewedOwnerOrder),
		eventDrivenUsecase:  &usecase,
		eventName:           events.OrderServiceEventTopic_TENANT_REVIEWED_OWNER_ORDER.String(),
		inTopics:            []string{events.OrderServiceEventTopic_TENANT_REVIEWED_OWNER_ORDER.String()},
		outTopic:            events.ItemReviewServiceEventTopic_ITEMS_REVIEWED.String()}

	ede.handle()
	return nil
}

func (troo *TenantReviewedOwnerOrder) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.ItemsReviewed)

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	dboInsert := mysql.NewDBOperation(mysql.CommandDB)
	log.Println(out.ItemReviews)
	for i := 0; i < len(out.ItemReviews); i++ {
		_, err := dbo.Command("UPDATE kudaki_store.items SET rating = ?, total_raw_rating = ? WHERE uuid = ?;",
			out.ItemReviews[i].Item.Rating,
			out.ItemReviews[i].Item.TotalRawRating,
			out.ItemReviews[i].Item.Uuid)
		errorkit.ErrorHandled(err)

		_, err = dboInsert.Command("INSERT INTO kudaki_store.item_reviews(uuid,user_uuid,item_uuid,review,rating,created_at) VALUES(?,?,?,?,?,UNIX_TIMESTAMP());",
			out.ItemReviews[i].Uuid,
			out.ItemReviews[i].UserUuid,
			out.ItemReviews[i].Item.Uuid,
			out.ItemReviews[i].Review,
			out.ItemReviews[i].Rating)
		errorkit.ErrorHandled(err)
	}
}
