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

type OwnerOrderApproved struct{}

func (ooa *OwnerOrderApproved) Work() interface{} {
	usecase := usecases.OwnerOrderApproved{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	edde := EventDrivenDownstreamExternal{
		PostUsecaseExecutor: ooa,
		eventDrivenAdapter:  new(adapters.OwnerOrderApproved),
		eventDrivenUsecase:  &usecase,
		eventName:           events.OrderServiceEventTopic_OWNER_ORDER_APPROVED.String(),
		inTopics:            []string{events.OrderServiceEventTopic_OWNER_ORDER_APPROVED.String()}}

	edde.handle()
	return nil
}

func (ooa *OwnerOrderApproved) ExecutePostDownstreamUsecase(inEvent proto.Message, usecaseRes *usecases.UsecaseHandlerResponse) {
	in := inEvent.(*events.OwnerOrderApproved)

	if usecaseRes.Ok {
		dbo := mysql.NewDBOperation(mysql.CommandDB)
		_, err := dbo.Command("UPDATE kudaki_order.orders SET status = ? WHERE uuid = ?;",
			order.OrderStatus_APPROVED.String(), in.OwnerOrder.Order.Uuid)
		errorkit.ErrorHandled(err)
	}
}
