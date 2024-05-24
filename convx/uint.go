package convx

import (
	"fmt"
	"reflect"
)

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

// ToUintE converts an interface to a uint type.
func ToUintE(i any) (uint, error) {
	return ToUnsignedE[uint](i)
}

// ToUint64E converts an interface to a uint64 type.
func ToUint64E(i any) (uint64, error) {
	return ToUnsignedE[uint64](i)
}

// ToUint32E converts an interface to a uint32 type.
func ToUint32E(i any) (uint32, error) {
	return ToUnsignedE[uint32](i)
}

// ToUint16E converts an interface to a uint16 type.
func ToUint16E(i any) (uint16, error) {
	return ToUnsignedE[uint16](i)
}

// ToUint8E converts an interface to a uint type.
func ToUint8E(i any) (uint8, error) {
	return ToUnsignedE[uint8](i)
}

// ToUintSlice casts an interface to a []uint type.
func ToUintSlice(i interface{}) []uint {
	v, _ := ToUintSliceE(i)
	return v
}

// ToUint64Slice casts an interface to a []uint64 type.
func ToUint64Slice(i interface{}) []uint64 {
	v, _ := ToUint64SliceE(i)
	return v
}

// ToUint32Slice casts an interface to a []uint32 type.
func ToUint32Slice(i interface{}) []uint32 {
	v, _ := ToUint32SliceE(i)
	return v
}

// ToUintSliceE casts an interface to a []uint type.
func ToUintSliceE(i interface{}) ([]uint, error) {
	if i == nil {
		return []uint{}, fmt.Errorf("unable to cast %#v of type %T to []uint", i, i)
	}

	switch v := i.(type) {
	case []uint:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]uint, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToUintE(s.Index(j).Interface())
			if err != nil {
				return []uint{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []uint{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
	}
}

// ToUint64SliceE casts an interface to a []uint64 type.
func ToUint64SliceE(i interface{}) ([]uint64, error) {
	if i == nil {
		return []uint64{}, fmt.Errorf("unable to cast %#v of type %T to []uint64", i, i)
	}

	switch v := i.(type) {
	case []uint64:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]uint64, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToUint64E(s.Index(j).Interface())
			if err != nil {
				return []uint64{}, fmt.Errorf("unable to cast %#v of type %T to []int64", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []uint64{}, fmt.Errorf("unable to cast %#v of type %T to []int64", i, i)
	}
}

// ToUint32SliceE casts an interface to a []int32 type.
func ToUint32SliceE(i interface{}) ([]uint32, error) {
	if i == nil {
		return []uint32{}, fmt.Errorf("unable to cast %#v of type %T to []int32", i, i)
	}

	switch v := i.(type) {
	case []uint32:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]uint32, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToUint32E(s.Index(j).Interface())
			if err != nil {
				return []uint32{}, fmt.Errorf("unable to cast %#v of type %T to []int32", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []uint32{}, fmt.Errorf("unable to cast %#v of type %T to []int32", i, i)
	}
}
