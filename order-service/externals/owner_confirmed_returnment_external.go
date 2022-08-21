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

type OwnerConfirmedReturnment struct{}

func (ocr *OwnerConfirmedReturnment) Work() interface{} {
	usecase := usecases.OwnerConfirmedReturnment{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	edde := EventDrivenDownstreamExternal{
		PostUsecaseExecutor: ocr,
		eventDrivenAdapter:  new(adapters.OwnerConfirmedReturnment),
		eventDrivenUsecase:  &usecase,
		eventName:           events.OrderServiceEventTopic_OWNER_CONFIRMED_RETURNMENT.String(),
		inTopics:            []string{events.OrderServiceEventTopic_OWNER_CONFIRMED_RETURNMENT.String()}}
	edde.handle()
	return nil
}

func (ocr *OwnerConfirmedReturnment) ExecutePostDownstreamUsecase(inEvent proto.Message, usecaseRes *usecases.UsecaseHandlerResponse) {
	in := inEvent.(*events.OwnerConfirmedReturnment)

	if usecaseRes.Ok {
		dbo := mysql.NewDBOperation(mysql.CommandDB)
		_, err := dbo.Command("UPDATE kudaki_order.orders SET status=? WHERE uuid=?;",
			order.OrderStatus_DONE.String(), in.OwnerOrder.Order.Uuid)
		errorkit.ErrorHandled(err)
	}
}
