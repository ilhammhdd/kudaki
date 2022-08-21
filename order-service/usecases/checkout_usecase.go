package usecases

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-order-service/entities/aggregates/order"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/entities/aggregates/rental"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type CheckOut struct {
	DBO DBOperator
}

func (co *CheckOut) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := co.initInOutEvent(in)

	cart := co.retrieveCart(inEvent.CartUuid, outEvent.Tenant.Uuid)
	if cart == nil {
		outEvent.EventStatus.Errors = []string{"authenticated user's cart with the given uuid not found"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	outEvent.Order = co.initOrder(cart.Uuid, outEvent.Tenant.Uuid)
	outEvent.OwnerOrders = co.initOwnerOrders(cart, outEvent.Order)

	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (co *CheckOut) initInOutEvent(in proto.Message) (inEvent *events.CheckOut, outEvent *events.CheckedOut) {
	inEvent = in.(*events.CheckOut)

	outEvent = new(events.CheckedOut)
	outEvent.CheckOut = inEvent
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Tenant = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.Uid = inEvent.Uid

	return
}

func (co *CheckOut) retrieveCart(cartUuid string, tenantUuid string) *rental.Cart {
	row, err := co.DBO.QueryRow("SELECT id,uuid,user_uuid,total_price,total_items,open,created_at FROM kudaki_rental.carts WHERE uuid=? AND user_uuid=?;", cartUuid, tenantUuid)
	errorkit.ErrorHandled(err)

	var cart rental.Cart
	var open int
	var createdAt int64

	if row.Scan(&cart.Id, &cart.Uuid, &cart.UserUuid, &cart.TotalPrice, &cart.TotalItems, &open, &createdAt) == sql.ErrNoRows {
		return nil
	}

	if open == 1 {
		cart.Open = true
	} else if open == 0 {
		cart.Open = false
	}

	createdAtProto, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
	errorkit.ErrorHandled(err)

	cart.CreatedAt = createdAtProto

	return &cart
}

func (co *CheckOut) initOrder(cartUuid string, tenantUuid string) *order.Order {
	return &order.Order{
		CartUuid:   cartUuid,
		CreatedAt:  ptypes.TimestampNow(),
		OrderNum:   uuid.New().String(),
		Status:     order.OrderStatus_PENDING,
		TenantUuid: tenantUuid,
		Uuid:       uuid.New().String()}
}

func (co *CheckOut) initOwnerOrders(cart *rental.Cart, initOrder *order.Order) []*order.OwnerOrder {
	rows, err := co.DBO.Query("SELECT ow.uuid, SUM(ci.total_item), SUM(ci.total_price) FROM kudaki_rental.cart_items ci JOIN kudaki_store.items i ON ci.item_uuid = i.uuid JOIN kudaki_store.storefronts sf ON i.storefront_uuid = sf.uuid JOIN kudaki_user.users ow ON sf.user_uuid = ow.uuid WHERE ci.cart_uuid = ? GROUP BY ow.uuid;", cart.Uuid)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var ownerOrders []*order.OwnerOrder
	for rows.Next() {
		ownerOrder := new(order.OwnerOrder)

		ownerOrder.Uuid = uuid.New().String()
		ownerOrder.OrderStatus = order.OrderStatus_PENDING
		ownerOrder.Order = initOrder

		rows.Scan(
			&ownerOrder.OwnerUuid,
			&ownerOrder.TotalQuantity,
			&ownerOrder.TotalPrice)

		ownerOrders = append(ownerOrders, ownerOrder)
	}

	return ownerOrders
}
