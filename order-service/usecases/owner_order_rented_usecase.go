package usecases

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/golang/protobuf/ptypes"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/entities/aggregates/order"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type OwnerOrderRented struct {
	DBO DBOperator
}

func (oor *OwnerOrderRented) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := oor.initInOutEvent(in)

	ownerOrder := oor.retrieveOwnerOrder(inEvent.OwnerOrderUuid, outEvent.Owner.Uuid)
	if ownerOrder == nil {
		outEvent.EventStatus.Errors = []string{"owner order with the given uuid not found, or the owner doesn't own the order, or the owner order status not APPROVED yet"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	oor.changeOwnerOrderStatus(ownerOrder)

	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.OwnerOrders = ownerOrder
	return outEvent
}

func (oor *OwnerOrderRented) initInOutEvent(in proto.Message) (inEvent *events.OwnerOrderRented, outEvent *events.OwnerOrderRentedOut) {
	inEvent = in.(*events.OwnerOrderRented)

	outEvent = new(events.OwnerOrderRentedOut)
	outEvent.EventStatus = new(events.Status)
	outEvent.Owner = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.Uid = inEvent.Uid

	return
}

func (oor *OwnerOrderRented) retrieveOwnerOrder(ownerOrderUuid string, ownerUuid string) *order.OwnerOrder {
	row, err := oor.DBO.QueryRow("SELECT oo.id, oo.uuid, oo.order_uuid, oo.owner_uuid, oo.total_price, oo.total_quantity, oo.status, oo.created_at, o.id, o.cart_uuid, o.tenant_uuid, o.order_num, o.status, o.shipment_fee, o.delivered, o.created_at FROM kudaki_order.owner_orders oo JOIN kudaki_order.orders o ON oo.order_uuid = o.uuid WHERE oo.uuid = ? AND oo.owner_uuid = ? AND oo.status = ?;",
		ownerOrderUuid, ownerUuid, order.OrderStatus_APPROVED.String())
	errorkit.ErrorHandled(err)

	var ownerOrder order.OwnerOrder
	ownerOrder.Order = new(order.Order)
	var ooStatus string
	var ooCreatedAt int64
	var oStatus string
	var oCreatedAt int64
	var oDelivered int32

	err = row.Scan(
		&ownerOrder.Id,
		&ownerOrder.Uuid,
		&ownerOrder.Order.Uuid,
		&ownerOrder.OwnerUuid,
		&ownerOrder.TotalPrice,
		&ownerOrder.TotalQuantity,
		&ooStatus,
		&ooCreatedAt,
		&ownerOrder.Order.Id,
		&ownerOrder.Order.CartUuid,
		&ownerOrder.Order.TenantUuid,
		&ownerOrder.Order.OrderNum,
		&oStatus,
		&ownerOrder.Order.ShipmentFee,
		&oDelivered,
		&oCreatedAt)
	if err == sql.ErrNoRows {
		return nil
	}

	ownerOrder.OrderStatus = order.OrderStatus(order.OrderStatus_value[ooStatus])
	ooCreatedAtProto, err := ptypes.TimestampProto(time.Unix(ooCreatedAt, 0))
	errorkit.ErrorHandled(err)
	ownerOrder.CreatedAt = ooCreatedAtProto
	ownerOrder.Order.Status = order.OrderStatus(order.OrderStatus_value[oStatus])
	oCreatedAtProto, err := ptypes.TimestampProto(time.Unix(oCreatedAt, 0))
	errorkit.ErrorHandled(err)
	ownerOrder.Order.CreatedAt = oCreatedAtProto
	if oDelivered == 0 {
		ownerOrder.Order.Delivered = false
	} else if oDelivered == 1 {
		ownerOrder.Order.Delivered = true
	}

	return &ownerOrder
}

func (oor *OwnerOrderRented) changeOwnerOrderStatus(ownerOrder *order.OwnerOrder) {
	(*ownerOrder).OrderStatus = order.OrderStatus_RENTED
}
