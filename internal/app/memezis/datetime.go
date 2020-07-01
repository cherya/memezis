package memezis

import (
	"time"

	"github.com/gogo/protobuf/types"
)

func fromProtoTime(timestamp *types.Timestamp) time.Time {
	t, _ := types.TimestampFromProto(timestamp)
	return t
}

func toProtoTime(time time.Time) *types.Timestamp {
	t, _ := types.TimestampProto(time)
	return t
}
