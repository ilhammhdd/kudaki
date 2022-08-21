package usecases

import (
	"github.com/golang/protobuf/proto"
)

type ReviewItem struct {
	DBO DBOperator
}

func (ri *ReviewItem) Handle(in proto.Message) (out proto.Message) {
	// inEvent, outEvent := ri.initInOutEvent(in)

	// item := ri.retrieveItem(inEvent)
	// itemReview := ri.initItemReview(outEvent.Requester, item, inEvent)
	// ri.recalculateItemRating(item, itemReview)

	// outEvent.ItemReview = itemReview
	// outEvent.EventStatus.HttpCode = http.StatusOK

	// return outEvent
	return nil
}

/* func (ri *ReviewItem) initInOutEvent(in proto.Message) (inEvent *events.ReviewItem, outEvent *events.ItemReviewed) {
	inEvent = in.(*events.ReviewItem)

	outEvent = new(events.ItemReviewed)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Requester = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.ReviewItem = inEvent
	outEvent.Uid = inEvent.Uid

	return
}

func (ri *ReviewItem) retrieveItem(inEvent *events.ReviewItem) *store.Item {
	row, err := ri.DBO.QueryRow("SELECT id,uuid,storefront_uuid,name,amount,unit,price,price_duration,description,photo,rating,length,width,height,color,unit_of_measurement,created_at FROM kudaki_store.items WHERE uuid=?", inEvent.ItemUuid)
	errorkit.ErrorHandled(err)

	var item store.Item
	item.Storefront = new(store.Storefront)
	var priceDuration string
	item.ItemDimension = new(store.ItemDimension)
	var createdAt int64
	var unitOfMeasurement string

	err = row.Scan(&item.Id, &item.Uuid, &item.Storefront.Uuid, &item.Name, &item.Amount, &item.Unit, &item.Price, &priceDuration, &item.Description, &item.Photo, &item.Rating, &item.ItemDimension.Length, &item.ItemDimension.Width, &item.ItemDimension.Height, &item.Color, &unitOfMeasurement, &createdAt)
	errorkit.ErrorHandled(err)

	item.ItemDimension.UnitOfMeasurement = store.UnitofMeasurement(store.UnitofMeasurement_value[unitOfMeasurement])

	createdAtProto, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
	errorkit.ErrorHandled(err)
	item.CreatedAt = createdAtProto

	item.PriceDuration = store.PriceDuration(store.PriceDuration_value[priceDuration])

	return &item
}

func (ri *ReviewItem) initItemReview(usr *user.User, item *store.Item, inEvent *events.ReviewItem) *store.ItemReview {
	return &store.ItemReview{
		CreatedAt: ptypes.TimestampNow(),
		Item:      item,
		Rating:    float64(inEvent.Rating),
		Review:    inEvent.Review,
		UserUuid:  usr.Uuid,
		Uuid:      uuid.New().String()}
}

func (ri *ReviewItem) recalculateItemRating(item *store.Item, latestItemReview *store.ItemReview) {
	row, err := ri.DBO.QueryRow("SELECT COUNT(uuid) FROM kudaki_store.item_reviews WHERE item_uuid=?;", item.Uuid)
	errorkit.ErrorHandled(err)

	var totalItemReview sql.NullInt64
	errorkit.ErrorHandled(row.Scan(&totalItemReview))

	row, err = ri.DBO.QueryRow("SELECT total_raw_rating FROM kudaki_store.items WHERE uuid=?;", item.Uuid)
	errorkit.ErrorHandled(err)

	var totalRawRating sql.NullFloat64
	errorkit.ErrorHandled(row.Scan(&totalRawRating))

	totalRawRating.Float64 += latestItemReview.Rating
	totalItemReview.Int64++

	(*item).Rating = totalRawRating.Float64 / float64(totalItemReview.Int64)
	(*item).TotalRawRating = totalRawRating.Float64
}
*/
