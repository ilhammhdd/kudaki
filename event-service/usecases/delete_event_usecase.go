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

type DeleteKudakiEvent struct {
	DBO DBOperator
}

func (dke *DeleteKudakiEvent) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := dke.initInOutEvent(in)

	kudakiEvent := dke.checkEventOwnership(outEvent.Organizer.Uuid, inEvent.EventUuid)
	if kudakiEvent == nil {
		outEvent.EventStatus.Errors = []string{"organizer's event with the given uuid not found"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	outEvent.KudakiEvent = kudakiEvent
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (dke *DeleteKudakiEvent) initInOutEvent(in proto.Message) (inEvent *events.DeleteKudakiEvent, outEvent *events.KudakiEventDeleted) {
	inEvent = in.(*events.DeleteKudakiEvent)

	outEvent = new(events.KudakiEventDeleted)
	outEvent.DeleteKudakiEvent = inEvent
	outEvent.EventStatus = new(events.Status)
	outEvent.Organizer = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.Uid = inEvent.Uid

	return
}

func (dke *DeleteKudakiEvent) checkEventOwnership(organizerUuid, eventUuid string) *kudaki_event.KudakiEvent {
	row, err := dke.DBO.QueryRow("SELECT id,uuid,organizer_user_uuid,name,venue,description,duration_from,duration_to,seen,status,created_at FROM kudaki_event.kudaki_events WHERE uuid=? AND organizer_user_uuid=?;",
		eventUuid, organizerUuid)
	errorkit.ErrorHandled(err)

	var kudakiEvent kudaki_event.KudakiEvent
	var durationFrom, durationTo, createdAt int64
	var status string

	if row.Scan(
		&kudakiEvent.Id,
		&kudakiEvent.Uuid,
		&kudakiEvent.OrganizerUserUuid,
		&kudakiEvent.Name,
		&kudakiEvent.Venue,
		&kudakiEvent.Description,
		&durationFrom,
		&durationTo,
		&kudakiEvent.Seen,
		&status,
		&createdAt) == sql.ErrNoRows {
		return nil
	}

	durationFromProto, err := ptypes.TimestampProto(time.Unix(durationFrom, 0))
	errorkit.ErrorHandled(err)
	durationToProto, err := ptypes.TimestampProto(time.Unix(durationTo, 0))
	errorkit.ErrorHandled(err)
	createdAtProto, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
	errorkit.ErrorHandled(err)

	kudakiEvent.DurationFrom = durationFromProto
	kudakiEvent.DurationTo = durationToProto
	kudakiEvent.CreatedAt = createdAtProto
	kudakiEvent.Status = kudaki_event.KudakiEventStatus(kudaki_event.KudakiEventStatus_value[status])

	return &kudakiEvent
}
