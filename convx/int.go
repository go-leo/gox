package convx

import (
	"database/sql"
	"fmt"
	"golang.org/x/exp/constraints"
	"reflect"
	"strconv"
	"time"
)

// ToInt converts an interface to an int type.
func ToInt(i any) int {
	v, _ := ToIntE(i)
	return v
}

// ToInt64 converts an interface to an int64 type.
func ToInt64(i any) int64 {
	v, _ := ToInt64E(i)
	return v
}

// ToInt32 converts an interface to an int32 type.
func ToInt32(i any) int32 {
	v, _ := ToInt32E(i)
	return v
}

// ToInt16 converts an interface to an int16 type.
func ToInt16(i any) int16 {
	v, _ := ToInt16E(i)
	return v
}

// ToInt8 converts an interface to an int8 type.
func ToInt8(i any) int8 {
	v, _ := ToInt8E(i)
	return v
}

// ToIntE converts an interface to an int type.
func ToIntE(i any) (int, error) {
	return ToSignedE[int](i)
}

// ToInt64E converts an interface to an int64 type.
func ToInt64E(i any) (int64, error) {
	return ToSignedE[int64](i)
}

// ToInt32E converts an interface to an int32 type.
func ToInt32E(i any) (int32, error) {
	return ToSignedE[int32](i)
}

// ToInt16E converts an interface to an int16 type.
func ToInt16E(i any) (int16, error) {
	return ToSignedE[int16](i)
}

// ToInt8E converts an interface to an int8 type.
func ToInt8E(i any) (int8, error) {
	return ToSignedE[int8](i)
}

// ToIntSlice casts an interface to a []int type.
func ToIntSlice(i interface{}) []int {
	v, _ := ToIntSliceE(i)
	return v
}

// ToInt64Slice casts an interface to a []int64 type.
func ToInt64Slice(i interface{}) []int64 {
	v, _ := ToInt64SliceE(i)
	return v
}

// ToInt32Slice casts an interface to a []int32 type.
func ToInt32Slice(i interface{}) []int32 {
	v, _ := ToInt32SliceE(i)
	return v
}

// ToIntSliceE casts an interface to a []int type.
func ToIntSliceE(i interface{}) ([]int, error) {
	if i == nil {
		return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
	}

	switch v := i.(type) {
	case []int:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToIntE(s.Index(j).Interface())
			if err != nil {
				return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
	}
}

// ToInt64SliceE casts an interface to a []int64 type.
func ToInt64SliceE(i interface{}) ([]int64, error) {
	if i == nil {
		return []int64{}, fmt.Errorf("unable to cast %#v of type %T to []int64", i, i)
	}

	switch v := i.(type) {
	case []int64:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int64, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToInt64E(s.Index(j).Interface())
			if err != nil {
				return []int64{}, fmt.Errorf("unable to cast %#v of type %T to []int64", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int64{}, fmt.Errorf("unable to cast %#v of type %T to []int64", i, i)
	}
}

// ToInt32SliceE casts an interface to a []int32 type.
func ToInt32SliceE(i interface{}) ([]int32, error) {
	if i == nil {
		return []int32{}, fmt.Errorf("unable to cast %#v of type %T to []int32", i, i)
	}

	switch v := i.(type) {
	case []int32:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int32, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToInt32E(s.Index(j).Interface())
			if err != nil {
				return []int32{}, fmt.Errorf("unable to cast %#v of type %T to []int32", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int32{}, fmt.Errorf("unable to cast %#v of type %T to []int32", i, i)
	}
}

// ToSigned converts an interface to a signed integer type.
func ToSigned[N constraints.Signed](i any) N {
	v, _ := ToSignedE[N](i)
	return v
}

// ToSignedE converts an interface to a signed integer type.
func ToSignedE[N constraints.Signed](i any) (N, error) {
	var zero N
	i = indirect(i)
	switch s := i.(type) {
	case int:
		return N(s), nil
	case int64:
		return N(s), nil
	case int32:
		return N(s), nil
	case int16:
		return N(s), nil
	case int8:
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
		return N(s), nil
	case float32:
		return N(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return zero, nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(s), 0, 0)
		if err == nil {
			return N(v), nil
		}
		return zero, fmt.Errorf("unable to convert %#v of type %T to %T", i, i, zero)
	case time.Weekday:
		return N(s), nil
	case time.Month:
		return N(s), nil
	case sql.NullInt64:
		if !s.Valid {
			return zero, ErrValueIsNULL
		}
		return N(s.Int64), nil
	case sql.NullInt32:
		if !s.Valid {
			return zero, ErrValueIsNULL
		}
		return N(s.Int32), nil
	case sql.NullInt16:
		if !s.Valid {
			return zero, ErrValueIsNULL
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
		return N(s.Float64), nil
	case sql.NullString:
		if !s.Valid {
			return zero, ErrValueIsNULL
		}
		v, err := strconv.ParseInt(trimZeroDecimal(s.String), 0, 0)
		if err == nil {
			return N(v), nil
		}
		return zero, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
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

// ToSignedSlice converts an interface to a signed integer slice type.
func ToSignedSlice[S []N, N constraints.Signed](i any) S {
	v, _ := ToSignedSliceE[S](i)
	return v
}

// ToSignedSliceE converts an interface to a signed integer slice type.
func ToSignedSliceE[S []N, N constraints.Signed](i any) (S, error) {
	var zero S
	if i == nil {
		return zero, fmt.Errorf("unable to cast %#v of type %T to %T", i, i, zero)
	}

	if v, ok := i.(S); ok {
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make(S, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToSignedE[N](s.Index(j).Interface())
			if err != nil {
				return zero, fmt.Errorf("unable to cast %#v of type %T to %T, %w", i, i, zero, err)
			}
			a[j] = val
		}
		return a, nil
	default:
		return zero, fmt.Errorf("unable to cast %#v of type %T to %T", i, i, zero)
	}
}
