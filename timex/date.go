package timex

import "time"

func Date(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func Today() time.Time {
	return Date(time.Now())
}

func Tomorrow() time.Time {
	return Today().AddDate(0, 0, 1)
}

func Yesterday() time.Time {
	return Today().AddDate(0, 0, -1)
}
