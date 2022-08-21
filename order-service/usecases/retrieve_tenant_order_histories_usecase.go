package usecases

import (
	"database/sql"
	"net/http"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/kudaki-order-service/entities/aggregates/order"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type RetrieveTenantOrderHistories struct {
	DBO           DBOperator
	ResultSchemer ResultSchemer
}

func (rtoh *RetrieveTenantOrderHistories) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := rtoh.initInOutEvent(in)

	cartItems := rtoh.retrieveCartItems(outEvent.Tenant.Uuid, inEvent)
	owners := rtoh.retrieveOwners(outEvent.Tenant.Uuid, inEvent)
	orders := rtoh.retrieveOrders(outEvent.Tenant.Uuid, inEvent)

	rtoh.pasteCartItemToOwner(cartItems, owners)
	rtoh.pasteOwnerToOrder(owners, orders)

	outEvent.Result = rtoh.ResultSchemer.SetResultSources(orders).ParseToResult()
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (rtoh *RetrieveTenantOrderHistories) initInOutEvent(in proto.Message) (inEvent *events.RetrieveTenantsOrderHistories, outEvent *events.TenantOrderHistoriesRetrieved) {
	inEvent = in.(*events.RetrieveTenantsOrderHistories)

	outEvent = new(events.TenantOrderHistoriesRetrieved)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.RetrieveTenantsOrderHistories = inEvent
	outEvent.Tenant = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.Uid = inEvent.Uid

	return
}

func (rtoh *RetrieveTenantOrderHistories) retrieveOrders(tenantUuid string, inEvent *events.RetrieveTenantsOrderHistories) []*OrderTemp {
	rows, err := rtoh.DBO.Query("SELECT o.id, o.uuid, o.cart_uuid, o.order_num, o.status, o.shipment_fee, o.created_at, o.delivered, c.total_price, c.total_items FROM(SELECT o_i.id FROM kudaki_order.orders o_i WHERE o_i.tenant_uuid = ? AND o_i.status = ? LIMIT ?, ?) o_ids JOIN kudaki_order.orders o ON o.id = o_ids.id JOIN kudaki_rental.carts c ON c.uuid = o.cart_uuid;",
		tenantUuid, inEvent.OrderStatus.String(), inEvent.Offset, inEvent.Limit)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var orderTemps []*OrderTemp
	for rows.Next() {
		var orderTemp OrderTemp
		var shipmentFee sql.NullInt64
		var delivered sql.NullBool

		err = rows.Scan(
			&orderTemp.Id,
			&orderTemp.Uuid,
			&orderTemp.CartUuid,
			&orderTemp.OrderNum,
			&orderTemp.StatusT,
			&shipmentFee,
			&orderTemp.CreatedAtT,
			&delivered,
			&orderTemp.TotalPrice,
			&orderTemp.TotalItem)
		errorkit.ErrorHandled(err)

		orderTemp.ShipmentFee = int32(shipmentFee.Int64)
		orderTemp.Delivered = delivered.Bool

		orderTemps = append(orderTemps, &orderTemp)
	}

	return orderTemps
}

func (rtoh *RetrieveTenantOrderHistories) retrieveOwners(tenantUuid string, inEvent *events.RetrieveTenantsOrderHistories) []*OwnerTemp {
	rows, err := rtoh.DBO.Query("SELECT p.full_name, u.email, u.phone_number, oo.total_price, oo.total_quantity, oo.status, oo.order_uuid, oo.uuid FROM(SELECT oo_i.id FROM kudaki_order.owner_orders oo_i JOIN kudaki_order.orders o_i ON oo_i.order_uuid = o_i.uuid WHERE o_i.tenant_uuid = ? AND o_i.status = ? LIMIT ?, ?) oo_ids JOIN kudaki_order.owner_orders oo ON oo.id = oo_ids.id JOIN kudaki_user.users u ON oo.owner_uuid = u.uuid JOIN kudaki_user.profiles p ON u.uuid = p.user_uuid;",
		tenantUuid, inEvent.OrderStatus.String(), inEvent.Offset, inEvent.Limit)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var owners []*OwnerTemp
	for rows.Next() {
		var owner OwnerTemp
		rows.Scan(
			&owner.FullName,
			&owner.Email,
			&owner.PhoneNumber,
			&owner.TotalPrice,
			&owner.TotalItem,
			&owner.OwnerApprovalStatus,
			&owner.OrderUuid,
			&owner.OwnerOrderUuid)
		owners = append(owners, &owner)
	}

	return owners
}

