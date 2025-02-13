package timex

import "time"

var UnixZero = time.Unix(0, 0)

func IsUnixZero(t time.Time) bool {
	return t.Equal(UnixZero)
}

func IsZero(t time.Time) bool {
	return t.IsZero() || IsUnixZero(t)
}

func IsNotZero(t time.Time) bool {
	return !IsZero(t)
}
