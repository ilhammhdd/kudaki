package usecases

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-order-service/entities/aggregates/user"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/entities/aggregates/order"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type TenantReviewsOwnerOrder struct {
	DBO DBOperator
}

func (tro *TenantReviewsOwnerOrder) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := tro.initInOutEvent(in)
	ownerOrder := tro.retrieveOwnerOrder(outEvent.Tenant, inEvent)
	if ownerOrder == nil {
		outEvent.EventStatus.Errors = []string{"owner order with the given uuid not found or not done yet or it doesn't belong to the user"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.OwnerOrderReview = tro.initOwnerOrderReview(outEvent.Tenant, ownerOrder, inEvent)
	return outEvent
}

func (tro *TenantReviewsOwnerOrder) initInOutEvent(in proto.Message) (inEvent *events.TenantReviewOwnerOrder, outEvent *events.TenantReviewedOwnerOrder) {
	inEvent = in.(*events.TenantReviewOwnerOrder)

	outEvent = new(events.TenantReviewedOwnerOrder)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Tenant = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.TenantReviewOwnerOrder = inEvent
	outEvent.Uid = inEvent.Uid

	return
}

func (tro *TenantReviewsOwnerOrder) retrieveOwnerOrder(tenant *user.User, inEvent *events.TenantReviewOwnerOrder) *order.OwnerOrder {
	row, err := tro.DBO.QueryRow("SELECT oo.id,oo.uuid,oo.owner_uuid,oo.order_uuid,oo.total_price,oo.total_quantity,oo.status,oo.created_at FROM kudaki_order.owner_orders oo JOIN kudaki_order.orders o ON oo.order_uuid = o.uuid WHERE oo.uuid = ? AND oo.status = ? AND o.status = ? AND o.tenant_uuid = ?;",
		inEvent.OwnerOrderUuid,
		order.OrderStatus_DONE.String(),
		order.OrderStatus_DONE.String(),
		tenant.Uuid)
	errorkit.ErrorHandled(err)

	var ownerOrder order.OwnerOrder
	var status string
	var createdAt int64
	ownerOrder.Order = new(order.Order)

	if row.Scan(
		&ownerOrder.Id,
		&ownerOrder.Uuid,
		&ownerOrder.OwnerUuid,
		&ownerOrder.Order.Uuid,
		&ownerOrder.TotalPrice,
		&ownerOrder.TotalQuantity,
		&status,
		&createdAt) == sql.ErrNoRows {
		return nil
	}

	createdAtProto, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
	errorkit.ErrorHandled(err)
	ownerOrder.CreatedAt = createdAtProto
	ownerOrder.OrderStatus = order.OrderStatus(order.OrderStatus_value[status])

	return &ownerOrder
}

func (tro *TenantReviewsOwnerOrder) initOwnerOrderReview(tenant *user.User, ownerOrder *order.OwnerOrder, inEvent *events.TenantReviewOwnerOrder) *order.OwnerOrderReview {
	return &order.OwnerOrderReview{
		CreatedAt:  ptypes.TimestampNow(),
		OwnerOrder: ownerOrder,
		Rating:     inEvent.Rating,
		Review:     inEvent.Review,
		TenantUuid: tenant.Uuid,
		Uuid:       uuid.New().String(),
	}
}
