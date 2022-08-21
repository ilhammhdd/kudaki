package externals

import (
	"github.com/ilhammhdd/kudaki-user-info-service/adapters"
	"github.com/ilhammhdd/kudaki-user-info-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases/events"
)

type RetrieveAddresses struct{}

func (ra *RetrieveAddresses) Work() interface{} {
	usecase := usecases.RetrieveAddresses{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}
	ede := EventDrivenExternal{
		eventDrivenAdapter: new(adapters.RetrieveAddresses),
		eventDrivenUsecase: &usecase,
		eventName:          events.UserInfoServiceCommandTopic_RETRIEVE_ADDRESSES.String(),
		inTopics:           []string{events.UserInfoServiceCommandTopic_RETRIEVE_ADDRESSES.String()},
		outTopic:           events.UserInfoServiceEventTopic_ADDRESSES_RETRIEVED.String()}

	ede.handle()
	return nil
}
