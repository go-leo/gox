package convx

import (
	"fmt"
	"reflect"
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
