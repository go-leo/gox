package timex

import "time"

func Max(x time.Time, y ...time.Time) time.Time {
	r := x
	for _, t := range y {
		if t.After(r) {
			r = t
		}
	}
	return r
}

func Min(x time.Time, y ...time.Time) time.Time {
	r := x
	for _, t := range y {
		if t.Before(r) {
			r = t
		}
	}
	return r
}
