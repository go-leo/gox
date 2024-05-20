package timex

import (
	"time"
)

func Month(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func CurrentMonth() time.Time {
	return Month(time.Now())
}

func ThisMonth(ts ...time.Time) time.Time {
	if len(ts) <= 0 {
		return CurrentMonth()
	}
	return Month(ts[0])
}
