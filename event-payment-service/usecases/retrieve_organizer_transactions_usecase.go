package usecases

import (
	"database/sql"
	"net/http"

	"github.com/ilhammhdd/kudaki-event-payment-service/entities/aggregates/user"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-event-payment-service/entities/aggregates/kudaki_event"
	"github.com/ilhammhdd/kudaki-event-payment-service/usecases/events"
)

type RetrieveOrganizerTransaction struct {
	DBO           DBOperator
	ResultSchemer ResultSchemer
}

func (rot *RetrieveOrganizerTransaction) Handle(in proto.Message) (out proto.Message) {
	_, outEvent := rot.initInOutEvent(in)

	dokuInvoices := rot.retrieveDokuInvoices(outEvent)
	outEvent.Result = rot.ResultSchemer.SetResultSources(dokuInvoices).ParseToResult()
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (rot *RetrieveOrganizerTransaction) initInOutEvent(in proto.Message) (inEvent *events.RetrieveOrganizerInvoices, outEvent *events.OrganizerInvoicesRetrieved) {
	inEvent = in.(*events.RetrieveOrganizerInvoices)

	outEvent = new(events.OrganizerInvoicesRetrieved)
	outEvent.EventStatus = new(events.Status)
	outEvent.Organizer = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.Uid = inEvent.Uid

	return
}

type DokuInvoiceTemp struct {
	kudaki_event.DokuInvoice
	Status      string                   `json:"status"`
	KudakiEvent kudaki_event.KudakiEvent `json:"kudaki_event"`
}

func (rot *RetrieveOrganizerTransaction) retrieveDokuInvoices(outEvent *events.OrganizerInvoicesRetrieved) []*DokuInvoiceTemp {
	var query string
	var rows *sql.Rows
	var err error

	if outEvent.Organizer.Role == user.UserRole_KUDAKI_TEAM || outEvent.Organizer.Role == user.UserRole_ADMIN {
		query = "SELECT di.id, di.uuid, di.kudaki_event_uuid, di.amount, di.purchase_amount, di.transaction_id_merchant, di.request_date_time, di.currency, di.purchase_currency, di.session_id, di.name, di.email, di.basket, di.status, ke.id AS kudaki_event_id, ke.uuid AS kudaki_event_uuid, ke.name AS kudaki_event_name, ke.latitude AS kudaki_event_latitude, ke.longitude AS kudaki_event_longitude, ke.venue AS kudaki_event_venue, ke.description AS kudaki_event_description, ke.ad_duration_from AS kudaki_event_ad_duration_from, ke.ad_duration_to AS kudaki_event_ad_duration_to, ke.duration_from AS kudaki_event_duration_from, ke.duration_to AS kudaki_event_duration_to, ke.seen AS kudaki_event_seen, ke.status AS kudaki_event_status FROM kudaki_event.doku_invoices di JOIN kudaki_event.kudaki_events ke ON di.kudaki_event_uuid = ke.uuid;"
		rows, err = rot.DBO.Query(query)
		errorkit.ErrorHandled(err)
	} else {
		query = "SELECT di.id, di.uuid, di.kudaki_event_uuid, di.amount, di.purchase_amount, di.transaction_id_merchant, di.request_date_time, di.currency, di.purchase_currency, di.session_id, di.name, di.email, di.basket, di.status, ke.id AS kudaki_event_id, ke.uuid AS kudaki_event_uuid, ke.name AS kudaki_event_name, ke.latitude AS kudaki_event_latitude, ke.longitude AS kudaki_event_longitude, ke.venue AS kudaki_event_venue, ke.description AS kudaki_event_description, ke.ad_duration_from AS kudaki_event_ad_duration_from, ke.ad_duration_to AS kudaki_event_ad_duration_to, ke.duration_from AS kudaki_event_duration_from, ke.duration_to AS kudaki_event_duration_to, ke.seen AS kudaki_event_seen, ke.status AS kudaki_event_status FROM kudaki_event.doku_invoices di JOIN kudaki_event.kudaki_events ke ON di.kudaki_event_uuid = ke.uuid WHERE ke.organizer_user_uuid = ?;"
		rows, err = rot.DBO.Query(query, outEvent.Organizer.Uuid)
		errorkit.ErrorHandled(err)
	}

	defer rows.Close()

	var dokuInvoices []*DokuInvoiceTemp
	for rows.Next() {
		var dokuInvoice DokuInvoiceTemp
		rows.Scan(
			&dokuInvoice.Id,
			&dokuInvoice.Uuid,
			&dokuInvoice.KudakiEventUuid,
			&dokuInvoice.Amount,
			&dokuInvoice.PurchaseAmount,
			&dokuInvoice.TransactionIdMerchant,
			&dokuInvoice.RequestDateTime,
			&dokuInvoice.Currency,
			&dokuInvoice.PurchaseCurrency,
			&dokuInvoice.SessionId,
			&dokuInvoice.Name,
			&dokuInvoice.Email,
			&dokuInvoice.Basket,
			&dokuInvoice.Status,
			&dokuInvoice.KudakiEvent.Id,
			&dokuInvoice.KudakiEvent.Uuid,
			&dokuInvoice.KudakiEvent.Name,
			&dokuInvoice.KudakiEvent.Latitude,
			&dokuInvoice.KudakiEvent.Longitude,
			&dokuInvoice.KudakiEvent.Venue,
			&dokuInvoice.KudakiEvent.Description,
			&dokuInvoice.KudakiEvent.AdDurationFrom,
			&dokuInvoice.KudakiEvent.AdDurationTo,
			&dokuInvoice.KudakiEvent.DurationFrom,
			&dokuInvoice.KudakiEvent.DurationTo,
			&dokuInvoice.KudakiEvent.Seen,
			&dokuInvoice.KudakiEvent.Status)

		dokuInvoices = append(dokuInvoices, &dokuInvoice)
	}

	return dokuInvoices
}
