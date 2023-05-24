package clock

import "time"

type Clock interface {
	NowAsUTC() time.Time
}

type RealClock struct{}

func (RealClock) NowAsUTC() time.Time {
	return time.Now().UTC()
}