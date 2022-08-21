package usecases

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-item-service/entities/aggregates/store"
	"github.com/ilhammhdd/kudaki-item-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-item-service/usecases/events"
)

type SearchItems struct {
	DBO           DBOperator
	ResultSchemer ResultSchemer
}

func (si *SearchItems) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := si.initInOutEvent(in)

	items := si.retrieveItems(outEvent.Requester, inEvent)

	outEvent.Result = si.ResultSchemer.SetResultSources(items).ParseToResult()
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (si *SearchItems) initInOutEvent(in proto.Message) (inEvent *events.SearchItems, outEvent *events.ItemsSearched) {
	inEvent = in.(*events.SearchItems)

	outEvent = new(events.ItemsSearched)
	outEvent.EventStatus = new(events.Status)
	outEvent.Requester = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.SearchItems = inEvent
	outEvent.Uid = inEvent.Uid

	return
}

func (si *SearchItems) retrieveItems(usr *user.User, inEvent *events.SearchItems) []*ItemTemp {
	rows, err := si.DBO.Query("SELECT i.id,i.uuid,i.storefront_uuid,i.name,i.amount,i.unit,i.price,i.price_duration,i.description,i.photo,i.rating,i.length,i.width,i.height,i.color,i.unit_of_measurement,i.created_at FROM (SELECT i_i.id FROM kudaki_store.items i_i WHERE MATCH(i_i.name,i_i.description) AGAINST(?) LIMIT ?, ?) i_ids JOIN kudaki_store.items i ON i.id = i_ids.id;", inEvent.Keyword, inEvent.Offset, inEvent.Limit)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var items []*ItemTemp
	for rows.Next() {
		var item ItemTemp
		var priceDuration string

		rows.Scan(&item.Id, &item.Uuid, &item.StorefrontUuidT, &item.Name, &item.Amount, &item.Unit, &item.Price, &priceDuration, &item.Description, &item.Photo, &item.Rating, &item.LengthT, &item.WidthT, &item.HeightT, &item.Color, &item.UnitofMeasurementT, &item.CreatedAtT)

		item.PriceDuration = store.PriceDuration(store.PriceDuration_value[priceDuration])

		items = append(items, &item)
	}
	return items
}
