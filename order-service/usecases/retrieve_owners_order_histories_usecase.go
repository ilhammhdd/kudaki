package usecases

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/entities/aggregates/order"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type RetrieveOwnersOrderHistories struct {
	DBO           DBOperator
	ResultSchemer ResultSchemer
}

func (rooh *RetrieveOwnersOrderHistories) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := rooh.initInOutEvent(in)

	ownerOrders := rooh.retrieveHistories(inEvent, outEvent.Owner.Uuid)

	outEvent.Result = rooh.ResultSchemer.SetResultSources(ownerOrders).ParseToResult()
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (rooh *RetrieveOwnersOrderHistories) initInOutEvent(in proto.Message) (inEvent *events.RetrieveOwnersOrderHistories, outEvent *events.OwnersOrderHistoriesRetrieved) {
	inEvent = in.(*events.RetrieveOwnersOrderHistories)

	outEvent = new(events.OwnersOrderHistoriesRetrieved)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.RetrieveOwnersOrderHistories = inEvent
	outEvent.Owner = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.Uid = inEvent.Uid

	return
}

func (rooh *RetrieveOwnersOrderHistories) retrieveHistories(inEvent *events.RetrieveOwnersOrderHistories, ownerUuid string) []*OwnerOrderTemp {
	rows, err := rooh.DBO.Query("SELECT oo.id, oo.uuid, oo.order_uuid AS tenant_order_uuid, o.cart_uuid, oo.status, oo.created_at, oo.total_price, oo.total_quantity, tu.email AS tenant_email, tp.full_name AS tenant_full_name FROM (SELECT oo_i.id FROM kudaki_order.owner_orders oo_i WHERE oo_i.status=? LIMIT ?, ? ) oo_ids JOIN kudaki_order.owner_orders oo ON oo_ids.id = oo.id JOIN kudaki_order.orders o ON oo.order_uuid = o.uuid JOIN kudaki_user.users tu ON o.tenant_uuid = tu.uuid JOIN kudaki_user.profiles tp ON tu.uuid = tp.user_uuid WHERE oo.owner_uuid = ?;",
		inEvent.OrderStatus.String(), inEvent.Offset, inEvent.Limit, ownerUuid)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var ownerOrders []*OwnerOrderTemp
	for rows.Next() {
		var ownerOrder OwnerOrderTemp
		ownerOrder.Tenant = new(TenantTemp)

		rows.Scan(
			&ownerOrder.Id,
			&ownerOrder.Uuid,
			&ownerOrder.TenantOrderUuid,
			&ownerOrder.CartUuid,
			&ownerOrder.StatusT,
			&ownerOrder.CreatedAtT,
			&ownerOrder.TotalPriceT,
			&ownerOrder.TotalItemT,
			&ownerOrder.Tenant.Email,
			&ownerOrder.Tenant.FullName)

		ownerOrders = append(ownerOrders, &ownerOrder)
	}

	return ownerOrders
}

type OwnerOrderTemp struct {
	order.OwnerOrder
	TenantOrderUuid string      `json:"tenant_order_uuid"`
	CartUuid        string      `json:"cart_uuid"`
	StatusT         string      `json:"status"`
	CreatedAtT      int64       `json:"created_at"`
	TotalPriceT     int64       `json:"total_price"`
	TotalItemT      int64       `json:"total_item"`
	Tenant          *TenantTemp `json:"tenant"`
}

type TenantTemp struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}
