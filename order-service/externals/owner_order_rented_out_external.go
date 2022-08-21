package externals

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/adapters"
	"github.com/ilhammhdd/kudaki-order-service/entities/aggregates/order"
	"github.com/ilhammhdd/kudaki-order-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-order-service/usecases"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type OwnerOrderRentedOut struct{}

func (ooro *OwnerOrderRentedOut) Work() interface{} {
	usecase := usecases.OwnerOrderRentedOut{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: ooro,
		eventDrivenAdapter:  new(adapters.OwnerOrderRentedOut),
		eventDrivenUsecase:  &usecase,
		eventName:           events.OrderServiceEventTopic_OWNER_ORDER_RENTED_OUT.String(),
		inTopics:            []string{events.OrderServiceEventTopic_OWNER_ORDER_RENTED_OUT.String()},
		outTopic:            events.OrderServiceEventTopic_ORDER_RENTED_OUT.String()}

	ede.handle()
	return nil
}

func (ooro *OwnerOrderRentedOut) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.OrderRentedOut)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("UPDATE kudaki_order.orders o SET o.status = ? where o.uuid = ?;",
		order.OrderStatus_RENTED.String(), out.Order.Uuid)
	errorkit.ErrorHandled(err)
}