func (rtoh *RetrieveTenantOrderHistories) retrieveCartItems(tenantUuid string, inEvent *events.RetrieveTenantsOrderHistories) []*CartItemTempOrder {
	rows, err := rtoh.DBO.Query("SELECT ci.total_item, ci.total_price, ci.unit_price, ci.duration_from, ci.duration_to, i.name, i.photo, i.price, i.price_duration, oo.uuid FROM kudaki_rental.cart_items ci JOIN kudaki_rental.carts c ON ci.cart_uuid = c.uuid JOIN kudaki_order.orders o ON c.uuid = o.cart_uuid JOIN kudaki_order.owner_orders oo ON o.uuid = oo.order_uuid JOIN kudaki_store.items i ON i.uuid = ci.item_uuid WHERE o.tenant_uuid = ? AND o.status = ?;",
		tenantUuid, inEvent.OrderStatus.String())
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var cartItems []*CartItemTempOrder
	for rows.Next() {
		var cartItem CartItemTempOrder
		var item ItemTempOrder
		rows.Scan(
			&cartItem.TotalItems,
			&cartItem.TotalPrice,
			&cartItem.UnitPrice,
			&cartItem.DurationFromT,
			&cartItem.DurationToT,
			&item.Name,
			&item.Photo,
			&item.Price,
			&item.PriceDuration,
			&cartItem.OwnerOrderUuid)
		cartItem.ItemTempOrder = &item
		cartItems = append(cartItems, &cartItem)
	}

	return cartItems
}

func (rtoh *RetrieveTenantOrderHistories) pasteCartItemToOwner(cartItems []*CartItemTempOrder, owners []*OwnerTemp) {
	for i := 0; i < len(cartItems); i++ {
		for j := 0; j < len(owners); j++ {
			if cartItems[i].OwnerOrderUuid == owners[j].OwnerOrderUuid {
				(*owners[j]).CartItem = append((*owners[j]).CartItem, cartItems[j])
			}
		}
	}
}

func (rtoh *RetrieveTenantOrderHistories) pasteOwnerToOrder(owners []*OwnerTemp, orders []*OrderTemp) {
	for i := 0; i < len(owners); i++ {
		for j := 0; j < len(orders); j++ {
			if owners[i].OrderUuid == orders[j].Uuid {
				(*orders[j]).OwnerTemp = append((*orders[j]).OwnerTemp, owners[j])
			}
		}
	}
}

type OrderTemp struct {
	order.Order
	StatusT    string       `json:"status"`
	CreatedAtT int64        `json:"created_at"`
	TotalPrice int32        `json:"total_price"`
	TotalItem  int32        `json:"total_item"`
	OwnerTemp  []*OwnerTemp `json:"owners"`
}

type OwnerTemp struct {
	FullName            string               `json:"full_name"`
	Email               string               `json:"email"`
	PhoneNumber         string               `json:"phone_number"`
	TotalPrice          int32                `json:"total_price"`
	TotalItem           int32                `json:"total_item"`
	OwnerApprovalStatus string               `json:"owner_approval_status"`
	CartItem            []*CartItemTempOrder `json:"cart_items"`
	OrderUuid           string               `json:"-"`
	OwnerOrderUuid      string               `json:"owner_order_uuid"`
}

type CartItemTempOrder struct {
	TotalItems     int32          `json:"total_items"`
	TotalPrice     int32          `json:"total_price"`
	UnitPrice      int32          `json:"unit_price"`
	DurationFromT  int64          `json:"duration_from"`
	DurationToT    int64          `json:"duration_to"`
	ItemTempOrder  *ItemTempOrder `json:"item"`
	OwnerOrderUuid string         `json:"-"`
}

type ItemTempOrder struct {
	Name          string `json:"name"`
	Photo         string `json:"photo"`
	Price         int32  `json:"price"`
	PriceDuration string `json:"price_duration"`
}
