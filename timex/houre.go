package timex

import "time"

func Hour(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
}

func CurrentHour() time.Time {
	return Hour(time.Now())
}

// ThisHour
func ThisHour(ts ...time.Time) time.Time {
	if len(ts) <= 0 {
		return CurrentMonth()
	}
	return Hour(ts[0])
}

// LastMonth
func LastHour(ts ...time.Time) time.Time {
	return ThisHour(ts...).Add(-time.Hour)
}

// NextMonth
func NextHour(ts ...time.Time) time.Time {
	return ThisMonth(ts...).Add(time.Hour)
}
