package adapters

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-mountain-service/usecases"
	"github.com/ilhammhdd/kudaki-mountain-service/usecases/events"
)

type RetrieveMountains struct {
	Result RetrieveMountainsResultSchema
}

func (rm *RetrieveMountains) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.RetrieveMountains
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (rm *RetrieveMountains) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.MountainsRetrieved)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type RetrieveMountainsResultSchema struct {
	Mountains []*usecases.MountainTemp `json:"mountains"`
}

func (rm *RetrieveMountains) SetResultSources(i ...interface{}) usecases.ResultSchemer {
	rm.Result = RetrieveMountainsResultSchema{
		Mountains: i[0].([]*usecases.MountainTemp)}

	return rm
}

func (rm *RetrieveMountains) ParseToResult() []byte {
	resultJSON, err := json.Marshal(rm.Result)
	errorkit.ErrorHandled(err)

	return resultJSON
}
