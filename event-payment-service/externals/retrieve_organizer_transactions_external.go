package externals

import (
	"github.com/ilhammhdd/kudaki-event-payment-service/adapters"
	"github.com/ilhammhdd/kudaki-event-payment-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-event-payment-service/usecases"
	"github.com/ilhammhdd/kudaki-event-payment-service/usecases/events"
)

type RetrieveOrganizerTransaction struct{}

func (rot *RetrieveOrganizerTransaction) Work() interface{} {
	adapter := new(adapters.RetrieveOrganizerTransaction)
	usecase := &usecases.RetrieveOrganizerTransaction{
		DBO:           mysql.NewDBOperation(mysql.QueryDB),
		ResultSchemer: adapter}

	ede := EventDrivenExternal{
		eventDrivenAdapter: adapter,
		eventDrivenUsecase: usecase,
		eventName:          events.EventServiceCommandTopic_RETRIEVE_ORGANIZER_INVOICES.String(),
		inTopics:           []string{events.EventServiceCommandTopic_RETRIEVE_ORGANIZER_INVOICES.String()},
		outTopic:           events.EventServiceEventTopic_ORGANIZER_INVOICES_RETRIEVED.String()}

	ede.handle()
	return nil
}
