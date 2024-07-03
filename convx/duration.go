package convx

import (
	"database/sql/driver"
	"fmt"
	"github.com/go-leo/gox/reflectx"
	"reflect"
	"time"
)

// ToDuration casts an interface to a time.Duration type.
func ToDuration(o any) time.Duration {
	v, _ := ToDurationE(o)
	return v
}

// ToDurationE casts an interface to a time.Duration type.
func ToDurationE(o any) (time.Duration, error) {
	return toDurationE(o)
}

// ToDurationSlice casts an interface to a []time.Duration type.
func ToDurationSlice(o any) []time.Duration {
	v, _ := ToDurationSliceE(o)
	return v
}

// ToDurationSliceE casts an interface to a []time.Duration type.
func ToDurationSliceE(o any) ([]time.Duration, error) {
	return toSliceE[[]time.Duration](o, ToDurationE)
}

func toDurationE(o any) (time.Duration, error) {
	var zero time.Duration
	o = reflectx.IndirectToInterface(o,
		reflect.TypeOf((*interface{ Int64() (int64, error) })(nil)).Elem(),
		reflect.TypeOf((*interface{ Float64() (float64, error) })(nil)).Elem(),
		reflect.TypeOf((*interface{ AsDuration() time.Duration })(nil)).Elem(),
		reflect.TypeOf((*driver.Valuer)(nil)).Elem(),
	)
	switch d := o.(type) {
	case time.Duration:
		return d, nil
	case string:
		v, err := time.ParseDuration(d)
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return v, nil
	case int, int64, int32, int16, int8,
		uint, uint64, uint32, uint16, uint8,
		float32, float64,
		interface{ Int64() (int64, error) }, interface{ Float64() (float64, error) }: // json.Number
		v, err := ToInt64E(o)
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return time.Duration(v), nil
	case interface{ AsDuration() time.Duration }:
		return d.AsDuration(), nil
	case driver.Valuer:
		v, err := d.Value()
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return toDurationE(v)
	default:
		return zero, fmt.Errorf(failedCast, o, o, zero)
	}
}
