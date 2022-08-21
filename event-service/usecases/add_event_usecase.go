package usecases

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/ilhammhdd/kudaki-event-service/entities/aggregates/user"

	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/ilhammhdd/kudaki-event-service/entities/aggregates/kudaki_event"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/kudaki-event-service/usecases/events"
)

type AddKudakiEvent struct {
	DBO DBOperator
}

func (ae *AddKudakiEvent) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := ae.initInOutEvent(in)

	outEvent.KudakiEvent = ae.initKudakiEvent(inEvent, outEvent)
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (ae *AddKudakiEvent) initInOutEvent(in proto.Message) (inEvent *events.AddKudakiEvent, outEvent *events.KudakiEventAdded) {
	inEvent = in.(*events.AddKudakiEvent)

	outEvent = new(events.KudakiEventAdded)
	outEvent.AddKudakiEvent = inEvent
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Organizer = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.Uid = inEvent.Uid

	return
}

func (ae *AddKudakiEvent) initKudakiEvent(inEvent *events.AddKudakiEvent, outEvent *events.KudakiEventAdded) *kudaki_event.KudakiEvent {
	durationFromProto, err := ptypes.TimestampProto(time.Unix(inEvent.DurationFrom, 0))
	errorkit.ErrorHandled(err)
	durationToProto, err := ptypes.TimestampProto(time.Unix(inEvent.DurationTo, 0))
	errorkit.ErrorHandled(err)

	adDurationFromProto, err := ptypes.TimestampProto(time.Unix(inEvent.AdDurationFrom, 0))
	errorkit.ErrorHandled(err)
	adDurationToProto, err := ptypes.TimestampProto(time.Unix(inEvent.AdDurationTo, 0))
	errorkit.ErrorHandled(err)

	return &kudaki_event.KudakiEvent{
		AdDurationFrom:    adDurationFromProto,
		AdDurationTo:      adDurationToProto,
		CreatedAt:         ptypes.TimestampNow(),
		Description:       inEvent.Description,
		DurationFrom:      durationFromProto,
		DurationTo:        durationToProto,
		FilePath:          inEvent.FilePath,
		Name:              inEvent.Name,
		OrganizerUserUuid: outEvent.Organizer.Uuid,
		Seen:              0,
		Status:            kudaki_event.KudakiEventStatus_UNPUBLISHED,
		Uuid:              uuid.New().String(),
		Venue:             inEvent.Venue}
}

func (ae *AddKudakiEvent) retrieveProfile(usr *user.User) *user.Profile {
	row, err := ae.DBO.QueryRow("SELECT id,uuid,full_name,photo,created_at FROM kudaki_user.profiles WHERE user_uuid = ?;", usr.Uuid)
	errorkit.ErrorHandled(err)

	var profile user.Profile
	var createdAt sql.NullInt64
	errorkit.ErrorHandled(row.Scan(&profile.Id, &profile.Uuid, &profile.FullName, &profile.Photo, &createdAt))
	createdAtProtoTime, err := ptypes.TimestampProto(time.Unix(createdAt.Int64, 0))
	errorkit.ErrorHandled(err)
	profile.CreatedAt = createdAtProtoTime
	profile.User = usr

	return &profile
}
