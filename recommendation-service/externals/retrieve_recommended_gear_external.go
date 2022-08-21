package externals

import (
	"github.com/ilhammhdd/kudaki-recommendation-service/adapters"
	"github.com/ilhammhdd/kudaki-recommendation-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-recommendation-service/usecases"
	"github.com/ilhammhdd/kudaki-recommendation-service/usecases/events"
)

type RetrieveRecommendationGear struct{}

func (rrg *RetrieveRecommendationGear) Work() interface{} {
	adapter := &adapters.RetrieveRecommendationGear{}
	usecase := &usecases.RetrieveRecommendationGear{
		DBO:           mysql.NewDBOperation(mysql.QueryDB),
		ResultSchemer: adapter}

	ede := EventDrivenExternal{
		eventDrivenAdapter: adapter,
		eventDrivenUsecase: usecase,
		eventName:          events.RecommendationServiceCommandTopic_RETRIEVE_RECOMMENDED_GEARS.String(),
		inTopics:           []string{events.RecommendationServiceCommandTopic_RETRIEVE_RECOMMENDED_GEARS.String()},
		outTopic:           events.RecommendationServiceEventTopic_RECOMMENDED_GEARS_RETRIEVED.String()}

	ede.handle()
	return nil
}
