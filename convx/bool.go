package convx

import (
	"database/sql/driver"
	"github.com/go-leo/gox/reflectx"
	"reflect"
	"strconv"
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
	if o == nil {
		return zero, nil
	}
	// fast path
	switch b := o.(type) {
	case bool:
		return E(b), nil
	case int, int64, int32, int16, int8,
		uint, uint64, uint32, uint16, uint8,
		float64, float32,
		int64er, float64er:
		n, err := ToFloat64E(o)
		if err != nil {
			return failedCastErrValue[E](o, err)
		}
		return n != 0, nil
	case string:
		v, err := strconv.ParseBool(o.(string))
		if err != nil {
			return failedCastErrValue[E](o, err)
		}
		return E(v), err
	case driver.Valuer:
		v, err := b.Value()
		if err != nil {
			return failedCastErrValue[E](o, err)
		}
		return toBoolE[E](v)
	default:
		// slow path
		return toBoolValueE[E](o)
	}
}

func toBoolValueE[E ~bool](o any) (E, error) {
	v := reflectx.IndirectValue(reflect.ValueOf(o))
	switch v.Kind() {
	case reflect.Bool:
		return E(v.Bool()), nil
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		return v.Int() != 0, nil
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		return v.Uint() != 0, nil
	case reflect.Float64, reflect.Float32:
		return v.Float() != 0, nil
	case reflect.String:
		b, err := strconv.ParseBool(v.String())
		if err != nil {
			return failedCastErrValue[E](o, err)
		}
		return E(b), err
	default:
		return failedCastValue[E](o)
	}
}
