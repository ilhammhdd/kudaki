package externals

import (
	"github.com/ilhammhdd/kudaki-item-service/adapters"
	"github.com/ilhammhdd/kudaki-item-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-item-service/usecases"
	"github.com/ilhammhdd/kudaki-item-service/usecases/events"
)

type RetrieveItems struct{}

func (ri *RetrieveItems) Work() interface{} {
	adapter := &adapters.RetrieveItems{}
	usecase := &usecases.RetrieveItems{
		DBO:           mysql.NewDBOperation(mysql.QueryDB),
		ResultSchemer: adapter}

	ede := EventDrivenExternal{
		eventDrivenAdapter: adapter,
		eventDrivenUsecase: usecase,
		eventName:          events.ItemServiceCommandTopic_RETRIEVE_ITEMS.String(),
		inTopics:           []string{events.ItemServiceCommandTopic_RETRIEVE_ITEMS.String()},
		outTopic:           events.ItemServiceEventTopic_ITEMS_RETRIEVED.String()}

	ede.handle()
	return nil
}
