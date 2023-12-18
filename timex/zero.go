package timex

import "time"

var UnixZero = time.Unix(0, 0)

func IsZero(t time.Time) bool {
	return t.IsZero() || t.Equal(UnixZero)
}

func IsNotZero(t time.Time) bool {
	return !IsZero(t)
}
