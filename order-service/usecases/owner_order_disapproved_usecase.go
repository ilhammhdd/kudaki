package usecases

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/entities/aggregates/order"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type OwnerOrderDisapproved struct {
	DBO DBOperator
}

func (ood *OwnerOrderDisapproved) Handle(in proto.Message) *UsecaseHandlerResponse {
	inEvent := in.(*events.OwnerOrderDisapproved)

	return &UsecaseHandlerResponse{
		Ok: ood.checkAllOwnerOrderDisapproval(inEvent)}
}

func (ood *OwnerOrderDisapproved) checkAllOwnerOrderDisapproval(inEvent *events.OwnerOrderDisapproved) (ok bool) {
	rows, err := ood.DBO.Query("SELECT oo.status FROM kudaki_order.owner_orders oo WHERE oo.order_uuid=?;", inEvent.OwnerOrder.Order.Uuid)
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
		if s != order.OrderStatus_DISAPPROVED.String() {
			ok = false
		}
	}

	return
}
