package externals

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/adapters"
	"github.com/ilhammhdd/kudaki-order-service/entities/aggregates/order"
	"github.com/ilhammhdd/kudaki-order-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-order-service/usecases"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type OwnerOrderDisapproved struct{}

func (ood *OwnerOrderDisapproved) Work() interface{} {
	usecase := &usecases.OwnerOrderDisapproved{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	edde := EventDrivenDownstreamExternal{
		PostUsecaseExecutor: ood,
		eventDrivenAdapter:  new(adapters.OwnerOrderDisapproved),
		eventDrivenUsecase:  usecase,
		eventName:           events.OrderServiceEventTopic_OWNER_ORDER_DISAPPROVED.String(),
		inTopics:            []string{events.OrderServiceEventTopic_OWNER_ORDER_DISAPPROVED.String()}}

	edde.handle()
	return nil
}

func (ood *OwnerOrderDisapproved) ExecutePostDownstreamUsecase(inEvent proto.Message, usecaseRes *usecases.UsecaseHandlerResponse) {
	in := inEvent.(*events.OwnerOrderDisapproved)

	if usecaseRes.Ok {
		dbo := mysql.NewDBOperation(mysql.CommandDB)
		_, err := dbo.Command("UPDATE kudaki_order.orders SET status = ? WHERE uuid = ?;",
			order.OrderStatus_DISAPPROVED.String(), in.OwnerOrder.Order.Uuid)
		errorkit.ErrorHandled(err)
	}
}
