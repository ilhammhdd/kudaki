package usecases

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-rental-service/entities/aggregates/rental"
	"github.com/ilhammhdd/kudaki-rental-service/usecases/events"
)

type DeleteCartItem struct {
	DBO DBOperator
}

func (dci *DeleteCartItem) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := dci.initInOutEvent(in)

	cartItem := dci.retrieveCartItem(inEvent)
	if cartItem == nil {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"cart item not found"}

		return outEvent
	}
	dci.recalculateCart(cartItem, cartItem.Cart)

	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.CartItem = cartItem

	return outEvent
}

func (dci *DeleteCartItem) initInOutEvent(in proto.Message) (inEvent *events.DeleteCartItem, outEvent *events.CartItemDeleted) {
	inEvent = in.(*events.DeleteCartItem)

	outEvent = new(events.CartItemDeleted)
	outEvent.DeleteCartItem = inEvent
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.User = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.Uid = inEvent.Uid

	return
}

func (dci *DeleteCartItem) retrieveCartItem(inEvent *events.DeleteCartItem) *rental.CartItem {
	row, err := dci.DBO.QueryRow("SELECT id,uuid,cart_uuid,item_uuid,total_item,total_price,unit_price,duration_from,duration_to,created_at FROM kudaki_rental.cart_items WHERE uuid=?;", inEvent.CartItemUuid)

	var cartItem rental.CartItem
	cartItem.Cart = new(rental.Cart)
	var durationFrom, durationTo, createdAt int64

	if row.Scan(
		&cartItem.Id,
		&cartItem.Uuid,
		&cartItem.Cart.Uuid,
		&cartItem.ItemUuid,
		&cartItem.TotalItem,
		&cartItem.TotalPrice,
		&cartItem.UnitPrice,
		&durationFrom,
		&durationTo,
		&createdAt) == sql.ErrNoRows {
		return nil
	}

	cartItem.Cart = dci.retrieveCart(cartItem.Cart.Uuid)

	durationFromProto, err := ptypes.TimestampProto(time.Unix(durationFrom, 0))
	errorkit.ErrorHandled(err)
	durationToProto, err := ptypes.TimestampProto(time.Unix(durationTo, 0))
	errorkit.ErrorHandled(err)
	createdAtProto, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
	errorkit.ErrorHandled(err)

	cartItem.DurationFrom = durationFromProto
	cartItem.DurationTo = durationToProto
	cartItem.CreatedAt = createdAtProto

	return &cartItem
}

func (dci *DeleteCartItem) retrieveCart(cartUuid string) *rental.Cart {
	row, err := dci.DBO.QueryRow("SELECT id,uuid,user_uuid,total_price,total_items,open,created_at FROM kudaki_rental.carts WHERE uuid=?;", cartUuid)
	errorkit.ErrorHandled(err)

	var cart rental.Cart
	var open int32
	var createdAt int64

	if row.Scan(
		&cart.Id,
		&cart.Uuid,
		&cart.UserUuid,
		&cart.TotalPrice,
		&cart.TotalItems,
		&open,
		&createdAt) == sql.ErrNoRows {
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

func (dci *DeleteCartItem) recalculateCart(cartItem *rental.CartItem, cart *rental.Cart) {
	(*cart).TotalItems -= cartItem.TotalItem
	(*cart).TotalPrice -= cartItem.TotalPrice
}
