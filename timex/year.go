package timex

import "time"

func Year(ts ...time.Time) time.Time {
	if len(ts) <= 0 {
		return year(time.Now())
	}
	if ts[0].IsZero() {
		return year(time.Now())
	}
	return year(ts[0])
}

func year(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}
