package externals

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-recommendation-service/adapters"
	"github.com/ilhammhdd/kudaki-recommendation-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-recommendation-service/usecases"
	"github.com/ilhammhdd/kudaki-recommendation-service/usecases/events"
)

type AddRecommendationGear struct{}

func (arg *AddRecommendationGear) Work() interface{} {
	usecase := &usecases.AddRecommendationGear{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: arg,
		eventDrivenAdapter:  new(adapters.AddRecommendationGear),
		eventDrivenUsecase:  usecase,
		eventName:           events.RecommendationServiceCommandTopic_ADD_RECOMMENDED_GEAR.String(),
		inTopics:            []string{events.RecommendationServiceCommandTopic_ADD_RECOMMENDED_GEAR.String()},
		outTopic:            events.RecommendationServiceEventTopic_RECOMMENDED_GEAR_ADDED.String()}

	ede.handle()
	return nil
}

func (arg *AddRecommendationGear) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.RecommendedGearAdded)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("INSERT INTO kudaki_mountain.recommended_gears(uuid,user_uuid,mountain_uuid,upvote,downvote,seen,created_at) VALUES(?,?,?,?,?,?,UNIX_TIMESTAMP());",
		out.RecommendedGear.Uuid,
		out.RecommendedGear.UserUuid,
		out.RecommendedGear.Mountain.Uuid,
		out.RecommendedGear.Upvote,
		out.RecommendedGear.Downvote,
		out.RecommendedGear.Seen)
	errorkit.ErrorHandled(err)

	dboItems := mysql.NewDBOperation(mysql.CommandDB)
	for i := 0; i < len(out.RecommendedGearItems); i++ {
		_, err = dboItems.Command("INSERT INTO kudaki_mountain.recommended_gear_items(uuid,recommended_gear_uuid,item_type,total,created_at) VALUES(?,?,?,?,UNIX_TIMESTAMP());",
			out.RecommendedGearItems[i].Uuid,
			out.RecommendedGearItems[i].RecommendedGear.Uuid,
			out.RecommendedGearItems[i].ItemType,
			out.RecommendedGearItems[i].Total)
		errorkit.ErrorHandled(err)
	}
}
