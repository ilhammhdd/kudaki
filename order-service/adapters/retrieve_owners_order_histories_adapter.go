package adapters

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/usecases"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type RetrieveOwnersOrderHistories struct {
	Result OwnerOrderResultSchema
}

func (rooh *RetrieveOwnersOrderHistories) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.RetrieveOwnersOrderHistories
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (rooh *RetrieveOwnersOrderHistories) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.OwnersOrderHistoriesRetrieved)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type OwnerOrderResultSchema struct {
	Orders []*usecases.OwnerOrderTemp `json:"owner_orders"`
}

func (rooh *RetrieveOwnersOrderHistories) SetResultSources(i ...interface{}) usecases.ResultSchemer {
	rooh.Result = OwnerOrderResultSchema{
		Orders: i[0].([]*usecases.OwnerOrderTemp)}

	return rooh
}

func (rooh *RetrieveOwnersOrderHistories) ParseToResult() []byte {
	resultJSON, err := json.Marshal(rooh.Result)
	errorkit.ErrorHandled(err)

	return resultJSON
}
