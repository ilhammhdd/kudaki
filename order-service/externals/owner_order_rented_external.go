package externals

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/adapters"
	"github.com/ilhammhdd/kudaki-order-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-order-service/usecases"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type OwnerOrderRented struct{}

func (oor *OwnerOrderRented) Work() interface{} {
	usecase := &usecases.OwnerOrderRented{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}
	ede := EventDrivenExternal{
		PostUsecaseExecutor: oor,
		eventDrivenAdapter:  new(adapters.OwnerOrderRented),
		eventDrivenUsecase:  usecase,
		eventName:           events.OrderServiceCommandTopic_OWNER_ORDER_RENTED.String(),
		inTopics:            []string{events.OrderServiceCommandTopic_OWNER_ORDER_RENTED.String()},
		outTopic:            events.OrderServiceEventTopic_OWNER_ORDER_RENTED_OUT.String()}

	ede.handle()
	return nil
}

func (oor *OwnerOrderRented) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.OwnerOrderRentedOut)
	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("UPDATE kudaki_order.owner_orders SET status = ? WHERE uuid = ?",
		out.OwnerOrders.OrderStatus.String(), out.OwnerOrders.Uuid)
	errorkit.ErrorHandled(err)
}
