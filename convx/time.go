package convx

import (
	"database/sql/driver"
	"fmt"
	"github.com/go-leo/gox/reflectx"
	"time"
)

// ToTime casts an interface to a time.Time type.
func ToTime(o any) time.Time {
	v, _ := ToTimeE(o)
	return v
}

// ToTimeE casts an interface to a time.Time type.
func ToTimeE(o any) (time.Time, error) {
	return ToTimeInLocationE(o, time.UTC)
}

// ToTimeInLocation casts an empty interface to time.Time,
func ToTimeInLocation(o any, location *time.Location) time.Time {
	v, _ := ToTimeInLocationE(o, location)
	return v
}

// ToTimeInLocationE casts an empty interface to time.Time,
// interpreting inputs without a timezone to be in the given location,
// or the local timezone if nil.
func ToTimeInLocationE(o any, location *time.Location) (time.Time, error) {
	return toTimeInLocationE(o, location)
}

func toTimeInLocationE(o any, location *time.Location) (time.Time, error) {
	zero := time.Time{}
	o = reflectx.Indirect(o)
	switch t := o.(type) {
	case time.Time:
		return t, nil
	case int, int64, int32, int16, int8,
		uint, uint64, uint32, uint16, uint8,
		float32, float64,
		interface{ Int64() (int64, error) }, interface{ Float64() (float64, error) }: // json.Number
		v, err := ToInt64E(t)
		if err != nil {
			return zero, err
		}
		return time.Unix(v, 0), nil
	case string:
		for _, format := range timeFormats {
			tim, err := time.ParseInLocation(format, t, location)
			if err != nil {
				continue
			}
			return tim, nil
		}
		return zero, fmt.Errorf(failedCast, o, o, zero)
	case driver.Valuer:
		v, err := t.Value()
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return ToTimeInLocationE(v, location)
	case interface{ AsTime() time.Time }:
		return t.AsTime(), nil
	default:
		return zero, fmt.Errorf("unable to cast %#s of type %T to Time", o, o)
	}
}

var (
	timeFormats = []string{
		time.Layout,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		// Handy time stamps.
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
		time.DateTime,
		time.DateOnly,
		time.TimeOnly,

		"2006-01-02 15:04:05Z07:00",
		"02 Jan 2006",
		"2006-01-02 15:04:05 -07:00",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02T15:04:05",                     // iso8601 without timezone
		"2006-01-02 15:04:05.999999999 -0700 MST", // Time.String()
		"2006-01-02T15:04:05-0700",                // RFC3339 without timezone hh:mm colon
		"2006-01-02 15:04:05Z0700",                // RFC3339 without T or timezone hh:mm colon

	}
)
