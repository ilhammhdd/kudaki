package externals

import (
	"log"
	"net/http"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-order-service/adapters"
	"github.com/ilhammhdd/kudaki-order-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-order-service/usecases"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type CheckOut struct{}

func (co *CheckOut) Work() interface{} {
	usecase := usecases.CheckOut{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: co,
		eventDrivenAdapter:  new(adapters.CheckOut),
		eventDrivenUsecase:  &usecase,
		eventName:           events.OrderServiceCommandTopic_CHECK_OUT.String(),
		inTopics:            []string{events.OrderServiceCommandTopic_CHECK_OUT.String()},
		outTopic:            events.OrderServiceEventTopic_CHECKED_OUT.String()}

	ede.handle()
	return nil
}

func (co *CheckOut) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.CheckedOut)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	co.insertOrder(out)
	co.insertOwnerOrder(out)
}

func (co *CheckOut) insertOrder(out *events.CheckedOut) {
	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("INSERT INTO kudaki_order.orders(uuid,cart_uuid,tenant_uuid,order_num,status,created_at) VALUES(?,?,?,?,?,UNIX_TIMESTAMP());",
		out.Order.Uuid, out.Order.CartUuid, out.Order.TenantUuid, out.Order.OrderNum, out.Order.Status.String())
	errorkit.ErrorHandled(err)
}

func (co *CheckOut) insertOwnerOrder(out *events.CheckedOut) {
	dbo := mysql.NewDBOperation(mysql.CommandDB)
	for i := 0; i < len(out.OwnerOrders); i++ {
		log.Println("owner orders uuid : ", out.OwnerOrders[i].Uuid)
		log.Println("owner orders order uuid : ", out.OwnerOrders[i].Order.Uuid)
		log.Println("owner orders owner uuid : ", out.OwnerOrders[i].OwnerUuid)
		log.Println("owner orders total price : ", out.OwnerOrders[i].TotalPrice)
		log.Println("owner orders total quantity : ", out.OwnerOrders[i].TotalQuantity)
		log.Println("owner orders order status : ", out.OwnerOrders[i].OrderStatus.String())

		_, err := dbo.Command("INSERT INTO kudaki_order.owner_orders(uuid,order_uuid,owner_uuid,total_price,total_quantity,status,created_at) VALUES(?,?,?,?,?,?,UNIX_TIMESTAMP());",
			out.OwnerOrders[i].Uuid,
			out.OwnerOrders[i].Order.Uuid,
			out.OwnerOrders[i].OwnerUuid,
			out.OwnerOrders[i].TotalPrice,
			out.OwnerOrders[i].TotalQuantity,
			out.OwnerOrders[i].OrderStatus.String())
		errorkit.ErrorHandled(err)
	}
}
