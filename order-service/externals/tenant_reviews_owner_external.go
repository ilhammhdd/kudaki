package externals

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/adapters"
	"github.com/ilhammhdd/kudaki-order-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-order-service/usecases"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type TenantReviewsOwnerOrder struct{}

func (tro *TenantReviewsOwnerOrder) Work() interface{} {
	uc := &usecases.TenantReviewsOwnerOrder{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: tro,
		eventDrivenAdapter:  new(adapters.TenantReviewsOwnerOrder),
		eventDrivenUsecase:  uc,
		eventName:           events.OrderServiceCommandTopic_TENANT_REVIEW_OWNER_ORDER.String(),
		inTopics:            []string{events.OrderServiceCommandTopic_TENANT_REVIEW_OWNER_ORDER.String()},
		outTopic:            events.OrderServiceEventTopic_TENANT_REVIEWED_OWNER_ORDER.String()}

	ede.handle()
	return nil
}

func (tro *TenantReviewsOwnerOrder) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.TenantReviewedOwnerOrder)

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("INSERT INTO kudaki_order.owner_order_reviews(uuid,tenant_uuid,owner_order_uuid,rating,review,created_at) VALUES(?,?,?,?,?,UNIX_TIMESTAMP());",
		out.OwnerOrderReview.Uuid,
		out.OwnerOrderReview.TenantUuid,
		out.OwnerOrderReview.OwnerOrder.Uuid,
		out.OwnerOrderReview.Rating,
		out.OwnerOrderReview.Review)
	errorkit.ErrorHandled(err)
}
