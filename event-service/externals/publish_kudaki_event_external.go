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

type PublishKudakiEvent struct{}

func (pke *PublishKudakiEvent) Work() interface{} {
	adapter := &adapters.PublishKudakiEvent{}
	usecase := &usecases.PublishKudakiEvent{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: pke,
		eventDrivenAdapter:  adapter,
		eventDrivenUsecase:  usecase,
		eventName:           events.EventServiceCommandTopic_PUBLISH_KUDAKI_EVENT.String(),
		inTopics:            []string{events.EventServiceCommandTopic_PUBLISH_KUDAKI_EVENT.String()},
		outTopic:            events.EventServiceEventTopic_KUDAKI_EVENT_PUBLISHED.String()}

	ede.handle()
	return nil
}

func (pke *PublishKudakiEvent) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.KudakiEventPublished)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("UPDATE kudaki_event.kudaki_events SET status = ? WHERE uuid = ?;",
		out.KudakiEvent.Status.String(), out.KudakiEvent.Uuid)
	errorkit.ErrorHandled(err)
}
