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

// ToUint converts an interface to a uint type.
func ToUint(o any) uint {
	v, _ := ToUintE(o)
	return v
}

// ToUint64 converts an interface to a uint64 type.
func ToUint64(o any) uint64 {
	v, _ := ToUint64E(o)
	return v
}

// ToUint32 converts an interface to a uint32 type.
func ToUint32(o any) uint32 {
	v, _ := ToUint32E(o)
	return v
}

// ToUint16 converts an interface to a uint16 type.
func ToUint16(o any) uint16 {
	v, _ := ToUint16E(o)
	return v
}

// ToUint8 converts an interface to a uint8 type.
func ToUint8(o any) uint8 {
	v, _ := ToUint8E(o)
	return v
}

// ToUintSlice casts an interface to a []uint type.
func ToUintSlice(o any) []uint {
	v, _ := ToUintSliceE(o)
	return v
}

// ToUint64Slice casts an interface to a []uint64 type.
func ToUint64Slice(o any) []uint64 {
	v, _ := ToUint64SliceE(o)
	return v
}

// ToUint32Slice casts an interface to a []uint32 type.
func ToUint32Slice(o any) []uint32 {
	v, _ := ToUint32SliceE(o)
	return v
}

// ToUint16Slice converts an interface to a []uint16 type.
func ToUint16Slice(o any) []uint16 {
	v, _ := ToUint16SliceE(o)
	return v
}

// ToUint8Slice converts an interface to a []uint8 type.
func ToUint8Slice(o any) []uint8 {
	v, _ := ToUint8SliceE(o)
	return v
}

// ToUintE converts an interface to a uint type.
func ToUintE(o any) (uint, error) {
	return ToUnsignedE[uint](o)
}

// ToUint64E converts an interface to a uint64 type.
func ToUint64E(o any) (uint64, error) {
	return ToUnsignedE[uint64](o)
}

// ToUint32E converts an interface to a uint32 type.
func ToUint32E(o any) (uint32, error) {
	return ToUnsignedE[uint32](o)
}

// ToUint16E converts an interface to a uint16 type.
func ToUint16E(o any) (uint16, error) {
	return ToUnsignedE[uint16](o)
}

// ToUint8E converts an interface to a uint type.
func ToUint8E(o any) (uint8, error) {
	return ToUnsignedE[uint8](o)
}

// ToUintSliceE casts an interface to a []uint type.
func ToUintSliceE(a any) ([]uint, error) {
	return ToUnsignedSliceE[[]uint](a)
}

// ToUint64SliceE casts an interface to a []uint64 type.
func ToUint64SliceE(o any) ([]uint64, error) {
	return ToUnsignedSliceE[[]uint64](o)
}

// ToUint32SliceE casts an interface to a []int32 type.
func ToUint32SliceE(o any) ([]uint32, error) {
	return ToUnsignedSliceE[[]uint32](o)
}

// ToUint16SliceE converts an interface to a uint16 type.
func ToUint16SliceE(o any) ([]uint16, error) {
	return ToUnsignedSliceE[[]uint16](o)
}

// ToUint8SliceE converts an interface to a uint type.
func ToUint8SliceE(o any) ([]uint8, error) {
	return ToUnsignedSliceE[[]uint8](o)
}

// ToUnsigned converts an interface to a unsigned integer type.
func ToUnsigned[N constraints.Unsigned](o any) N {
	v, _ := ToUnsignedE[N](o)
	return v
}

// ToUnsignedE converts an interface to a unsigned integer type.
func ToUnsignedE[E constraints.Unsigned](o any) (E, error) {
	return toUnsignedE[E](o)
}

// ToUnsignedSlice converts an interface to an unsigned integer slice type.
func ToUnsignedSlice[S ~[]E, E constraints.Unsigned](o any) S {
	v, _ := ToUnsignedSliceE[S](o)
	return v
}

// ToUnsignedSliceE converts an interface to an unsigned integer slice type.
func ToUnsignedSliceE[S ~[]E, E constraints.Unsigned](o any) (S, error) {
	return toSliceE[S](o, toUnsignedE[E])
}

func toUnsignedE[E constraints.Unsigned](o any) (E, error) {
	var zero E
	o = reflectx.IndirectToInterface(o,
		reflect.TypeOf((*interface{ Int64() (int64, error) })(nil)).Elem(),
		reflect.TypeOf((*interface{ Float64() (float64, error) })(nil)).Elem(),
		reflect.TypeOf((*driver.Valuer)(nil)).Elem(),
	)
	switch u := o.(type) {
	case int:
		if u < 0 {
			return zero, fmt.Errorf(failedCast, o, o, zero)
		}
		return E(u), nil
	case int64:
		if u < 0 {
			return zero, fmt.Errorf(failedCast, o, o, zero)
		}
		return E(u), nil
	case int32:
		if u < 0 {
			return zero, fmt.Errorf(failedCast, o, o, zero)
		}
		return E(u), nil
	case int16:
		if u < 0 {
			return zero, fmt.Errorf(failedCast, o, o, zero)
		}
		return E(u), nil
	case int8:
		if u < 0 {
			return zero, fmt.Errorf(failedCast, o, o, zero)
		}
		return E(u), nil
	case uint:
		return E(u), nil
	case uint64:
		return E(u), nil
	case uint32:
		return E(u), nil
	case uint16:
		return E(u), nil
	case uint8:
		return E(u), nil
	case float64:
		if u < 0 {
			return zero, fmt.Errorf(failedCast, o, o, zero)
		}
		return E(u), nil
	case float32:
		if u < 0 {
			return zero, fmt.Errorf(failedCast, o, o, zero)
		}
		return E(u), nil
	case interface{ Int64() (int64, error) }: // json.Number
		v, err := u.Int64()
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		if v < 0 {
			return zero, fmt.Errorf(failedCast, o, o, zero)
		}
		return E(v), err
	case interface{ Float64() (float64, error) }: // json.Number
		v, err := u.Float64()
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		if v < 0 {
			return zero, fmt.Errorf(failedCast, o, o, zero)
		}
		return E(v), err
	case string:
		v, err := strconv.ParseUint(trimZeroDecimal(u), 0, 0)
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		if v < 0 {
			return zero, fmt.Errorf(failedCast, o, o, zero)
		}
		return E(v), nil
	case bool:
		if u {
			return 1, nil
		}
		return zero, nil
	case time.Duration:
		if u < 0 {
			return zero, fmt.Errorf(failedCast, o, o, zero)
		}
		return E(u), nil
	case time.Weekday:
		if u < 0 {
			return zero, fmt.Errorf(failedCast, o, o, zero)
		}
		return E(u), nil
	case time.Month:
		if u < 0 {
			return zero, fmt.Errorf(failedCast, o, o, zero)
		}
		return E(u), nil
	case driver.Valuer:
		v, err := u.Value()
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return toUnsignedE[E](v)
	case nil:
		return zero, nil
	default:
		return zero, fmt.Errorf(failedCast, o, o, zero)
	}
}
