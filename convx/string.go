package convx

import (
	"fmt"
	"github.com/go-leo/gox/reflectx"
	"html/template"
	"strconv"
)

// ToString casts an interface to a string type.
func ToString[E ~string](o any) E {
	v, _ := ToStringE[E](o)
	return v
}

// ToStringE casts an interface to a string type.
func ToStringE[E ~string](o any) (E, error) {
	return toStringE[E](o)
}

// ToStringSlice casts an interface to a []string type.
func ToStringSlice[S ~[]E, E ~string](o any) S {
	v, _ := ToStringSliceE[S](o)
	return v
}

// ToStringSliceE casts an interface to a []string type.
func ToStringSliceE[S ~[]E, E ~string](o any) (S, error) {
	return toSliceE[S](o, toStringE[E])
}

func toStringE[E ~string](o any) (E, error) {
	var zero E
	o = reflectx.Indirect(o)
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
	case fmt.Stringer: // json.Number
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
	case nil:
		return "", nil
	default:
		return zero, fmt.Errorf(failedCast, o, o, zero)
	}
}
