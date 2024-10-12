package convx

import (
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"fmt"
	"github.com/go-leo/gox/reflectx"
	"html/template"
	"reflect"
	"strconv"
)

// ToString casts an interface to a string type.
func ToString(o any) string {
	return ToText[string](o)
}

// ToStringE casts an interface to a string type.
func ToStringE(o any) (string, error) {
	return ToTextE[string](o)
}

// ToStringSlice casts an interface to a []string type.
func ToStringSlice(o any) []string {
	return ToTextSlice[[]string](o)
}

// ToStringSliceE casts an interface to a []string type.
func ToStringSliceE(o any) ([]string, error) {
	return ToTextSliceE[[]string](o)
}

// ToText casts an interface to a string type.
func ToText[E ~string](o any) E {
	v, _ := ToTextE[E](o)
	return v
}

// ToTextE casts an interface to a string type.
func ToTextE[E ~string](o any) (E, error) {
	return toTextE[E](o)
}

// ToTextSlice casts an interface to a []string type.
func ToTextSlice[S ~[]E, E ~string](o any) S {
	v, _ := ToTextSliceE[S](o)
	return v
}

// ToTextSliceE casts an interface to a []string type.
func ToTextSliceE[S ~[]E, E ~string](o any) (S, error) {
	return toSliceE[S](o, toTextE[E])
}

func toTextE[E ~string](o any) (E, error) {
	var zero E
	if o == nil {
		return zero, nil
	}
	// fast path
	switch s := o.(type) {
	case bool:
		return E(strconv.FormatBool(s)), nil
	case float64:
		return E(strconv.FormatFloat(s, 'f', -1, 64)), nil
	case float32:
		return E(strconv.FormatFloat(float64(s), 'f', -1, 32)), nil
	case int:
		return E(strconv.Itoa(s)), nil
	case int64:
		return E(strconv.FormatInt(s, 10)), nil
	case int32:
		return E(strconv.FormatInt(int64(s), 10)), nil
	case int16:
		return E(strconv.FormatInt(int64(s), 10)), nil
	case int8:
		return E(strconv.FormatInt(int64(s), 10)), nil
	case uint:
		return E(strconv.FormatUint(uint64(s), 10)), nil
	case uint64:
		return E(strconv.FormatUint(s, 10)), nil
	case uint32:
		return E(strconv.FormatUint(uint64(s), 10)), nil
	case uint16:
		return E(strconv.FormatUint(uint64(s), 10)), nil
	case uint8:
		return E(strconv.FormatUint(uint64(s), 10)), nil
	case string:
		return E(s), nil
	case []byte:
		return E(string(s)), nil
	case template.HTML:
		return E(string(s)), nil
	case template.URL:
		return E(string(s)), nil
	case template.JS:
		return E(string(s)), nil
	case template.CSS:
		return E(string(s)), nil
	case template.HTMLAttr:
		return E(string(s)), nil
	case fmt.Stringer:
		return E(s.String()), nil
	case error:
		return E(s.Error()), nil
	case encoding.TextMarshaler:
		v, err := s.MarshalText()
		if err != nil {
			return failedCastErrValue[E](o, err)
		}
		return E(string(v)), nil
	case json.Marshaler:
		v, err := s.MarshalJSON()
		if err != nil {
			return failedCastErrValue[E](o, err)
		}
		return E(string(v)), nil
	case driver.Valuer:
		v, err := s.Value()
		if err != nil {
			return failedCastErrValue[E](o, err)
		}
		return toTextE[E](v)
	default:
		// slow path
		return toTextValueE[E](o)
	}
}

func toTextValueE[E ~string](o any) (E, error) {
	v := reflectx.IndirectValue(reflect.ValueOf(o))
	switch v.Kind() {
	case reflect.Bool:
		return E(strconv.FormatBool(v.Bool())), nil
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		return E(strconv.FormatInt(v.Int(), 10)), nil
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		return E(strconv.FormatUint(v.Uint(), 10)), nil
	case reflect.Float64, reflect.Float32:
		return E(strconv.FormatFloat(v.Float(), 'f', -1, 64)), nil
	case reflect.String:
		return E(v.String()), nil
	case reflect.Slice:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			return E(string(v.Bytes())), nil
		}
		return failedCastValue[E](o)
	default:
		return failedCastValue[E](o)
	}
}
