package convx

import (
	"database/sql/driver"
	"github.com/go-leo/gox/reflectx"
	"golang.org/x/exp/constraints"
	"reflect"
	"strconv"
	"time"
)

// ToFloat64 casts an interface to a float64 type.
func ToFloat64(o any) float64 {
	v, _ := ToFloat64E(o)
	return v
}

// ToFloat32 casts an interface to a float32 type.
func ToFloat32(o any) float32 {
	v, _ := ToFloat32E(o)
	return v
}

// ToFloat64Slice casts an interface to a float64 slice type.
func ToFloat64Slice(o any) []float64 {
	v, _ := ToFloat64SliceE(o)
	return v
}

// ToFloat32Slice casts an interface to a float32 slice type.
func ToFloat32Slice(o any) []float32 {
	v, _ := ToFloat32SliceE(o)
	return v
}

// ToFloat64E casts an interface to a float64 type.
func ToFloat64E(o any) (float64, error) {
	return ToFloatE[float64](o)
}

// ToFloat32E casts an interface to a float32 type.
func ToFloat32E(o any) (float32, error) {
	return ToFloatE[float32](o)
}

// ToFloat64SliceE casts an interface to a float64 type.
func ToFloat64SliceE(o any) ([]float64, error) {
	return ToFloatSliceE[[]float64](o)
}

// ToFloat32SliceE casts an interface to a float32 type.
func ToFloat32SliceE(o any) ([]float32, error) {
	return ToFloatSliceE[[]float32](o)
}

// ToFloat converts an interface to a floating-point type.
func ToFloat[E constraints.Float](o any) E {
	v, _ := ToFloatE[E](o)
	return v
}

// ToFloatE converts an interface to a floating-point type.
func ToFloatE[E constraints.Float](o any) (E, error) {
	return toFloatE[E](o)
}

// ToFloatSlice converts an interface to a floating-point slice type.
func ToFloatSlice[S ~[]E, E constraints.Float](o any) S {
	v, _ := ToFloatSliceE[S](o)
	return v
}

// ToFloatSliceE converts an interface to a floating-point slice type.
func ToFloatSliceE[S ~[]E, E constraints.Float](o any) (S, error) {
	return toSliceE[S](o, toFloatE[E])
}

func toFloatE[E constraints.Float](o any) (E, error) {
	var zero E
	if o == nil {
		return zero, nil
	}
	// fast path
	switch f := o.(type) {
	case bool:
		if f {
			return 1, nil
		}
		return zero, nil
	case int:
		return E(f), nil
	case int64:
		return E(f), nil
	case int32:
		return E(f), nil
	case int16:
		return E(f), nil
	case int8:
		return E(f), nil
	case uint:
		return E(f), nil
	case uint64:
		return E(f), nil
	case uint32:
		return E(f), nil
	case uint16:
		return E(f), nil
	case uint8:
		return E(f), nil
	case float64:
		return E(f), nil
	case float32:
		return E(f), nil
	case string:
		v, err := strconv.ParseFloat(f, 64)
		if err != nil {
			return failedCastErrValue[E](o, err)
		}
		return E(v), nil
	case time.Duration:
		return E(f), nil
	case time.Weekday:
		return E(f), nil
	case time.Month:
		return E(f), nil
	case int64er:
		v, err := f.Int64()
		if err != nil {
			return failedCastErrValue[E](o, err)
		}
		return E(v), nil
	case float64er:
		v, err := f.Float64()
		if err != nil {
			return failedCastErrValue[E](o, err)
		}
		return E(v), nil
	case driver.Valuer:
		v, err := f.Value()
		if err != nil {
			return failedCastErrValue[E](o, err)
		}
		return toFloatE[E](v)
	default:
		// slow path
		return toFloatValueE[E](o)
	}
}

func toFloatValueE[E constraints.Float](o any) (E, error) {
	var zero E
	v := reflectx.IndirectValue(reflect.ValueOf(o))
	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return zero, nil
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		return E(v.Int()), nil
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		return E(v.Uint()), nil
	case reflect.Float64, reflect.Float32:
		return E(v.Float()), nil
	case reflect.String:
		f, err := strconv.ParseFloat(v.String(), 64)
		if err != nil {
			return failedCastErrValue[E](o, err)
		}
		return E(f), nil
	default:
		return failedCastValue[E](o)
	}
}
