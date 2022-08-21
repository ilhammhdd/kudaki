package externals

import (
	"github.com/ilhammhdd/kudaki-mountain-service/adapters"
	"github.com/ilhammhdd/kudaki-mountain-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-mountain-service/usecases"
	"github.com/ilhammhdd/kudaki-mountain-service/usecases/events"
)

type RetrieveMountains struct{}

func (rm *RetrieveMountains) Work() interface{} {
	adapter := new(adapters.RetrieveMountains)
	usecase := &usecases.RetrieveMountains{
		DBO:           mysql.NewDBOperation(mysql.QueryDB),
		ResultSchemer: adapter}

	ede := EventDrivenExternal{
		eventDrivenAdapter: adapter,
		eventDrivenUsecase: usecase,
		eventName:          events.MountainServiceCommandTopic_RETRIEVE_MOUNTAINS.String(),
		inTopics:           []string{events.MountainServiceCommandTopic_RETRIEVE_MOUNTAINS.String()},
		outTopic:           events.MountainServiceEventTopic_MOUNTAINS_RETRIEVED.String()}

	ede.handle()
	return nil
}
