package reflectx

import (
	"fmt"
	"reflect"
	"unsafe"
)

type UnsafeAccessor struct {
	fields  map[string]reflect.StructField
	address unsafe.Pointer
}

func UnsafeAccessorOf(obj any) (*UnsafeAccessor, error) {
	objType := reflect.TypeOf(obj)
	if objType.Kind() != reflect.Pointer {
		return nil, fmt.Errorf("reflectx: %T is not pointer", obj)
	}
	objStructType := IndirectType(objType)
	numField := objStructType.NumField()
	fields := make(map[string]reflect.StructField, numField)
	for i := 0; i < numField; i++ {
		structField := objStructType.Field(i)
		fields[structField.Name] = structField
	}
	objValue := reflect.ValueOf(obj)
	return &UnsafeAccessor{fields: fields, address: objValue.UnsafePointer()}, nil
}

func (accessor *UnsafeAccessor) Field(field string) (any, error) {
	structField, ok := accessor.fields[field]
	if !ok {
		return nil, fmt.Errorf("reflectx: failed to get %s field", field)
	}
	// field address = objAddress + fieldOffset
	fdAddress := unsafe.Pointer(uintptr(accessor.address) + structField.Offset)
	return reflect.NewAt(structField.Type, fdAddress).Elem().Interface(), nil
}

func (accessor *UnsafeAccessor) SetField(field string, val any) error {
	structField, ok := accessor.fields[field]
	if !ok {
		return fmt.Errorf("reflectx: failed to get %s field", field)
	}
	// field address = objAddress + fieldOffset
	fdAddress := unsafe.Pointer(uintptr(accessor.address) + structField.Offset)
	reflect.NewAt(structField.Type, fdAddress).Elem().Set(reflect.ValueOf(val))
	return nil
}
