package usecases

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/entities/aggregates/order"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type OwnerConfirmReturnment struct {
	DBO DBOperator
}

func (ocr *OwnerConfirmReturnment) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := ocr.initInOutEvent(in)

	ownerOrder := ocr.retrieveOwnerOrder(outEvent.Owner.Uuid, inEvent.OwnerOrderUuid)
	if ownerOrder == nil {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"owner order with the given uuid not found"}
		return outEvent
	}

	outEvent.OwnerOrder = ownerOrder
	outEvent.EventStatus.HttpCode = http.StatusOK
	return outEvent
}

func (ocr *OwnerConfirmReturnment) initInOutEvent(in proto.Message) (inEvent *events.OwnerConfirmReturnment, outEvent *events.OwnerConfirmedReturnment) {
	inEvent = in.(*events.OwnerConfirmReturnment)

	outEvent = new(events.OwnerConfirmedReturnment)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Owner = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.OwnerConfirmReturnment = inEvent
	outEvent.Uid = inEvent.Uid

	return
}

func (ocr *OwnerConfirmReturnment) retrieveOwnerOrder(ownerUuid, ownerOrderUuid string) *order.OwnerOrder {
	row, err := ocr.DBO.QueryRow("SELECT id, uuid, order_uuid, owner_uuid, total_price, total_quantity, status, created_at FROM kudaki_order.owner_orders WHERE uuid=? AND owner_uuid=?;",
		ownerOrderUuid, ownerUuid)
	errorkit.ErrorHandled(err)

	var ownerOrder order.OwnerOrder
	ownerOrder.Order = new(order.Order)
	var status string
	var createdAt int64

	if row.Scan(
		&ownerOrder.Id,
		&ownerOrder.Uuid,
		&ownerOrder.Order.Uuid,
		&ownerOrder.OwnerUuid,
		&ownerOrder.TotalPrice,
		&ownerOrder.TotalQuantity,
		&status,
		&createdAt) == sql.ErrNoRows {
		return nil
	}

	ownerOrder.OrderStatus = order.OrderStatus(order.OrderStatus_value[status])
	createdAtProto, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
	errorkit.ErrorHandled(err)
	ownerOrder.CreatedAt = createdAtProto

	return &ownerOrder
}
