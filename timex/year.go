package timex

import "time"

func Year(t time.Time) time.Time {
	return time.Date(t.Year(), 0, 0, 0, 0, 0, 0, t.Location())
}

func ThisYear(ts ...time.Time) time.Time {
	if len(ts) <= 0 {
		return Year(time.Now())
	}
	if ts[0].IsZero() {
		return Year(time.Now())
	}
	return Year(ts[0])
}
