package externals

import (
	"github.com/ilhammhdd/kudaki-rental-service/adapters"
	"github.com/ilhammhdd/kudaki-rental-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-rental-service/usecases"
	"github.com/ilhammhdd/kudaki-rental-service/usecases/events"
)

type RetrieveCartItems struct{}

func (rci *RetrieveCartItems) Work() interface{} {
	adapter := new(adapters.RetrieveCartItems)
	usecase := usecases.RetrieveCartItems{
		DBO:           mysql.NewDBOperation(mysql.QueryDB),
		ResultSchemer: adapter}

	ede := EventDrivenExternal{
		eventDrivenAdapter: adapter,
		eventDrivenUsecase: &usecase,
		eventName:          events.RentalServiceCommandTopic_RETRIEVE_CART_ITEMS.String(),
		inTopics:           []string{events.RentalServiceCommandTopic_RETRIEVE_CART_ITEMS.String()},
		outTopic:           events.RentalServiceEventTopic_CART_ITEMS_RETRIEVED.String()}

	ede.handle()
	return nil
}
