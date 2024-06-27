package convx

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/exp/constraints"
	"reflect"
	"strconv"
	"time"
)

// ToUint converts an interface to a uint type.
func ToUint(i any) uint {
	v, _ := ToUintE(i)
	return v
}

// ToUint64 converts an interface to a uint64 type.
func ToUint64(i any) uint64 {
	v, _ := ToUint64E(i)
	return v
}

// ToUint32 converts an interface to a uint32 type.
func ToUint32(i any) uint32 {
	v, _ := ToUint32E(i)
	return v
}

// ToUint16 converts an interface to a uint16 type.
func ToUint16(i any) uint16 {
	v, _ := ToUint16E(i)
	return v
}

// ToUint8 converts an interface to a uint8 type.
func ToUint8(i any) uint8 {
	v, _ := ToUint8E(i)
	return v
}

// ToUintE converts an interface to a uint type.
func ToUintE(i any) (uint, error) {
	return ToUnsignedE[uint](i)
}

// ToUint64E converts an interface to a uint64 type.
func ToUint64E(i any) (uint64, error) {
	return ToUnsignedE[uint64](i)
}

// ToUint32E converts an interface to a uint32 type.
func ToUint32E(i any) (uint32, error) {
	return ToUnsignedE[uint32](i)
}

// ToUint16E converts an interface to a uint16 type.
func ToUint16E(i any) (uint16, error) {
	return ToUnsignedE[uint16](i)
}

// ToUint8E converts an interface to a uint type.
func ToUint8E(i any) (uint8, error) {
	return ToUnsignedE[uint8](i)
}

// ToUintSlice casts an interface to a []uint type.
func ToUintSlice(i interface{}) []uint {
	v, _ := ToUintSliceE(i)
	return v
}

// ToUint64Slice casts an interface to a []uint64 type.
func ToUint64Slice(i interface{}) []uint64 {
	v, _ := ToUint64SliceE(i)
	return v
}

// ToUint32Slice casts an interface to a []uint32 type.
func ToUint32Slice(i interface{}) []uint32 {
	v, _ := ToUint32SliceE(i)
	return v
}

// ToUintSliceE casts an interface to a []uint type.
func ToUintSliceE(i interface{}) ([]uint, error) {
	if i == nil {
		return []uint{}, fmt.Errorf("unable to cast %#v of type %T to []uint", i, i)
	}

	switch v := i.(type) {
	case []uint:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]uint, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToUintE(s.Index(j).Interface())
			if err != nil {
				return []uint{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []uint{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
	}
}

// ToUint64SliceE casts an interface to a []uint64 type.
func ToUint64SliceE(i interface{}) ([]uint64, error) {
	if i == nil {
		return []uint64{}, fmt.Errorf("unable to cast %#v of type %T to []uint64", i, i)
	}

	switch v := i.(type) {
	case []uint64:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]uint64, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToUint64E(s.Index(j).Interface())
			if err != nil {
				return []uint64{}, fmt.Errorf("unable to cast %#v of type %T to []int64", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []uint64{}, fmt.Errorf("unable to cast %#v of type %T to []int64", i, i)
	}
}

// ToUint32SliceE casts an interface to a []int32 type.
func ToUint32SliceE(i interface{}) ([]uint32, error) {
	if i == nil {
		return []uint32{}, fmt.Errorf("unable to cast %#v of type %T to []int32", i, i)
	}

	switch v := i.(type) {
	case []uint32:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]uint32, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToUint32E(s.Index(j).Interface())
			if err != nil {
				return []uint32{}, fmt.Errorf("unable to cast %#v of type %T to []int32", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []uint32{}, fmt.Errorf("unable to cast %#v of type %T to []int32", i, i)
	}
}

// ToUnsigned converts an interface to a unsigned integer type.
func ToUnsigned[N constraints.Unsigned](i any) N {
	v, _ := ToUnsignedE[N](i)
	return v
}

// ToUnsignedE converts an interface to a unsigned integer type.
func ToUnsignedE[N constraints.Unsigned](i any) (N, error) {
	var zero N
	i = indirect(i)
	switch s := i.(type) {
	case int:
		if s < 0 {
			return zero, ErrNegativeNotAllowed
		}
		return N(s), nil
	case int64:
		if s < 0 {
			return zero, ErrNegativeNotAllowed
		}
		return N(s), nil
	case int32:
		if s < 0 {
			return zero, ErrNegativeNotAllowed
		}
		return N(s), nil
	case int16:
		if s < 0 {
			return zero, ErrNegativeNotAllowed
		}
		return N(s), nil
	case int8:
		if s < 0 {
			return zero, ErrNegativeNotAllowed
		}
		return N(s), nil
	case uint:
		return N(s), nil
	case uint64:
		return N(s), nil
	case uint32:
		return N(s), nil
	case uint16:
		return N(s), nil
	case uint8:
		return N(s), nil
	case float64:
		if s < 0 {
			return zero, ErrNegativeNotAllowed
		}
		return N(s), nil
	case float32:
		if s < 0 {
			return zero, ErrNegativeNotAllowed
		}
		return N(s), nil
	case string:
		v, err := strconv.ParseUint(trimZeroDecimal(s), 0, 0)
		if err == nil {
			if v < 0 {
				return zero, ErrNegativeNotAllowed
			}
			return N(v), nil
		}
		return zero, fmt.Errorf("unable to convert %#v of type %T to int", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return zero, nil
	case json.Number:
		return ToUnsignedE[N](string(s))
	case time.Weekday:
		if s < 0 {
			return zero, ErrNegativeNotAllowed
		}
		return N(s), nil
	case time.Month:
		if s < 0 {
			return zero, ErrNegativeNotAllowed
		}
		return N(s), nil
	case sql.NullInt64:
		if !s.Valid {
			return zero, ErrValueIsNULL
		}
		if s.Int64 < 0 {
			return zero, ErrNegativeNotAllowed
		}
		return N(s.Int64), nil
	case sql.NullInt32:
		if !s.Valid {
			return zero, ErrValueIsNULL
		}
		if s.Int32 < 0 {
			return zero, ErrNegativeNotAllowed
		}
		return N(s.Int32), nil
	case sql.NullInt16:
		if !s.Valid {
			return zero, ErrValueIsNULL
		}
		if s.Int16 < 0 {
			return zero, ErrNegativeNotAllowed
		}
		return N(s.Int16), nil
	case sql.NullByte:
		if !s.Valid {
			return zero, ErrValueIsNULL
		}
		return N(s.Byte), nil
	case sql.NullFloat64:
		if !s.Valid {
			return zero, ErrValueIsNULL
		}
		if s.Float64 < 0 {
			return zero, ErrNegativeNotAllowed
		}
		return N(s.Float64), nil
	case sql.NullString:
		if !s.Valid {
			return zero, ErrValueIsNULL
		}
		v, err := strconv.ParseInt(trimZeroDecimal(s.String), 0, 0)
		if err == nil {
			if v < 0 {
				return zero, ErrNegativeNotAllowed
			}
			return N(v), nil
		}
		return zero, fmt.Errorf("unable to convert %#v of type %T to %T", i, i, zero)
	case interface{ Int64() (int64, error) }:
		v, err := s.Int64()
		return N(v), err
	case interface{ Float64() (float64, error) }:
		v, err := s.Float64()
		return N(v), err
	case nil:
		return zero, nil
	default:
		return zero, fmt.Errorf("unable to convert %#v of type %T to %T", i, i, zero)
	}
}
