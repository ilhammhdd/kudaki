package externals

import (
	"net/http"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/ilhammhdd/kudaki-order-service/entities/aggregates/order"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-order-service/adapters"
	"github.com/ilhammhdd/kudaki-order-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-order-service/usecases"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type DisapproveOwnerOrder struct{}

func (ao *DisapproveOwnerOrder) Work() interface{} {
	usecase := &usecases.DisapproveOwnerOrder{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: ao,
		eventDrivenAdapter:  new(adapters.DisapproveOwnerOrder),
		eventDrivenUsecase:  usecase,
		eventName:           events.OrderServiceCommandTopic_DISAPPROVE_OWNER_ORDER.String(),
		inTopics:            []string{events.OrderServiceCommandTopic_DISAPPROVE_OWNER_ORDER.String()},
		outTopic:            events.OrderServiceEventTopic_OWNER_ORDER_DISAPPROVED.String()}

	ede.handle()
	return nil
}

func (ao *DisapproveOwnerOrder) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.OwnerOrderDisapproved)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("UPDATE kudaki_order.owner_orders SET status=? WHERE uuid=?;",
		order.OrderStatus_DISAPPROVED.String(), out.OwnerOrder.Uuid)
	errorkit.ErrorHandled(err)
}
