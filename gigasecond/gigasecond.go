package gigasecond

import (
	"time"
)

const gigaSecond = 1000000000

func AddGigasecond(t time.Time) time.Time {
	duration := time.Second * gigaSecond
	return t.Add(duration)
}
