package convx

import (
	"golang.org/x/exp/slices"
	"reflect"
)

// ToSlice casts an interface to a []any type.
func ToSlice(o any) []any {
	v, _ := ToSliceE(o)
	return v
}

// ToSliceE casts an interface to a []any type.
func ToSliceE(o any) ([]any, error) {
	return toSliceE[[]any](o, func(o any) (any, error) { return o, nil })
}

func toSliceE[S ~[]E, E any](o any, to func(o any) (E, error)) (S, error) {
	var zero S
	if o == nil {
		return zero, nil
	}
	if v, ok := o.(S); ok {
		return slices.Clone(v), nil
	}
	kind := reflect.TypeOf(o).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		value := reflect.ValueOf(o)
		res := make(S, value.Len())
		for i := 0; i < value.Len(); i++ {
			val, err := to(value.Index(i).Interface())
			if err != nil {
				return zero, err
			}
			res[i] = val
		}
		return res, nil
	default:
		return failedCastValue[S](o)
	}
}
