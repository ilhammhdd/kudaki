package usecases

import (
	"database/sql"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-order-service/entities/aggregates/order"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type OwnerOrderApproved struct {
	DBO DBOperator
}

func (ooa *OwnerOrderApproved) Handle(in proto.Message) *UsecaseHandlerResponse {
	inEvent := in.(*events.OwnerOrderApproved)

	return &UsecaseHandlerResponse{
		Ok: ooa.checkAllOwnerOrderApproval(inEvent)}
}

func (ooa *OwnerOrderApproved) checkAllOwnerOrderApproval(inEvent *events.OwnerOrderApproved) (ok bool) {
	rows, err := ooa.DBO.Query("SELECT oo.status FROM kudaki_order.owner_orders oo WHERE oo.order_uuid=?;", inEvent.OwnerOrder.Order.Uuid)
	if err == sql.ErrNoRows {
		return false
	}
	defer rows.Close()

	var statuses []string
	for rows.Next() {
		var status string
		rows.Scan(&status)
		statuses = append(statuses, status)
	}

	ok = true
	for _, s := range statuses {
		if s != order.OrderStatus_APPROVED.String() {
			ok = false
		}
	}

	return
}
