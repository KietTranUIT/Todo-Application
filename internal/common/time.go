package common

import (
	"time"
)

func GetTimeNow() time.Time {
	return time.Now()
}

func GetExpireTime() time.Time {
	now := time.Now()
	now = now.Add(time.Minute * 10)
	return now
}
