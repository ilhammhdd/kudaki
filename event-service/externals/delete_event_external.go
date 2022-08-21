package externals

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-event-service/adapters"
	"github.com/ilhammhdd/kudaki-event-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-event-service/usecases"
	"github.com/ilhammhdd/kudaki-event-service/usecases/events"
)

type DeleteKudakiEvent struct{}

func (ae *DeleteKudakiEvent) Work() interface{} {
	usecase := usecases.DeleteKudakiEvent{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: ae,
		eventDrivenAdapter:  new(adapters.DeleteKudakiEvent),
		eventDrivenUsecase:  &usecase,
		eventName:           events.EventServiceCommandTopic_DELETE_KUDAKI_EVENT.String(),
		inTopics:            []string{events.EventServiceCommandTopic_DELETE_KUDAKI_EVENT.String()},
		outTopic:            events.EventServiceEventTopic_KUDAKI_EVENT_DELETED.String()}

	ede.handle()
	return nil
}

func (ae *DeleteKudakiEvent) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.KudakiEventDeleted)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("DELETE FROM kudaki_event.kudaki_events WHERE uuid=?;",
		out.KudakiEvent.Uuid)
	errorkit.ErrorHandled(err)
}
