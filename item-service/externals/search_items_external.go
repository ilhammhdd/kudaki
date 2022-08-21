package externals

import (
	"github.com/ilhammhdd/kudaki-item-service/adapters"
	"github.com/ilhammhdd/kudaki-item-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-item-service/usecases"
	"github.com/ilhammhdd/kudaki-item-service/usecases/events"
)

type SearchItems struct{}

func (si *SearchItems) Work() interface{} {
	adapter := &adapters.SearchItems{}
	usecase := &usecases.SearchItems{
		DBO:           mysql.NewDBOperation(mysql.QueryDB),
		ResultSchemer: adapter}

	ede := EventDrivenExternal{
		eventDrivenAdapter: adapter,
		eventDrivenUsecase: usecase,
		eventName:          events.ItemServiceCommandTopic_SEARCH_ITEMS.String(),
		inTopics:           []string{events.ItemServiceCommandTopic_SEARCH_ITEMS.String()},
		outTopic:           events.ItemServiceEventTopic_ITEMS_SEARCHED.String()}

	ede.handle()
	return nil
}
