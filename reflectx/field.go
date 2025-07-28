package reflectx

import (
	"fmt"
	"reflect"
	"unsafe"
)

// FindFieldByTag searches through the fields of a given struct value for a field
// that has a specific tag and whose tag value satisfies a given condition.
// It returns the reflect.Value of the matched field along with a boolean indicating
// whether a match was found.
//
// Parameters:
//
//	v: The reflect.Value representing the struct to search.
//	tagKey: The key of the tag to look for.
//	match: A function that takes the tag value as a string and returns true if it satisfies the condition.
//
// Returns:
//
//	A tuple containing the reflect.Value of the matched field and a boolean indicating if a match was found.
func FindFieldByTag(objValue reflect.Value, tagKey string, match func(tagVal string) bool) (reflect.Value, bool) {
	// Indirect the value to get the underlying value.
	structValue := IndirectValue(objValue)

	// Check if the value is a struct.
	if structValue.Kind() != reflect.Struct {
		return reflect.Value{}, false
	}

	// Directly access the type once instead of on each iteration.
	structType := structValue.Type()

	// Iterate over all fields in the given struct.
	for i := 0; i < structType.NumField(); i++ {
		// Get the current field and its corresponding struct field type.
		field := structValue.Field(i)
		structField := structType.Field(i)

		// Check if the field has the specified tag with the given value.
		if tagVal, ok := structField.Tag.Lookup(tagKey); ok && match(tagVal) {
			// Return the field value if the tag matches.
			return field, true
		}
	}

	// Return zero Value if no matching field is found.
	return reflect.Value{}, false
}

func GetField(objValue reflect.Value, field string) (any, error) {
	structValue := IndirectValue(objValue)
	for structValue.Kind() != reflect.Struct {
		return nil, fmt.Errorf("reflectx: %T is not struct", structValue.Interface())
	}
	fieldVal := structValue.FieldByName(field)
	if fieldVal.IsZero() {
		return nil, fmt.Errorf("reflectx: field %s not found", field)
	}
	return fieldVal.Interface(), nil
}

// GetUnexportedField 通过反射获取结构体中的非导出字段值，支持嵌套字段访问
// objValue: 包含目标字段的对象的反射值
// fields: 要访问的字段名称序列，支持多级嵌套访问
// 返回值: 目标字段的反射值，如果字段不存在则返回无效的reflect.Value
func GetUnexportedField(objValue reflect.Value, fields ...string) reflect.Value {
	// 遍历字段路径，逐级深入获取目标字段
	for _, field := range fields {
		v := IndirectValue(objValue)
		field := v.FieldByName(field)
		if !field.IsValid() {
			return reflect.Value{}
		}
		objValue = field
	}
	// 通过不安全指针创建可访问的字段值副本
	return reflect.NewAt(objValue.Type(), unsafe.Pointer(objValue.UnsafeAddr())).Elem()
}

func SetField(objValue reflect.Value, field string, newValue any) error {
	if objValue.Kind() != reflect.Pointer {
		return fmt.Errorf("reflectx: %T is not pointer", objValue.Interface())
	}
	structValue := IndirectValue(objValue)
	for structValue.Kind() != reflect.Struct {
		return fmt.Errorf("reflectx: %T is not struct", structValue.Interface())
	}
	fieldVal := structValue.FieldByName(field)
	if fieldVal.IsZero() {
		return fmt.Errorf("reflectx: field %s not found", field)
	}
	if !fieldVal.CanSet() {
		return fmt.Errorf("reflectx: cannot set field %s", field)
	}
	fieldVal.Set(reflect.ValueOf(newValue))
	return nil
}

func RangeFields(objValue reflect.Value) (map[string]any, error) {
	if objValue.IsZero() {
		return nil, fmt.Errorf("reflectx: unsupport zero value")
	}
	objValue = IndirectValue(objValue)
	objType := IndirectType(objValue.Type())
	if objType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("reflectx: %s is not struct", objType.Kind())
	}

	numField := objType.NumField()
	res := make(map[string]any, numField)
	for i := 0; i < numField; i++ {
		fieldType := objType.Field(i)
		fieldValue := objValue.Field(i)
		if fieldType.IsExported() {
			res[fieldType.Name] = fieldValue.Interface()
		} else {
			res[fieldType.Name] = reflect.Zero(fieldType.Type).Interface()
		}
	}
	return res, nil
}
