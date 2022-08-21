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

type UpdateCartItem struct{}

func (uci *UpdateCartItem) Work() interface{} {
	usecase := usecases.UpdateCartItem{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: uci,
		eventDrivenAdapter:  new(adapters.UpdateCartItem),
		eventDrivenUsecase:  &usecase,
		eventName:           events.RentalServiceCommandTopic_UPDATE_CART_ITEM.String(),
		inTopics:            []string{events.RentalServiceCommandTopic_UPDATE_CART_ITEM.String()},
		outTopic:            events.RentalServiceEventTopic_CART_ITEMS_UPDATED.String()}

	ede.handle()
	return nil
}

func (uci *UpdateCartItem) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.CartItemsUpdated)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	uci.updateCartItem(out)
	uci.updateCart(out)
}

func (uci *UpdateCartItem) updateCartItem(out *events.CartItemsUpdated) {
	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("UPDATE kudaki_rental.cart_items SET total_item = ?, total_price = ? WHERE uuid=?",
		out.UpdatedCartItem[0].TotalItem, out.UpdatedCartItem[0].TotalPrice, out.UpdatedCartItem[0].Uuid)
	errorkit.ErrorHandled(err)
}

func (uci *UpdateCartItem) updateCart(out *events.CartItemsUpdated) {
	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("UPDATE kudaki_rental.carts SET total_price = ?,total_items = ? WHERE uuid = ?;",
		out.UpdatedCartItem[0].Cart.TotalPrice, out.UpdatedCartItem[0].Cart.TotalItems, out.UpdatedCartItem[0].Cart.Uuid)
	errorkit.ErrorHandled(err)
}
