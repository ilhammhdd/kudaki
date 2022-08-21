package externals

import (
	"github.com/ilhammhdd/kudaki-user-info-service/adapters"
	"github.com/ilhammhdd/kudaki-user-info-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases/events"
)

type RetrieveProfile struct{}

func (rp *RetrieveProfile) Work() interface{} {
	adapter := new(adapters.RetrieveProfile)
	usecase := usecases.RetrieveProfile{
		DBO:           mysql.NewDBOperation(mysql.QueryDB),
		ResultSchemer: adapter}

	ede := EventDrivenExternal{
		eventDrivenAdapter: adapter,
		eventDrivenUsecase: &usecase,
		eventName:          events.UserInfoServiceCommandTopic_RETRIEVE_PROFILE.String(),
		inTopics:           []string{events.UserInfoServiceCommandTopic_RETRIEVE_PROFILE.String()},
		outTopic:           events.UserInfoServiceEventTopic_PROFILE_RETRIEVED.String()}

	ede.handle()
	return nil
}
