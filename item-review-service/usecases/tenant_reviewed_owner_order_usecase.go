package usecases

import (
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/ilhammhdd/kudaki-item-review-service/entities/aggregates/order"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/ilhammhdd/kudaki-item-review-service/entities/aggregates/store"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases/events"
)

type TenantReviewedOwnerOrder struct {
	DBO DBOperator
}

func (troo *TenantReviewedOwnerOrder) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := troo.initInOutEvent(in)

	itemsWithCount := troo.retrieveItems(inEvent)
	itemReviews := troo.initItemReviews(inEvent.OwnerOrderReview, itemsWithCount)
	troo.recalculateItemsRating(itemsWithCount, itemReviews)

	outEvent.ItemReviews = itemReviews
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (troo *TenantReviewedOwnerOrder) initInOutEvent(in proto.Message) (inEvent *events.TenantReviewedOwnerOrder, outEvent *events.ItemsReviewed) {
	inEvent = in.(*events.TenantReviewedOwnerOrder)

	outEvent = new(events.ItemsReviewed)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Requester = inEvent.Tenant
	outEvent.Uid = inEvent.Uid

	return
}

type ItemWithCountReview struct {
	store.Item
	CountedReviews int64
}

func (troo *TenantReviewedOwnerOrder) retrieveItems(inEvent *events.TenantReviewedOwnerOrder) []*ItemWithCountReview {
	rows, err := troo.DBO.Query("SELECT i.id, i.uuid, i.storefront_uuid, i.name, i.amount, i.unit, i.price, i.price_duration, i.description, i.photo, i.rating, i.total_raw_rating, i.length, i.width, i.height, i.color, i.unit_of_measurement, i.created_at, (SELECT COUNT(ir_i.id) FROM kudaki_store.item_reviews ir_i JOIN kudaki_store.items i_i ON ir_i.item_uuid = i_i.uuid JOIN kudaki_store.storefronts sf_i ON i_i.storefront_uuid = sf_i.uuid WHERE sf_i.user_uuid = ?) counted_reviews FROM kudaki_store.items i JOIN kudaki_rental.cart_items ci ON i.uuid = ci.item_uuid JOIN kudaki_store.storefronts sf ON i.storefront_uuid = sf.uuid WHERE ci.cart_uuid =(SELECT c_i.uuid FROM kudaki_rental.carts c_i JOIN kudaki_order.orders o_i ON c_i.uuid = o_i.cart_uuid WHERE o_i.uuid = ?) AND sf.user_uuid = ?;",
		inEvent.OwnerOrderReview.OwnerOrder.OwnerUuid,
		inEvent.OwnerOrderReview.OwnerOrder.Order.Uuid,
		inEvent.OwnerOrderReview.OwnerOrder.OwnerUuid)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var items []*ItemWithCountReview
	for rows.Next() {
		var item ItemWithCountReview
		item.Storefront = new(store.Storefront)
		item.ItemDimension = new(store.ItemDimension)
		var priceDuration string
		var unitOfMeasurement string
		var createdAt int64
		rows.Scan(
			&item.Id,
			&item.Uuid,
			&item.Storefront.Uuid,
			&item.Name,
			&item.Amount,
			&item.Unit,
			&item.Price,
			&priceDuration,
			&item.Description,
			&item.Photo,
			&item.Rating,
			&item.TotalRawRating,
			&item.ItemDimension.Length,
			&item.ItemDimension.Width,
			&item.ItemDimension.Height,
			&item.Color,
			&unitOfMeasurement,
			&createdAt,
			&item.CountedReviews)

		createdAtProto, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
		errorkit.ErrorHandled(err)
		item.ItemDimension.UnitOfMeasurement = store.UnitofMeasurement(store.UnitofMeasurement_value[unitOfMeasurement])
		item.PriceDuration = store.PriceDuration(store.PriceDuration_value[priceDuration])
		item.CreatedAt = createdAtProto

		items = append(items, &item)
	}

	return items
}

func (troo *TenantReviewedOwnerOrder) initItemReviews(ownerOrderReview *order.OwnerOrderReview, items []*ItemWithCountReview) []*store.ItemReview {
	var itemReviews = make([]*store.ItemReview, len(items))

	for i := 0; i < len(items); i++ {
		itemReviews[i] = &store.ItemReview{
			CreatedAt: ptypes.TimestampNow(),
			Item:      &items[i].Item,
			Rating:    ownerOrderReview.Rating,
			Review:    ownerOrderReview.Review,
			UserUuid:  ownerOrderReview.TenantUuid,
			Uuid:      uuid.New().String()}
	}

	return itemReviews
}

func (troo *TenantReviewedOwnerOrder) recalculateItemsRating(items []*ItemWithCountReview, itemReviews []*store.ItemReview) {
	for i := 0; i < len(itemReviews); i++ {
		(*items[i]).CountedReviews++
		(*itemReviews[i].Item).TotalRawRating += itemReviews[i].Rating
		(*itemReviews[i].Item).Rating = itemReviews[i].Item.TotalRawRating / float64(items[i].CountedReviews)
	}
}
