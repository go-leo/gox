package reflectx

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Accessor interface {
	Get(field string) (any, error)
	Set(field string, val any) error
}

type defaultAccessor struct {
	fields  map[string]reflect.StructField
	address unsafe.Pointer
}

func (accessor *defaultAccessor) Get(field string) (any, error) {
	structField, ok := accessor.fields[field]
	if !ok {
		return nil, fmt.Errorf("reflectx: failed to get %s field", field)
	}
	return accessor.get(structField)
}

func (accessor *defaultAccessor) Set(field string, val any) error {
	structField, ok := accessor.fields[field]
	if !ok {
		return fmt.Errorf("reflectx: failed to get %s field", field)
	}
	accessor.set(structField, val)
	return nil
}

func (accessor *defaultAccessor) fieldAddress(structField reflect.StructField) unsafe.Pointer {
	// field address = objAddress + fieldOffset
	return unsafe.Pointer(uintptr(accessor.address) + structField.Offset)
}

func (accessor *defaultAccessor) get(structField reflect.StructField) (any, error) {
	return reflect.NewAt(structField.Type, accessor.fieldAddress(structField)).Elem().Interface(), nil
}

func (accessor *defaultAccessor) set(structField reflect.StructField, val any) {
	reflect.NewAt(structField.Type, accessor.fieldAddress(structField)).Elem().Set(reflect.ValueOf(val))
}

func FieldAccessorOf(objValue reflect.Value) (Accessor, error) {
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

	return &defaultAccessor{fields: fields, address: address}, nil
}

func TagAccessorOf(objValue reflect.Value, key string) (Accessor, error) {
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

	return &defaultAccessor{fields: fields, address: address}, nil
}
