package externals

import (
	"github.com/ilhammhdd/kudaki-order-service/adapters"
	"github.com/ilhammhdd/kudaki-order-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-order-service/usecases"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type RetrieveTenantOrderHistories struct{}

func (rtoh *RetrieveTenantOrderHistories) Work() interface{} {
	adapter := new(adapters.RetrieveTenantOrderHistories)
	usecase := &usecases.RetrieveTenantOrderHistories{
		DBO:           mysql.NewDBOperation(mysql.QueryDB),
		ResultSchemer: adapter}

	ede := EventDrivenExternal{
		eventDrivenAdapter: adapter,
		eventDrivenUsecase: usecase,
		eventName:          events.OrderServiceCommandTopic_RETRIEVE_TENANTS_ORDER_HISTORIES.String(),
		inTopics:           []string{events.OrderServiceCommandTopic_RETRIEVE_TENANTS_ORDER_HISTORIES.String()},
		outTopic:           events.OrderServiceEventTopic_TENANTS_ORDER_HISTORIES_RETRIEVED.String()}

	ede.handle()
	return nil
}
