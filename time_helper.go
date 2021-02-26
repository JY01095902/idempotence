package idempotence

import (
	"time"
)

func GetTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}
