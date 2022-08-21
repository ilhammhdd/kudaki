package externals

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-rental-service/adapters"
	"github.com/ilhammhdd/kudaki-rental-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-rental-service/usecases"
	"github.com/ilhammhdd/kudaki-rental-service/usecases/events"
)

type AddCartItem struct{}

func (aci *AddCartItem) Work() interface{} {
	usecase := usecases.AddCartItem{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: aci,
		eventDrivenAdapter:  new(adapters.AddCartItem),
		eventDrivenUsecase:  &usecase,
		eventName:           events.RentalServiceCommandTopic_ADD_CART_ITEM.String(),
		inTopics:            []string{events.RentalServiceCommandTopic_ADD_CART_ITEM.String()},
		outTopic:            events.RentalServiceEventTopic_CART_ITEM_ADDED.String()}

	ede.handle()

	return nil
}

func (aci *AddCartItem) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.CartItemAdded)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	aci.upsertCart(out)
	aci.upsertCartItem(out)
}

func (aci *AddCartItem) upsertCart(out *events.CartItemAdded) {
	dboCart := mysql.NewDBOperation(mysql.CommandDB)

	var open int
	if out.CartItem.Cart.Open {
		open = 1
	} else {
		open = 0
	}

	_, err := dboCart.Command("INSERT INTO kudaki_rental.carts(id,uuid,user_uuid,total_price,total_items,open,created_at) VALUES(?,?,?,?,?,?,UNIX_TIMESTAMP()) ON DUPLICATE KEY UPDATE total_price=?,total_items=?;",
		out.CartItem.Cart.Id,
		out.CartItem.Cart.Uuid,
		out.CartItem.Cart.UserUuid,
		out.CartItem.Cart.TotalPrice,
		out.CartItem.Cart.TotalItems,
		open,
		out.CartItem.Cart.TotalPrice,
		out.CartItem.Cart.TotalItems)
	errorkit.ErrorHandled(err)
}

func (aci *AddCartItem) upsertCartItem(out *events.CartItemAdded) {
	dboCartItem := mysql.NewDBOperation(mysql.CommandDB)

	_, err := dboCartItem.Command("INSERT INTO kudaki_rental.cart_items(id,uuid,cart_uuid,item_uuid,total_item,total_price,unit_price,duration_from,duration_to,created_at) VALUES(?,?,?,?,?,?,?,?,?,UNIX_TIMESTAMP()) ON DUPLICATE KEY UPDATE total_item=?, total_price=?;",
		out.CartItem.Id,
		out.CartItem.Uuid,
		out.CartItem.Cart.Uuid,
		out.CartItem.ItemUuid,
		out.CartItem.TotalItem,
		out.CartItem.TotalPrice,
		out.CartItem.UnitPrice,
		out.CartItem.DurationFrom.GetSeconds(),
		out.CartItem.DurationTo.GetSeconds(),
		out.CartItem.TotalItem,
		out.CartItem.TotalPrice)
	errorkit.ErrorHandled(err)
}
