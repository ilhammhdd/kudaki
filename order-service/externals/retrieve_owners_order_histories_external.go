package externals

import (
	"github.com/ilhammhdd/kudaki-order-service/adapters"
	"github.com/ilhammhdd/kudaki-order-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-order-service/usecases"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type RetrieveOwnersOrderHistories struct{}

func (rooh *RetrieveOwnersOrderHistories) Work() interface{} {
	adapter := new(adapters.RetrieveOwnersOrderHistories)
	usecase := &usecases.RetrieveOwnersOrderHistories{
		DBO:           mysql.NewDBOperation(mysql.QueryDB),
		ResultSchemer: adapter}

	ede := EventDrivenExternal{
		eventDrivenAdapter: adapter,
		eventDrivenUsecase: usecase,
		eventName:          events.OrderServiceCommandTopic_RETRIEVE_OWNERS_ORDER_HISTORIES.String(),
		inTopics:           []string{events.OrderServiceCommandTopic_RETRIEVE_OWNERS_ORDER_HISTORIES.String()},
		outTopic:           events.OrderServiceEventTopic_OWNERS_ORDER_HISTORIES_RETRIEVED.String()}

	ede.handle()
	return nil
}
