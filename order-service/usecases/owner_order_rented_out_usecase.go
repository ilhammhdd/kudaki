package usecases

import (
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/entities/aggregates/order"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type OwnerOrderRentedOut struct {
	DBO DBOperator
}

func (ooro *OwnerOrderRentedOut) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := ooro.initInOutEvent(in)

	if ooro.checkAllOwnerOrderStatusRented(inEvent.OwnerOrders) {
		outEvent.Order = ooro.retrieveOrder(inEvent.OwnerOrders.Order.Uuid)
		outEvent.EventStatus.HttpCode = http.StatusOK
		return outEvent
	}

	outEvent.EventStatus.HttpCode = http.StatusNoContent
	return outEvent
}

func (ooro *OwnerOrderRentedOut) initInOutEvent(in proto.Message) (inEvent *events.OwnerOrderRentedOut, outEvent *events.OrderRentedOut) {
	inEvent = in.(*events.OwnerOrderRentedOut)

	outEvent = new(events.OrderRentedOut)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Owner = inEvent.Owner
	outEvent.Uid = inEvent.Uid

	return
}

func (ooro *OwnerOrderRentedOut) checkAllOwnerOrderStatusRented(ownerOrder *order.OwnerOrder) (ok bool) {
	rows, err := ooro.DBO.Query("SELECT oo.status FROM kudaki_order.owner_orders oo WHERE oo.order_uuid = ?;", ownerOrder.Order.Uuid)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	ok = true
	for rows.Next() {
		var stat string
		rows.Scan(&stat)

		if stat != order.OrderStatus_RENTED.String() {
			ok = false
			return
		}
	}

	return
}

func (ooro *OwnerOrderRentedOut) retrieveOrder(orderUuid string) *order.Order {
	row, err := ooro.DBO.QueryRow("SELECT id,uuid,cart_uuid,tenant_uuid,order_num,status,shipment_fee,delivered,created_at WHERE uuid = ?", orderUuid)
	errorkit.ErrorHandled(err)

	var ord order.Order
	var stat string
	var createdAt int64
	row.Scan(
		&ord.Id,
		&ord.Uuid,
		&ord.CartUuid,
		&ord.TenantUuid,
		&ord.OrderNum,
		&stat,
		&ord.ShipmentFee,
		&ord.Delivered,
		&createdAt)

	ord.Status = order.OrderStatus(order.OrderStatus_value[stat])
	createdAtProto, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
	errorkit.ErrorHandled(err)
	ord.CreatedAt = createdAtProto

	return &ord
}
