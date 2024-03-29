package timex

import "time"

func Month(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func ThisMonth(ts ...time.Time) time.Time {
	if len(ts) <= 0 {
		return Month(time.Now())
	}
	if ts[0].IsZero() {
		return Month(time.Now())
	}
	return Month(ts[0])
}
