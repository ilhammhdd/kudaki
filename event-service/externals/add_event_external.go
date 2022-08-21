package externals

import (
	"net/http"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-event-service/adapters"
	"github.com/ilhammhdd/kudaki-event-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-event-service/usecases"
	"github.com/ilhammhdd/kudaki-event-service/usecases/events"
)

type AddKudakiEvent struct{}

func (ae *AddKudakiEvent) Work() interface{} {
	usecase := usecases.AddKudakiEvent{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: ae,
		eventDrivenAdapter:  new(adapters.AddKudakiEvent),
		eventDrivenUsecase:  &usecase,
		eventName:           events.EventServiceCommandTopic_ADD_KUDAKI_EVENT.String(),
		inTopics:            []string{events.EventServiceCommandTopic_ADD_KUDAKI_EVENT.String()},
		outTopic:            events.EventServiceEventTopic_KUDAKI_EVENT_ADDED.String()}

	ede.handle()
	return nil
}

func (ae *AddKudakiEvent) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.KudakiEventAdded)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("INSERT INTO kudaki_event.kudaki_events(uuid,organizer_user_uuid,name,venue,description,duration_from,duration_to,ad_duration_from,ad_duration_to,seen,status,file_path,created_at) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,UNIX_TIMESTAMP());",
		out.KudakiEvent.Uuid,
		out.KudakiEvent.OrganizerUserUuid,
		out.KudakiEvent.Name,
		out.KudakiEvent.Venue,
		out.KudakiEvent.Description,
		out.KudakiEvent.DurationFrom.Seconds,
		out.KudakiEvent.DurationTo.Seconds,
		out.KudakiEvent.AdDurationFrom.Seconds,
		out.KudakiEvent.AdDurationTo.Seconds,
		out.KudakiEvent.Seen,
		out.KudakiEvent.Status.String(),
		out.KudakiEvent.FilePath,
	)
	errorkit.ErrorHandled(err)
}
