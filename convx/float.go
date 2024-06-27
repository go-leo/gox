package convx

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-leo/gox/operator"
	"golang.org/x/exp/constraints"
	"strconv"
	"time"
)

// ToFloat64 casts an interface to a float64 type.
func ToFloat64(i any) float64 {
	v, _ := ToFloat64E(i)
	return v
}

// ToFloat32 casts an interface to a float32 type.
func ToFloat32(i any) float32 {
	v, _ := ToFloat32E(i)
	return v
}

// ToFloat64E casts an interface to a float64 type.
func ToFloat64E(i any) (float64, error) {
	return ToFloatE[float64](i)
}

// ToFloat32E casts an interface to a float32 type.
func ToFloat32E(i any) (float32, error) {
	return ToFloatE[float32](i)
}

//// ToFloat64Slice casts an interface to a float64 slice type.
//func ToFloat64Slice(i any) []float64 {
//	v, _ := ToFloat64SliceE(i)
//	return v
//}
//
//// ToFloat32Slice casts an interface to a float32 slice type.
//func ToFloat32Slice(i any) []float32 {
//	v, _ := ToFloat32SliceE(i)
//	return v
//}

//// ToFloat64E casts an interface to a float64 type.
//func ToFloat64SliceE(i any) ([]float64, error) {
//	return ToFloatE[float64](i)
//}
//
//// ToFloat32E casts an interface to a float32 type.
//func ToFloat32SliceE(i any) ([]float32, error) {
//	return ToFloatE[float32](i)
//}

// ToFloat converts an interface to a floating-point type.
func ToFloat[N constraints.Float](i any) N {
	v, _ := ToFloatE[N](i)
	return v
}

// ToFloatE converts an interface to a floating-point type.
func ToFloatE[N constraints.Float](i any) (N, error) {
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
	case string:
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return N(v), nil
		}
		return zero, fmt.Errorf("unable to convert %#v of type %T to %T", i, i, zero)
	case bool:
		if s {
			return 1, nil
		}
		return zero, nil
	case json.Number:
		return ToFloatE[N](string(s))
	case time.Weekday:
		return N(s), nil
	case time.Month:
		return N(s), nil
	case sql.NullInt64:
		return N(s.Int64), operator.Ternary(s.Valid, nil, ErrValueIsNULL)
	case sql.NullInt32:
		return N(s.Int32), operator.Ternary(s.Valid, nil, ErrValueIsNULL)
	case sql.NullInt16:
		return N(s.Int16), operator.Ternary(s.Valid, nil, ErrValueIsNULL)
	case sql.NullByte:
		return N(s.Byte), operator.Ternary(s.Valid, nil, ErrValueIsNULL)
	case sql.NullFloat64:
		return N(s.Float64), operator.Ternary(s.Valid, nil, ErrValueIsNULL)
	case sql.NullString:
		if !s.Valid {
			return zero, ErrValueIsNULL
		}
		v, err := strconv.ParseFloat(s.String, 64)
		if err == nil {
			return N(v), nil
		}
		return zero, fmt.Errorf("unable to cast %#v of type %T to %T", i, i, zero)
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
