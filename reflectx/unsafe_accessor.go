package reflectx

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Accessor struct {
	fields  map[string]reflect.StructField
	address unsafe.Pointer
}

func (accessor *Accessor) Get(field string) (any, error) {
	structField, ok := accessor.fields[field]
	if !ok {
		return nil, fmt.Errorf("reflectx: failed to get %s field", field)
	}
	return accessor.get(structField)
}

func (accessor *Accessor) Set(field string, val any) error {
	structField, ok := accessor.fields[field]
	if !ok {
		return fmt.Errorf("reflectx: failed to get %s field", field)
	}
	accessor.set(structField, val)
	return nil
}

func (accessor *Accessor) fieldAddress(structField reflect.StructField) unsafe.Pointer {
	// field address = objAddress + fieldOffset
	return unsafe.Pointer(uintptr(accessor.address) + structField.Offset)
}

func (accessor *Accessor) get(structField reflect.StructField) (any, error) {
	return reflect.NewAt(structField.Type, accessor.fieldAddress(structField)).Elem().Interface(), nil
}

func (accessor *Accessor) set(structField reflect.StructField, val any) {
	reflect.NewAt(structField.Type, accessor.fieldAddress(structField)).Elem().Set(reflect.ValueOf(val))
}

func FieldAccessorOf(objValue reflect.Value) (*Accessor, error) {
	if !objValue.IsValid() {
		return nil, fmt.Errorf("reflectx: invalid reflect.Value")
	}
	if objValue.Kind() != reflect.Pointer {
		return nil, fmt.Errorf("reflectx: %T is not a pointer", objValue.Interface())
	}
	objStructValue := IndirectValue(objValue)

	// Ensure the value points to a struct or struct field.
	if objStructValue.Kind() != reflect.Struct {
		return nil, fmt.Errorf("reflectx: %T does not point to a struct", objValue.Interface())
	}

	objStructType := objStructValue.Type()
	address := objStructValue.Addr().UnsafePointer()

	numField := objStructType.NumField()
	fields := make(map[string]reflect.StructField, numField)
	for i := 0; i < numField; i++ {
		structField := objStructType.Field(i)
		fields[structField.Name] = structField
	}

	return &Accessor{fields: fields, address: address}, nil
}

func TagAccessorOf(objValue reflect.Value, key string) (*Accessor, error) {
	if !objValue.IsValid() {
		return nil, fmt.Errorf("reflectx: invalid reflect.Value")
	}
	if objValue.Kind() != reflect.Pointer {
		return nil, fmt.Errorf("reflectx: %T is not a pointer", objValue.Interface())
	}
	objStructValue := IndirectValue(objValue)

	// Ensure the value points to a struct or struct field.
	if objStructValue.Kind() != reflect.Struct {
		return nil, fmt.Errorf("reflectx: %T does not point to a struct", objValue.Interface())
	}

	objStructType := objStructValue.Type()
	address := objStructValue.Addr().UnsafePointer()

	numField := objStructType.NumField()
	fields := make(map[string]reflect.StructField, numField)
	for i := 0; i < numField; i++ {
		structField := objStructType.Field(i)
		value, ok := structField.Tag.Lookup(key)
		if !ok {
			continue
		}
		if _, ok := fields[value]; ok {
			return nil, fmt.Errorf("reflectx: %s tag value %s is duplicated", key, value)
		}
		fields[value] = structField
	}

	return &Accessor{fields: fields, address: address}, nil
}

func GetByAccessor[T any](accessor *Accessor, field string) (T, error) {
	var res T
	structField, ok := accessor.fields[field]
	if !ok {
		return res, fmt.Errorf("reflectx: failed to get %s field", field)
	}
	res = *(*T)(accessor.fieldAddress(structField))
	return res, nil
}

func SetByAccessor[T any](accessor *Accessor, field string, val T) error {
	structField, ok := accessor.fields[field]
	if !ok {
		return fmt.Errorf("reflectx: failed to get %s field", field)
	}
	*(*T)(accessor.fieldAddress(structField)) = val
	return nil
}
