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

type RetrieveKudakiEvent struct{}

func (rke *RetrieveKudakiEvent) Work() interface{} {
	adapter := &adapters.RetrieveKudakiEvent{}
	usecase := &usecases.RetrieveKudakiEvent{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: rke,
		eventDrivenAdapter:  adapter,
		eventDrivenUsecase:  usecase,
		eventName:           events.EventServiceCommandTopic_RETRIEVE_KUDAKI_EVENT.String(),
		inTopics:            []string{events.EventServiceCommandTopic_RETRIEVE_KUDAKI_EVENT.String()},
		outTopic:            events.EventServiceEventTopic_KUDAKI_EVENT_RETRIEVED.String()}

	ede.handle()
	return nil
}

func (rke *RetrieveKudakiEvent) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.KudakiEventRetrieved)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("UPDATE kudaki_event.kudaki_events SET seen = ? WHERE uuid = ?;",
		out.KudakiEvent.Seen, out.KudakiEvent.Uuid)
	errorkit.ErrorHandled(err)
}
