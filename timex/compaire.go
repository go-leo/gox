package timex

import "time"

// Max return the maximum time.Time value from a set of times.
func Max(x time.Time, y ...time.Time) time.Time {
	r := x
	for _, t := range y {
		if t.After(r) {
			r = t
		}
	}
	return r
}

// Min return the smallest time.Time value from a set of times.
func Min(x time.Time, y ...time.Time) time.Time {
	r := x
	for _, t := range y {
		if t.Before(r) {
			r = t
		}
	}
	return r
}
