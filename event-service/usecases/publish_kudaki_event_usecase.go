package usecases

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/ilhammhdd/kudaki-event-service/entities/aggregates/kudaki_event"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/kudaki-event-service/usecases/events"
)

type PublishKudakiEvent struct {
	DBO DBOperator
}

func (pke *PublishKudakiEvent) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := pke.initInOutEvent(in)

	kudakiEvent := pke.checkPaymentStatus(inEvent.KudakiEventUuid)
	if kudakiEvent == nil {
		outEvent.EventStatus.Errors = []string{"kudaki event with the given uuid not exists, or the payment is not SUCCESS"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	outEvent.KudakiEvent = kudakiEvent
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (pke *PublishKudakiEvent) initInOutEvent(in proto.Message) (inEvent *events.PublishKudakiEvent, outEvent *events.KudakiEventPublished) {
	inEvent = in.(*events.PublishKudakiEvent)

	outEvent = new(events.KudakiEventPublished)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Publisher = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.Uid = inEvent.Uid

	return
}

func (pke *PublishKudakiEvent) checkPaymentStatus(kudakiEventUuid string) *kudaki_event.KudakiEvent {
	log.Println(kudakiEventUuid, kudaki_event.DokuInvoiceStatus_SUCCESS.String())
	row, err := pke.DBO.QueryRow("SELECT ke.ad_duration_from, ke.ad_duration_to, ke.created_at, ke.description, ke.duration_from, ke.duration_to, ke.file_path, ke.id, ke.name, ke.organizer_user_uuid, ke.seen, ke.status, ke.uuid, ke.venue FROM kudaki_event.doku_invoices di JOIN kudaki_event.kudaki_events ke ON di.kudaki_event_uuid = ke.uuid WHERE ke.uuid = ? AND di.status = ?;",
		kudakiEventUuid, kudaki_event.DokuInvoiceStatus_SUCCESS.String())
	errorkit.ErrorHandled(err)

	var ke kudaki_event.KudakiEvent

	var adDurationFrom int64
	var adDurationTo int64
	var createdAt int64
	var durationFrom int64
	var durationTo int64
	var status string

	if errScan := row.Scan(
		&adDurationFrom,
		&adDurationTo,
		&createdAt,
		&ke.Description,
		&durationFrom,
		&durationTo,
		&ke.FilePath,
		&ke.Id,
		&ke.Name,
		&ke.OrganizerUserUuid,
		&ke.Seen,
		&status,
		&ke.Uuid,
		&ke.Venue); errScan == sql.ErrNoRows {
		return nil
	}

	adDurationFromProto, err := ptypes.TimestampProto(time.Unix(adDurationFrom, 0))
	errorkit.ErrorHandled(err)
	adDurationToProto, err := ptypes.TimestampProto(time.Unix(adDurationTo, 0))
	errorkit.ErrorHandled(err)
	createdAtProto, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
	errorkit.ErrorHandled(err)
	durationFromProto, err := ptypes.TimestampProto(time.Unix(durationFrom, 0))
	errorkit.ErrorHandled(err)
	durationToProto, err := ptypes.TimestampProto(time.Unix(durationTo, 0))
	errorkit.ErrorHandled(err)

	ke.AdDurationFrom = adDurationFromProto
	ke.AdDurationTo = adDurationToProto
	ke.CreatedAt = createdAtProto
	ke.DurationFrom = durationFromProto
	ke.DurationTo = durationToProto
	ke.Status = kudaki_event.KudakiEventStatus_PUBLISHED

	return &ke
}
