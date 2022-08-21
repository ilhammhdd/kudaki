package usecases

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/kudaki-order-service/entities/aggregates/order"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type DisapproveOwnerOrder struct {
	DBO DBOperator
}

func (ao *DisapproveOwnerOrder) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := ao.initInOutEvent(in)

	ownerOrder := ao.retrieveOwnerOrder(outEvent.Owner.Uuid, inEvent.OwnerOrderUuid)
	if ownerOrder == nil {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"owner order with the given uuid not found"}
		return outEvent
	}

	ownerOrder.OrderStatus = order.OrderStatus_DISAPPROVED
	outEvent.OwnerOrder = ownerOrder
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (ao *DisapproveOwnerOrder) initInOutEvent(in proto.Message) (inEvent *events.DisapproveOwnerOrder, outEvent *events.OwnerOrderDisapproved) {
	inEvent = in.(*events.DisapproveOwnerOrder)

	outEvent = new(events.OwnerOrderDisapproved)
	outEvent.DisapproveOwnerOrder = inEvent
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Owner = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.Uid = inEvent.Uid

	return
}

func (ao *DisapproveOwnerOrder) retrieveOwnerOrder(ownerUuid, ownerOrderUuid string) *order.OwnerOrder {
	row, err := ao.DBO.QueryRow("SELECT oo.id, oo.uuid, oo.owner_uuid, oo.total_price, oo.total_quantity, oo.status, oo.created_at, o.id AS o_id, o.uuid AS o_uuid,o.cart_uuid AS o_cart_uuid, o.tenant_uuid AS o_tenant_uuid, o.order_num AS o_order_num, o.status AS o_status, o.shipment_fee AS o_shipment_fee, o.delivered AS o_delivered, o.created_at AS o_created_at FROM kudaki_order.owner_orders oo JOIN kudaki_order.orders o ON oo.order_uuid = o.uuid WHERE oo.uuid = ? AND oo.owner_uuid = ?;",
		ownerOrderUuid, ownerUuid)
	errorkit.ErrorHandled(err)

	var ownerOrder order.OwnerOrder
	ownerOrder.Order = new(order.Order)

	var ownerOrderStat string
	var ownerOrderCreatedAt int64
	var orderStat string
	var orderCreatedAt int64
	var shipmentFee sql.NullInt64
	var delivered sql.NullBool

	err = row.Scan(
		&ownerOrder.Id,
		&ownerOrder.Uuid,
		&ownerOrder.OwnerUuid,
		&ownerOrder.TotalPrice,
		&ownerOrder.TotalQuantity,
		&ownerOrderStat,
		&ownerOrderCreatedAt,
		&ownerOrder.Order.Id,
		&ownerOrder.Order.Uuid,
		&ownerOrder.Order.CartUuid,
		&ownerOrder.Order.TenantUuid,
		&ownerOrder.Order.OrderNum,
		&orderStat,
		&shipmentFee,
		&delivered,
		&orderCreatedAt)
	if errorkit.ErrorHandled(err) && err == sql.ErrNoRows {
		return nil
	}

	ownerOrder.Order.Status = order.OrderStatus(order.OrderStatus_value[orderStat])
	orderCreatedAtProto, err := ptypes.TimestampProto(time.Unix(orderCreatedAt, 0))
	errorkit.ErrorHandled(err)
	ownerOrder.Order.CreatedAt = orderCreatedAtProto
	ownerOrder.Order.ShipmentFee = int32(shipmentFee.Int64)
	ownerOrder.Order.Delivered = delivered.Bool

	ownerOrder.OrderStatus = order.OrderStatus(order.OrderStatus_value[ownerOrderStat])
	ownerOrderCreatedAtProto, err := ptypes.TimestampProto(time.Unix(ownerOrderCreatedAt, 0))
	errorkit.ErrorHandled(err)
	ownerOrder.CreatedAt = ownerOrderCreatedAtProto

	return &ownerOrder
}
