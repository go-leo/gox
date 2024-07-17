package convx

import (
	"fmt"
	"github.com/go-leo/gox/encodingx/jsonx"
	"reflect"
)

// ToStringMap casts an interface to a map[string]any type.
func ToStringMap(o any) map[string]any {
	v, _ := ToStringMapE(o)
	return v
}

// ToStringMapString casts an interface to a map[string]string type.
func ToStringMapString(o any) map[string]string {
	v, _ := ToStringMapStringE(o)
	return v
}

// ToStringMapBool casts an interface to a map[string]bool type.
func ToStringMapBool(o any) map[string]bool {
	v, _ := ToStringMapBoolE(o)
	return v
}

// ToStringMapInt casts an interface to a map[string]int type.
func ToStringMapInt(o any) map[string]int {
	v, _ := ToStringMapIntE(o)
	return v
}

// ToStringMapInt64 casts an interface to a map[string]int64 type.
func ToStringMapInt64(o any) map[string]int64 {
	v, _ := ToStringMapInt64E(o)
	return v
}

// ToStringMapStringSlice casts an interface to a map[string][]string type.
func ToStringMapStringSlice(o any) map[string][]string {
	v, _ := ToStringMapStringSliceE(o)
	return v
}

// ToStringMapE casts an interface to a map[string]any type.
func ToStringMapE(o any) (map[string]any, error) {
	return toMapE[map[string]any](o, ToTextE[string], func(o any) (any, error) { return o, nil })
}

// ToStringMapStringE casts an interface to a map[string]string type.
func ToStringMapStringE(o any) (map[string]string, error) {
	return toMapE[map[string]string](o, ToTextE[string], ToTextE[string])
}

// ToStringMapBoolE casts an interface to a map[string]bool type.
func ToStringMapBoolE(o any) (map[string]bool, error) {
	return toMapE[map[string]bool](o, ToTextE[string], ToBoolE[bool])
}

// ToStringMapIntE casts an interface to a map[string]int{} type.
func ToStringMapIntE(o any) (map[string]int, error) {
	return toMapE[map[string]int](o, ToTextE[string], ToSignedE[int])
}

// ToStringMapInt64E casts an interface to a map[string]int64{} type.
func ToStringMapInt64E(o any) (map[string]int64, error) {
	return toMapE[map[string]int64](o, ToTextE[string], ToSignedE[int64])
}

// ToStringMapStringSliceE casts an interface to a map[string][]string type.
func ToStringMapStringSliceE(o any) (map[string][]string, error) {
	return toMapE[map[string][]string](o, ToTextE[string], ToTextSliceE[[]string])
}

func toMapE[M ~map[K]V, K comparable, V any](o any, key func(o any) (K, error), val func(o any) (V, error)) (M, error) {
	var zero M
	if o == nil {
		return zero, nil
	}
	if s, ok := o.(string); ok {
		res := make(M)
		err := jsonx.Unmarshal([]byte(s), &res)
		if err != nil {
			return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
		}
		return res, nil
	}
	oType := reflect.TypeOf(o)
	if oType.Kind() != reflect.Map {
		return zero, fmt.Errorf(failedCast, o, o, zero)
	}

	res := make(M)
	resVal := reflect.ValueOf(res)
	oValue := reflect.ValueOf(o)
	for _, keyVal := range oValue.MapKeys() {
		k, err := key(oValue.MapIndex(keyVal).Interface())
		if err != nil {
			return zero, err
		}
		v, err := val(oValue.MapIndex(keyVal).Interface())
		if err != nil {
			return zero, err
		}
		resVal.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
	}
	return res, nil
}
