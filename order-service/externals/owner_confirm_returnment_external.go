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

type OwnerConfirmReturnment struct{}

func (ocr *OwnerConfirmReturnment) Work() interface{} {
	usecase := usecases.OwnerConfirmReturnment{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: ocr,
		eventDrivenAdapter:  new(adapters.OwnerConfirmReturnment),
		eventDrivenUsecase:  &usecase,
		eventName:           events.OrderServiceCommandTopic_OWNER_CONFIRM_RETURNMENT.String(),
		inTopics:            []string{events.OrderServiceCommandTopic_OWNER_CONFIRM_RETURNMENT.String()},
		outTopic:            events.OrderServiceEventTopic_OWNER_CONFIRMED_RETURNMENT.String()}

	ede.handle()
	return nil
}

func (ocr *OwnerConfirmReturnment) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.OwnerConfirmedReturnment)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("UPDATE kudaki_order.owner_orders SET status=? WHERE uuid=?;",
		order.OrderStatus_DONE.String(), out.OwnerOrder.Uuid)
	errorkit.ErrorHandled(err)
}
