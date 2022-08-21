package usecases

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-event-service/entities/aggregates/kudaki_event"
	"github.com/ilhammhdd/kudaki-event-service/usecases/events"
)

type RetrieveKudakiEvent struct {
	DBO DBOperator
}

func (rke *RetrieveKudakiEvent) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := rke.initInOutEvent(in)

	kudakiEvent, paymentStat := rke.retrieveFromDB(inEvent)
	if kudakiEvent == nil {
		outEvent.EventStatus.Errors = []string{"kudaki event with the given uuid doesn't exists"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	outEvent.KudakiEvent = kudakiEvent
	outEvent.PaymentStatus = paymentStat
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (rke *RetrieveKudakiEvent) initInOutEvent(in proto.Message) (inEvent *events.RetrieveKudakiEvent, outEvent *events.KudakiEventRetrieved) {
	inEvent = in.(*events.RetrieveKudakiEvent)

	outEvent = new(events.KudakiEventRetrieved)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Uid = inEvent.Uid

	return
}

func (rke *RetrieveKudakiEvent) retrieveFromDB(inEvent *events.RetrieveKudakiEvent) (*kudaki_event.KudakiEvent, string) {
	row, err := rke.DBO.QueryRow("SELECT ke.uuid, ke.seen, ke.name, ke.venue, ke.description, ke.duration_from, ke.duration_to, ke.ad_duration_from, ke.ad_duration_to, ke.status, ke.file_path, di.status FROM kudaki_event.kudaki_events ke JOIN kudaki_event.doku_invoices di ON ke.uuid = di.kudaki_event_uuid WHERE ke.uuid = ?;",
		inEvent.KudakiEventUuid)
	errorkit.ErrorHandled(err)

	var kudakiEvent kudaki_event.KudakiEvent
	var durationFrom int64
	var durationTo int64
	var adDurationFrom int64
	var adDurationTo int64
	var status string
	var paymentStatus string

	if err = row.Scan(
		&kudakiEvent.Uuid,
		&kudakiEvent.Seen,
		&kudakiEvent.Name,
		&kudakiEvent.Venue,
		&kudakiEvent.Description,
		&durationFrom,
		&durationTo,
		&adDurationFrom,
		&adDurationTo,
		&status,
		&kudakiEvent.FilePath,
		&paymentStatus); err == sql.ErrNoRows {
		return nil, ""
	}

	durationFromProto, err := ptypes.TimestampProto(time.Unix(durationFrom, 0))
	errorkit.ErrorHandled(err)
	durationToProto, err := ptypes.TimestampProto(time.Unix(durationTo, 0))
	errorkit.ErrorHandled(err)
	adDurationFromProto, err := ptypes.TimestampProto(time.Unix(adDurationFrom, 0))
	errorkit.ErrorHandled(err)
	adDurationToProto, err := ptypes.TimestampProto(time.Unix(adDurationTo, 0))
	errorkit.ErrorHandled(err)

	kudakiEvent.DurationFrom = durationFromProto
	kudakiEvent.DurationTo = durationToProto
	kudakiEvent.AdDurationFrom = adDurationFromProto
	kudakiEvent.AdDurationTo = adDurationToProto
	kudakiEvent.Status = kudaki_event.KudakiEventStatus(kudaki_event.KudakiEventStatus_value[status])
	kudakiEvent.Seen++
	kudakiEvent.Uuid = inEvent.KudakiEventUuid

	return &kudakiEvent, paymentStatus
}
