package externals_test

import (
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
)

func TestProtoUnix(t *testing.T) {
	customtime := time.Unix(1562825326, 0)
	t.Log("custom time : ", customtime.Unix())
	timeNow := time.Now()
	t.Log("time now : ", timeNow.Unix())

	customProtoTime, err := ptypes.TimestampProto(customtime)
	t.Error(err)
	t.Log("proto time : ", customProtoTime)
	nowProtoTime, err := ptypes.TimestampProto(timeNow)
	t.Error(err)
	t.Log("time now : ", nowProtoTime)
}
