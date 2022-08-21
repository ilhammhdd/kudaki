package usecases

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/entities/aggregates/order"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type OwnerConfirmedReturnment struct {
	DBO DBOperator
}

func (ocr *OwnerConfirmedReturnment) Handle(in proto.Message) *UsecaseHandlerResponse {
	inEvent := in.(*events.OwnerConfirmedReturnment)

	return &UsecaseHandlerResponse{
		Ok: ocr.checkAllOwnerReturnmentConfirmation(inEvent)}
}

func (ocr *OwnerConfirmedReturnment) checkAllOwnerReturnmentConfirmation(inEvent *events.OwnerConfirmedReturnment) (ok bool) {
	rows, err := ocr.DBO.Query("SELECT oo.status FROM kudaki_order.owner_orders oo WHERE oo.order_uuid=?;",
		inEvent.OwnerOrder.Order.Uuid)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var statuses []string
	for rows.Next() {
		var status string
		rows.Scan(&status)
		statuses = append(statuses, status)
	}

	ok = true
	for _, s := range statuses {
		if s != order.OrderStatus_DONE.String() {
			ok = false
		}
	}

	return
}
