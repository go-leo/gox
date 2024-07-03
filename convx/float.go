package convx

import (
	"database/sql/driver"
	"fmt"
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
	o = reflectx.IndirectToInterface(o,
		reflect.TypeOf((*interface{ Int64() (int64, error) })(nil)).Elem(),
		reflect.TypeOf((*interface{ Float64() (float64, error) })(nil)).Elem(),
		reflect.TypeOf((*driver.Valuer)(nil)).Elem(),
	)
	switch f := o.(type) {
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
	case interface{ Int64() (int64, error) }: // json.Number
		v, err := f.Int64()
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return E(v), nil
	case interface{ Float64() (float64, error) }: // json.Number
		v, err := f.Float64()
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return E(v), nil
	case string:
		v, err := strconv.ParseFloat(f, 64)
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return E(v), nil
	case bool:
		if f {
			return 1, nil
		}
		return zero, nil
	case time.Duration:
		return E(f), nil
	case time.Weekday:
		return E(f), nil
	case time.Month:
		return E(f), nil
	case driver.Valuer:
		v, err := f.Value()
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return toFloatE[E](v)
	case nil:
		return zero, nil
	default:
		return zero, fmt.Errorf(failedCast, o, o, zero)
	}
}
