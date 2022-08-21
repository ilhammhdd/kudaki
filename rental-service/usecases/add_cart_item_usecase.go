package usecases

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/ilhammhdd/kudaki-rental-service/entities/aggregates/store"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-rental-service/entities/aggregates/rental"
	"github.com/ilhammhdd/kudaki-rental-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-rental-service/usecases/events"
)

type AddCartItem struct {
	DBO DBOperator
}

func (aci *AddCartItem) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := aci.initInOutEvent(in)

	openCart := aci.retrieveOpenCart(outEvent.Requester)
	item := aci.retrieveItem(inEvent.ItemUuid)
	if item == nil {
		outEvent.EventStatus.Errors = []string{"item with the given uuid not found"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound

		return outEvent
	}

	if item.Amount < inEvent.ItemAmount {
		outEvent.EventStatus.Errors = []string{"insufficient item amount"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound

		return outEvent
	}

	cartItem := aci.retrieveCartItem(item, openCart)
	if cartItem != nil {
		aci.updateCartItem(inEvent, cartItem, item)
	} else {
		cartItem = aci.initCartItem(inEvent, item, openCart)
	}

	cartItem.Cart = openCart

	aci.updateCart(inEvent, cartItem, item)

	outEvent.CartItem = cartItem
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (aci *AddCartItem) initInOutEvent(in proto.Message) (inEvent *events.AddCartItem, outEvent *events.CartItemAdded) {
	inEvent = in.(*events.AddCartItem)

	outEvent = new(events.CartItemAdded)
	outEvent.AddCartItem = inEvent
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Requester = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.Uid = inEvent.Uid

	return
}

func (aci *AddCartItem) retrieveOpenCart(usr *user.User) *rental.Cart {
	row, err := aci.DBO.QueryRow("SELECT id,uuid,user_uuid,total_price,total_items,created_at FROM kudaki_rental.carts WHERE user_uuid=? AND open=1;", usr.Uuid)
	errorkit.ErrorHandled(err)

	var cart rental.Cart
	var createdAt int64
	if row.Scan(
		&cart.Id,
		&cart.Uuid,
		&cart.UserUuid,
		&cart.TotalPrice,
		&cart.TotalItems,
		&createdAt) == sql.ErrNoRows {
		return nil
	}

	cart.CreatedAt, err = ptypes.TimestampProto(time.Unix(createdAt, 0))
	errorkit.ErrorHandled(err)

	return &cart
}

func (aci *AddCartItem) retrieveItem(itemUuid string) *store.Item {
	row, err := aci.DBO.QueryRow("SELECT id,uuid,storefront_uuid,name,amount,unit,price,price_duration,description,photo,rating,length,width,height,color,unit_of_measurement,created_at FROM kudaki_store.items WHERE uuid=?", itemUuid)
	errorkit.ErrorHandled(err)

	var item store.Item
	item.Storefront = new(store.Storefront)
	var priceDuration string
	item.ItemDimension = new(store.ItemDimension)
	var createdAt int64
	var unitOfMeasurement string

	if row.Scan(
		&item.Id, &item.Uuid,
		&item.Storefront.Uuid, &item.Name,
		&item.Amount, &item.Unit,
		&item.Price, &priceDuration,
		&item.Description, &item.Photo,
		&item.Rating, &item.ItemDimension.Length,
		&item.ItemDimension.Width, &item.ItemDimension.Height,
		&item.Color, &unitOfMeasurement,
		&createdAt) == sql.ErrNoRows {
		return nil
	}

	item.ItemDimension.UnitOfMeasurement = store.UnitofMeasurement(store.UnitofMeasurement_value[unitOfMeasurement])

	createdAtProto, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
	errorkit.ErrorHandled(err)
	item.CreatedAt = createdAtProto

	item.PriceDuration = store.PriceDuration(store.PriceDuration_value[priceDuration])

	return &item
}

func (aci *AddCartItem) retrieveCartItem(item *store.Item, cart *rental.Cart) *rental.CartItem {
	row, err := aci.DBO.QueryRow("SELECT id,uuid,item_uuid,total_item,total_price,unit_price,duration_from,duration_to,created_at FROM kudaki_rental.cart_items WHERE cart_uuid=? AND item_uuid=?;", cart.Uuid, item.Uuid)
	errorkit.ErrorHandled(err)

	var cartItem rental.CartItem
	var durationFrom int64
	var durationTo int64
	var createdAt int64

	if row.Scan(
		&cartItem.Id,
		&cartItem.Uuid,
		&cartItem.ItemUuid,
		&cartItem.TotalItem,
		&cartItem.TotalPrice,
		&cartItem.UnitPrice,
		&durationFrom,
		&durationTo,
		&createdAt) == sql.ErrNoRows {
		return nil
	}

	cartItem.Cart = cart

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

func (aci *AddCartItem) updateCartItem(inEvent *events.AddCartItem, cartItem *rental.CartItem, item *store.Item) {
	(*cartItem).TotalItem += inEvent.ItemAmount
	(*cartItem).TotalPrice += (inEvent.ItemAmount * item.Price)
}

func (aci *AddCartItem) initCartItem(inEvent *events.AddCartItem, item *store.Item, cart *rental.Cart) *rental.CartItem {
	durationFromProto, err := ptypes.TimestampProto(time.Unix(inEvent.DurationFrom, 0))
	errorkit.ErrorHandled(err)

	var durationTo int64
	switch item.PriceDuration {
	case store.PriceDuration_DAY:
		durationTo = inEvent.DurationFrom + int64(inEvent.Duration)*8.64e4
	case store.PriceDuration_WEEK:
		durationTo = inEvent.DurationFrom + int64(inEvent.Duration)*604800
	case store.PriceDuration_MONTH:
		durationTo = inEvent.DurationFrom + int64(inEvent.Duration)*2.628e6
	case store.PriceDuration_YEAR:
		durationTo = inEvent.DurationFrom + int64(inEvent.Duration)*3.154e7
	}

	durationToProto, err := ptypes.TimestampProto(time.Unix(durationTo, 0))
	errorkit.ErrorHandled(err)

	return &rental.CartItem{
		Cart:         cart,
		DurationFrom: durationFromProto,
		DurationTo:   durationToProto,
		ItemUuid:     item.Uuid,
		TotalItem:    inEvent.ItemAmount,
		TotalPrice:   item.Price * inEvent.ItemAmount * inEvent.Duration,
		UnitPrice:    item.Price,
		Uuid:         uuid.New().String()}
}

func (aci *AddCartItem) updateCart(inEvent *events.AddCartItem, cartItem *rental.CartItem, item *store.Item) {
	(*cartItem.Cart).TotalItems += inEvent.ItemAmount
	(*cartItem.Cart).TotalPrice += inEvent.ItemAmount * item.Price
}
