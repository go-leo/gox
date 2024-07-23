package convx

import (
	"database/sql/driver"
	"fmt"
	"github.com/go-leo/gox/reflectx"
	"golang.org/x/exp/constraints"
	"strconv"
	"time"
)

// ToInt converts an interface to an int type.
func ToInt(o any) int {
	v, _ := ToIntE(o)
	return v
}

// ToInt64 converts an interface to an int64 type.
func ToInt64(o any) int64 {
	v, _ := ToInt64E(o)
	return v
}

// ToInt32 converts an interface to an int32 type.
func ToInt32(o any) int32 {
	v, _ := ToInt32E(o)
	return v
}

// ToInt16 converts an interface to an int16 type.
func ToInt16(o any) int16 {
	v, _ := ToInt16E(o)
	return v
}

// ToInt8 converts an interface to an int8 type.
func ToInt8(o any) int8 {
	v, _ := ToInt8E(o)
	return v
}

// ToIntSlice casts an interface to a []int type.
func ToIntSlice(o any) []int {
	v, _ := ToIntSliceE(o)
	return v
}

// ToInt64Slice casts an interface to a []int64 type.
func ToInt64Slice(o any) []int64 {
	v, _ := ToInt64SliceE(o)
	return v
}

// ToInt32Slice casts an interface to a []int32 type.
func ToInt32Slice(o any) []int32 {
	v, _ := ToInt32SliceE(o)
	return v
}

// ToInt16Slice converts an interface to an int16 type.
func ToInt16Slice(o any) []int16 {
	v, _ := ToInt16SliceE(o)
	return v
}

// ToInt8Slice converts an interface to an int8 type.
func ToInt8Slice(o any) []int8 {
	v, _ := ToInt8SliceE(o)
	return v
}

// ToIntE converts an interface to an int type.
func ToIntE(o any) (int, error) {
	return ToSignedE[int](o)
}

// ToInt64E converts an interface to an int64 type.
func ToInt64E(o any) (int64, error) {
	return ToSignedE[int64](o)
}

// ToInt32E converts an interface to an int32 type.
func ToInt32E(o any) (int32, error) {
	return ToSignedE[int32](o)
}

// ToInt16E converts an interface to an int16 type.
func ToInt16E(o any) (int16, error) {
	return ToSignedE[int16](o)
}

// ToInt8E converts an interface to an int8 type.
func ToInt8E(o any) (int8, error) {
	return ToSignedE[int8](o)
}

// ToIntSliceE casts an interface to a []int type.
func ToIntSliceE(o any) ([]int, error) {
	return ToSignedSliceE[[]int](o)
}

// ToInt64SliceE casts an interface to a []int64 type.
func ToInt64SliceE(o any) ([]int64, error) {
	return ToSignedSliceE[[]int64](o)
}

// ToInt32SliceE casts an interface to a []int32 type.
func ToInt32SliceE(o any) ([]int32, error) {
	return ToSignedSliceE[[]int32](o)
}

// ToInt16SliceE converts an interface to an []int16 type.
func ToInt16SliceE(o any) ([]int16, error) {
	return ToSignedSliceE[[]int16](o)
}

// ToInt8SliceE converts an interface to an []int8 type.
func ToInt8SliceE(o any) ([]int8, error) {
	return ToSignedSliceE[[]int8](o)
}

// ToSigned converts an interface to a signed integer type.
func ToSigned[E constraints.Signed](o any) E {
	v, _ := ToSignedE[E](o)
	return v
}

// ToSignedE converts an interface to a signed integer type.
func ToSignedE[E constraints.Signed](o any) (E, error) {
	return toSignedE[E](o)
}

// ToSignedSlice converts an interface to a signed integer slice type.
func ToSignedSlice[S ~[]E, E constraints.Signed](o any) S {
	v, _ := ToSignedSliceE[S](o)
	return v
}

// ToSignedSliceE converts an interface to a signed integer slice type.
func ToSignedSliceE[S ~[]E, E constraints.Signed](o any) (S, error) {
	return toSliceE[S](o, ToSignedE[E])
}

func toSignedE[E constraints.Signed](o any) (E, error) {
	var zero E
	o = reflectx.IndirectToInterface(o, emptyInt64er, emptyFloat64er, emptyValuer)
	switch s := o.(type) {
	case int:
		return E(s), nil
	case int64:
		return E(s), nil
	case int32:
		return E(s), nil
	case int16:
		return E(s), nil
	case int8:
		return E(s), nil
	case uint:
		return E(s), nil
	case uint64:
		return E(s), nil
	case uint32:
		return E(s), nil
	case uint16:
		return E(s), nil
	case uint8:
		return E(s), nil
	case float64:
		return E(s), nil
	case float32:
		return E(s), nil
	case int64er:
		v, err := s.Int64()
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return E(v), nil
	case float64er:
		v, err := s.Float64()
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return E(v), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(s), 0, 0)
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return E(v), nil
	case bool:
		if s {
			return 1, nil
		}
		return zero, nil
	case time.Duration:
		return E(s), nil
	case time.Weekday:
		return E(s), nil
	case time.Month:
		return E(s), nil
	case driver.Valuer:
		v, err := s.Value()
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return toSignedE[E](v)
	case nil:
		return zero, nil
	default:
		return zero, fmt.Errorf(failedCast, o, o, zero)
	}
}
