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
	ErrValueIsNULL        = errors.New("convx: unable to convert NULL value")
	ErrNegativeNotAllowed = errors.New("convx: unable to convert negative value")
)

// ToSigned converts an interface to a signed integer type.
func ToSigned[N constraints.Signed](i any) N {
	v, _ := ToSignedE[N](i)
	return v
}

// ToUnsigned converts an interface to a unsigned integer type.
func ToUnsigned[N constraints.Unsigned](i any) N {
	v, _ := ToUnsignedE[N](i)
	return v
}

// ToFloat converts an interface to a floating-point type.
func ToFloat[N constraints.Float](i any) N {
	v, _ := ToFloatE[N](i)
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
