package convx

import (
	"database/sql/driver"
	"fmt"
	"github.com/go-leo/gox/reflectx"
	"strconv"
	"time"
)

// ToBool casts an interface to a bool type.
func ToBool[E ~bool](o any) E {
	v, _ := ToBoolE[E](o)
	return v
}

// ToBoolE casts an interface to a bool type.
func ToBoolE[E ~bool](o any) (E, error) {
	return toBoolE[E](o)
}

// ToBoolSlice casts an interface to a []bool type.
func ToBoolSlice[S ~[]E, E ~bool](o any) S {
	v, _ := ToBoolSliceE[S](o)
	return v
}

// ToBoolSliceE casts an interface to a []bool type.
func ToBoolSliceE[S ~[]E, E ~bool](o any) (S, error) {
	return toSliceE[S](o, toBoolE[E])
}

func toBoolE[E ~bool](o any) (E, error) {
	var zero E
	o = reflectx.IndirectToInterface(o, emptyInt64er, emptyFloat64er, emptyValuer)
	switch b := o.(type) {
	case bool:
		return E(b), nil
	case int, int64, int32, int16, int8,
		uint, uint64, uint32, uint16, uint8,
		float64, float32,
		int64er, float64er,
		time.Duration:
		v, err := ToIntE(o)
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return v != 0, nil
	case string:
		v, err := strconv.ParseBool(o.(string))
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return E(v), err
	case driver.Valuer:
		v, err := b.Value()
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return toBoolE[E](v)
	case nil:
		return zero, nil
	default:
		return zero, fmt.Errorf(failedCast, o, o, zero)
	}
}
