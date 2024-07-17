package convx

import (
	"database/sql/driver"
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
	o = reflectx.IndirectToInterface(o,
		reflect.TypeOf((*fmt.Stringer)(nil)).Elem(),
		reflect.TypeOf((*error)(nil)).Elem(),
		reflect.TypeOf((*driver.Valuer)(nil)).Elem(),
	)
	switch s := o.(type) {
	case string:
		return E(s), nil
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
	case []byte:
		return E(string(s)), nil
	case fmt.Stringer:
		return E(s.String()), nil
	case error:
		return E(s.Error()), nil
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
	case driver.Valuer:
		v, err := s.Value()
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return toTextE[E](v)
	case nil:
		return "", nil
	default:
		return zero, fmt.Errorf(failedCast, o, o, zero)
	}
}
