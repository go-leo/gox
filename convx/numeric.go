package convx

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-leo/gox/operator"
	"golang.org/x/exp/constraints"
	"strconv"
	"time"
)

var (
	ErrValueIsNULL        = errors.New("unable to convert NULL value")
	ErrNegativeNotAllowed = errors.New("unable to convert negative value")
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

// ToIntE converts an interface to an int type.
func ToIntE(i any) (int, error) {
	return toSignedE[int](i)
}

// ToInt64E converts an interface to an int64 type.
func ToInt64E(i any) (int64, error) {
	return toSignedE[int64](i)
}

// ToInt32E converts an interface to an int32 type.
func ToInt32E(i any) (int32, error) {
	return toSignedE[int32](i)
}

// ToInt16E converts an interface to an int16 type.
func ToInt16E(i any) (int16, error) {
	return toSignedE[int16](i)
}

// ToInt8E converts an interface to an int8 type.
func ToInt8E(i any) (int8, error) {
	return toSignedE[int8](i)
}

// ToUintE converts an interface to a uint type.
func ToUintE(i any) (uint, error) {
	return toUnsignedE[uint](i)
}

// ToUint64E converts an interface to a uint64 type.
func ToUint64E(i any) (uint64, error) {
	return toUnsignedE[uint64](i)
}

// ToUint32E converts an interface to a uint32 type.
func ToUint32E(i any) (uint32, error) {
	return toUnsignedE[uint32](i)
}

// ToUint16E converts an interface to a uint16 type.
func ToUint16E(i any) (uint16, error) {
	return toUnsignedE[uint16](i)
}

// ToUint8E converts an interface to a uint type.
func ToUint8E(i any) (uint8, error) {
	return toUnsignedE[uint8](i)
}

// ToFloat64E casts an interface to a float64 type.
func ToFloat64E(i any) (float64, error) {
	return toFloatE[float64](i)
}

// ToFloat32E casts an interface to a float32 type.
func ToFloat32E(i any) (float32, error) {
	return toFloatE[float32](i)
}

// toSignedE converts an interface to a signed integer type.
func toSignedE[N constraints.Signed](i any) (N, error) {
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

// toUnsignedE converts an interface to a unsigned integer type.
func toUnsignedE[N constraints.Unsigned](i any) (N, error) {
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
		return toUnsignedE[N](string(s))
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

// toFloatE converts an interface to a floating-point type.
func toFloatE[N constraints.Float](i any) (N, error) {
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
		return toFloatE[N](string(s))
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

func trimZeroDecimal(s string) string {
	var foundZero bool
	for i := len(s); i > 0; i-- {
		switch s[i-1] {
		case '.':
			if foundZero {
				return s[:i-1]
			}
		case '0':
			foundZero = true
		default:
			return s
		}
	}
	return s
}
