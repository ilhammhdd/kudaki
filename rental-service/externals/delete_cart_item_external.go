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

type DeleteCartItem struct{}

func (dci *DeleteCartItem) Work() interface{} {
	usecase := usecases.DeleteCartItem{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: dci,
		eventDrivenAdapter:  new(adapters.DeleteCartItem),
		eventDrivenUsecase:  &usecase,
		eventName:           events.RentalServiceCommandTopic_DELETE_CART_ITEM.String(),
		inTopics:            []string{events.RentalServiceCommandTopic_DELETE_CART_ITEM.String()},
		outTopic:            events.RentalServiceEventTopic_CART_ITEM_DELETED.String()}

	ede.handle()
	return nil
}

func (dci *DeleteCartItem) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.CartItemDeleted)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dci.deleteCartItem(out)
	dci.updateCart(out)
}

func (dci *DeleteCartItem) deleteCartItem(out *events.CartItemDeleted) {
	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("DELETE FROM kudaki_rental.cart_items WHERE uuid=?", out.CartItem.Uuid)
	errorkit.ErrorHandled(err)
}

func (dci *DeleteCartItem) updateCart(out *events.CartItemDeleted) {
	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("UPDATE kudaki_rental.carts SET total_price = ?,total_items = ? WHERE uuid = ?;",
		out.CartItem.Cart.TotalPrice, out.CartItem.Cart.TotalItems, out.CartItem.Cart.Uuid)
	errorkit.ErrorHandled(err)
}
