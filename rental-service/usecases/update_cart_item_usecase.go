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

type UpdateCartItem struct {
	DBO DBOperator
}

func (uci *UpdateCartItem) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := uci.initInOutEvent(in)

	initialCartItem := uci.retrieveCartItem(inEvent)
	if initialCartItem == nil {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"cart item with the given uuid not found"}

		return outEvent
	}
	updatedCartItem := uci.recalculateCartItemAndCart(inEvent, initialCartItem)

	outEvent.InitialCartItem = append(outEvent.InitialCartItem, initialCartItem)
	outEvent.UpdatedCartItem = append(outEvent.UpdatedCartItem, updatedCartItem)
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (uci *UpdateCartItem) initInOutEvent(in proto.Message) (inEvent *events.UpdateCartItem, outEvent *events.CartItemsUpdated) {
	inEvent = in.(*events.UpdateCartItem)

	outEvent = new(events.CartItemsUpdated)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Uid = inEvent.Uid
	outEvent.UpdateCartItem = inEvent
	outEvent.User = GetUserFromKudakiToken(inEvent.KudakiToken)

	return
}

func (uci *UpdateCartItem) retrieveCartItem(inEvent *events.UpdateCartItem) *rental.CartItem {
	row, err := uci.DBO.QueryRow("SELECT id,uuid,cart_uuid,item_uuid,total_item,total_price,unit_price,duration_from,duration_to,created_at FROM kudaki_rental.cart_items WHERE uuid=?;", inEvent.CartItemUuid)

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

	cartItem.Cart = uci.retrieveCart(cartItem.Cart.Uuid)

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

func (uci *UpdateCartItem) retrieveCart(cartUuid string) *rental.Cart {
	row, err := uci.DBO.QueryRow("SELECT id,uuid,user_uuid,total_price,total_items,open,created_at FROM kudaki_rental.carts WHERE uuid=?;", cartUuid)
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

func (uci *UpdateCartItem) recalculateCartItemAndCart(inEvent *events.UpdateCartItem, initialCartItem *rental.CartItem) (updatedCartItem *rental.CartItem) {
	updatedCartItemTemp := *initialCartItem
	updatedCart := *initialCartItem.Cart

	var totalItemDiff int32
	if initialCartItem.TotalItem > inEvent.TotalItem {
		totalItemDiff = initialCartItem.TotalItem - inEvent.TotalItem
		updatedCartItemTemp.TotalItem -= totalItemDiff
		updatedCartItemTemp.TotalPrice -= (totalItemDiff * initialCartItem.UnitPrice)
		if updatedCartItemTemp.TotalItem < 0 {
			updatedCartItemTemp.TotalItem = 0
			updatedCartItemTemp.TotalPrice = 0
		}

		updatedCart.TotalItems -= totalItemDiff
		updatedCart.TotalPrice -= (totalItemDiff * initialCartItem.UnitPrice)
		if updatedCart.TotalItems < 0 {
			updatedCart.TotalItems = 0
			updatedCart.TotalPrice = 0
		}
	} else if initialCartItem.TotalItem < inEvent.TotalItem {
		totalItemDiff = inEvent.TotalItem - initialCartItem.TotalItem
		updatedCartItemTemp.TotalItem += totalItemDiff
		updatedCartItemTemp.TotalPrice += (totalItemDiff * initialCartItem.UnitPrice)

		updatedCart.TotalItems += totalItemDiff
		updatedCart.TotalPrice += (totalItemDiff * initialCartItem.UnitPrice)
	}

	updatedCartItemTemp.Cart = &updatedCart
	return &updatedCartItemTemp
}
