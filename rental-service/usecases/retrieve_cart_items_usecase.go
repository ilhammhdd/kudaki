package usecases

import (
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-rental-service/entities/aggregates/rental"
	"github.com/ilhammhdd/kudaki-rental-service/entities/aggregates/store"
	"github.com/ilhammhdd/kudaki-rental-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-rental-service/usecases/events"
)

type RetrieveCartItems struct {
	DBO           DBOperator
	ResultSchemer ResultSchemer
}

func (rci *RetrieveCartItems) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := rci.initIntOutEvent(in)

	cart := rci.retrieveCart(outEvent.Requester)
	storefront := rci.retrieveStorefronts(cart, inEvent)
	cartItems := rci.retrieveCartItems(inEvent, cart)
	rci.pasteCartItemToStorefront(storefront, cartItems)

	outEvent.Result = rci.ResultSchemer.SetResultSources(cart, storefront).ParseToResult()
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (rci *RetrieveCartItems) initIntOutEvent(in proto.Message) (inEvent *events.RetrieveCartItems, outEvent *events.CartItemsRetrieved) {
	inEvent = in.(*events.RetrieveCartItems)

	outEvent = new(events.CartItemsRetrieved)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Requester = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.RetrieveCartItems = inEvent
	outEvent.Uid = inEvent.Uid

	return
}

func (rci *RetrieveCartItems) retrieveCart(usr *user.User) *CartTemp {
	row, err := rci.DBO.QueryRow("SELECT id,uuid,user_uuid,total_price,total_items,created_at FROM kudaki_rental.carts WHERE user_uuid=? AND open=1;", usr.Uuid)
	errorkit.ErrorHandled(err)

	var cartTemp CartTemp
	row.Scan(
		&cartTemp.Id,
		&cartTemp.Uuid,
		&cartTemp.UserUuid,
		&cartTemp.TotalPrice,
		&cartTemp.TotalItems,
		&cartTemp.CreatedAtT)

	return &cartTemp
}

func (rci *RetrieveCartItems) retrieveCartItems(inEvent *events.RetrieveCartItems, cartTemp *CartTemp) []*CartItemTemp {
	rows, err := rci.DBO.Query("SELECT ci.id, ci.uuid, ci.total_item, ci.total_price, ci.unit_price, ci.duration_from, ci.duration_to, i.id AS item_id, i.uuid AS item_uuid, i.name, i.photo, i.price, i.price_duration, sf.uuid AS storefront_uuid FROM (SELECT ci_i.id FROM kudaki_rental.cart_items ci_i JOIN kudaki_store.items i_i ON ci_i.item_uuid = i_i.uuid JOIN kudaki_store.storefronts sf_i ON i_i.storefront_uuid = sf_i.uuid WHERE ci_i.cart_uuid = ? LIMIT ?, ?) ci_ids JOIN kudaki_rental.cart_items ci ON ci.id = ci_ids.id JOIN kudaki_store.items i ON ci.item_uuid = i.uuid JOIN kudaki_store.storefronts sf ON i.storefront_uuid = sf.uuid;",
		cartTemp.Uuid, inEvent.Offset, inEvent.Limit)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var cartItemTemps []*CartItemTemp
	for rows.Next() {
		var cartItemTemp CartItemTemp
		cartItemTemp.Item = new(store.Item)

		err = rows.Scan(
			&cartItemTemp.Id,
			&cartItemTemp.Uuid,
			&cartItemTemp.TotalItem,
			&cartItemTemp.TotalPrice,
			&cartItemTemp.UnitPrice,
			&cartItemTemp.DurationFromT,
			&cartItemTemp.DurationToT,
			&cartItemTemp.Item.Id,
			&cartItemTemp.Item.Uuid,
			&cartItemTemp.Item.Name,
			&cartItemTemp.Item.Photo,
			&cartItemTemp.Item.Price,
			&cartItemTemp.PriceDurationT,
			&cartItemTemp.StorefrontUuid)
		errorkit.ErrorHandled(err)

		cartItemTemps = append(cartItemTemps, &cartItemTemp)
	}

	return cartItemTemps
}

func (rci *RetrieveCartItems) retrieveStorefronts(cartTemp *CartTemp, inEvent *events.RetrieveCartItems) []*StorefrontTemp {
	rows, err := rci.DBO.Query("SELECT sf.id, sf.uuid, p.full_name, u.email, u.phone_number FROM (SELECT id FROM kudaki_rental.cart_items ci_i WHERE ci_i.cart_uuid = ? LIMIT ?, ? ) ci_ids JOIN kudaki_rental.cart_items ci ON ci.id = ci_ids.id JOIN kudaki_store.items i ON ci.item_uuid = i.uuid JOIN kudaki_store.storefronts sf ON sf.uuid = i.storefront_uuid JOIN kudaki_user.users u ON sf.user_uuid = u.uuid JOIN kudaki_user.profiles p ON u.uuid = p.user_uuid GROUP BY sf.uuid;",
		cartTemp.Uuid, inEvent.Offset, inEvent.Limit)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var storefrontTemps []*StorefrontTemp

	for rows.Next() {
		var storefrontTemp StorefrontTemp

		err = rows.Scan(
			&storefrontTemp.Id,
			&storefrontTemp.Uuid,
			&storefrontTemp.OwnerName,
			&storefrontTemp.OwnerEmail,
			&storefrontTemp.OwnerPhoneNumber)
		errorkit.ErrorHandled(err)

		storefrontTemps = append(storefrontTemps, &storefrontTemp)
	}

	return storefrontTemps
}

func (rci *RetrieveCartItems) pasteCartItemToStorefront(storefrontTemps []*StorefrontTemp, cartItemTemps []*CartItemTemp) {
	for i := 0; i < len(cartItemTemps); i++ {
		for j := 0; j < len(storefrontTemps); j++ {
			if cartItemTemps[i].StorefrontUuid == storefrontTemps[j].Uuid {
				log.Println("storefront uuid equals : ", cartItemTemps[i].StorefrontUuid, " == ", storefrontTemps[j].Uuid)
				(*storefrontTemps[j]).CartItems = append((*storefrontTemps[j]).CartItems, cartItemTemps[i])
			}
		}
	}
}

type CartTemp struct {
	rental.Cart
	CreatedAtT int64 `json:"created_at"`
}

type StorefrontTemp struct {
	Id               int64           `json:"-"`
	Uuid             string          `json:"-"`
	OwnerName        string          `json:"owner_name"`
	OwnerEmail       string          `json:"owner_email"`
	OwnerPhoneNumber string          `json:"owner_phone_number"`
	CartItems        []*CartItemTemp `json:"cart_items"`
}

type CartItemTemp struct {
	Id             int64       `json:"id"`
	Uuid           string      `json:"uuid"`
	TotalItem      int32       `json:"total_item"`
	TotalPrice     int32       `json:"total_price"`
	UnitPrice      int32       `json:"unit_price"`
	DurationFromT  int64       `json:"duration_from"`
	DurationToT    int64       `json:"duration_to"`
	PriceDurationT string      `json:"price_duration"`
	Item           *store.Item `json:"item"`
	StorefrontUuid string      `json:"-"`
}
