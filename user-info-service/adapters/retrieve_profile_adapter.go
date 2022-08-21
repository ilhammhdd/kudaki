package adapters

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases/events"
)

type RetrieveProfile struct {
	Result RetrieveProfileResultSchema
}

func (rp *RetrieveProfile) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.RetrieveProfile
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (rp *RetrieveProfile) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.ProfileRetrieved)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type RetrieveProfileResultSchema struct {
	Profile *usecases.RetrieveProfileTemp `json:"profile"`
}

func (rp *RetrieveProfile) SetResultSources(i ...interface{}) usecases.ResultSchemer {
	rp.Result = RetrieveProfileResultSchema{
		Profile: i[0].(*usecases.RetrieveProfileTemp)}

	return rp
}

func (rp *RetrieveProfile) ParseToResult() []byte {
	resultJSON, err := json.Marshal(rp.Result)
	errorkit.ErrorHandled(err)

	return resultJSON
}
