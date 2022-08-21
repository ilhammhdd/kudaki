package adapters

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-item-service/usecases"
	"github.com/ilhammhdd/kudaki-item-service/usecases/events"
)

type RetrieveItems struct {
	Result ItemsRetrievedSchema
}

func (ri *RetrieveItems) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.RetrieveItems
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (ri *RetrieveItems) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.ItemsRetrieved)

	outByte, err := proto.Marshal(out)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type ItemsRetrievedSchema struct {
	Items []*usecases.ItemTemp `json:"items"`
}

func (ri *RetrieveItems) SetResultSources(i ...interface{}) usecases.ResultSchemer {
	ri.Result = ItemsRetrievedSchema{
		Items: i[0].([]*usecases.ItemTemp)}

	return ri
}

func (ri *RetrieveItems) ParseToResult() []byte {
	resultJSON, err := json.Marshal(ri.Result)
	errorkit.ErrorHandled(err)

	return resultJSON
}
