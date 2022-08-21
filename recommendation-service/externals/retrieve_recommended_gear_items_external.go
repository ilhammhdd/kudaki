package externals

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-recommendation-service/adapters"
	"github.com/ilhammhdd/kudaki-recommendation-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-recommendation-service/usecases"
	"github.com/ilhammhdd/kudaki-recommendation-service/usecases/events"
)

type RetrieveRecommendedGearItems struct{}

func (rrgi *RetrieveRecommendedGearItems) Work() interface{} {
	adapter := new(adapters.RetrieveRecommendedGearItems)
	usecase := &usecases.RetrieveRecommendedGearItems{
		DBO:           mysql.NewDBOperation(mysql.QueryDB),
		ResultSchemer: adapter}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: rrgi,
		eventDrivenAdapter:  adapter,
		eventDrivenUsecase:  usecase,
		eventName:           events.RecommendationServiceCommandTopic_RETRIEVE_RECOMMENDED_GEAR_ITEMS.String(),
		inTopics:            []string{events.RecommendationServiceCommandTopic_RETRIEVE_RECOMMENDED_GEAR_ITEMS.String()},
		outTopic:            events.RecommendationServiceEventTopic_RECOMMENDED_GEAR_ITEMS_RETRIEVED.String()}

	ede.handle()
	return nil
}

func (rrgi *RetrieveRecommendedGearItems) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.RecommendedGearItemsRetrieved)

	var rg adapters.RetrieveRecommendedGearItemsResult
	err := json.Unmarshal(out.Result, &rg)
	errorkit.ErrorHandled(err)

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err = dbo.Command("UPDATE kudaki_mountain.recommended_gear rg SET seen = ? WHERE uuid = ?;",
		rg.RecommendedGear.Seen, rg.RecommendedGear.Uuid)
	errorkit.ErrorHandled(err)
}
