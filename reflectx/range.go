package reflectx

import (
	"fmt"
	"reflect"
)

func RangeArrayOrSlice(obj any) ([]any, error) {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Array && objValue.Kind() != reflect.Slice {
		return nil, fmt.Errorf("reflectx: %T is not array or slice", obj)
	}
	res := make([]any, 0, objValue.Len())
	for i := 0; i < objValue.Len(); i++ {
		elem := objValue.Index(i)
		res = append(res, elem.Interface())
	}
	return res, nil
}

func RangeMap(obj any) ([]any, []any, error) {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Map {
		return nil, nil, fmt.Errorf("reflectx: %T is not map", obj)
	}
	resKeys := make([]any, 0, objValue.Len())
	resValues := make([]any, 0, objValue.Len())
	itr := objValue.MapRange()
	for itr.Next() {
		resKeys = append(resKeys, itr.Key().Interface())
		resValues = append(resValues, itr.Value().Interface())
	}
	return resKeys, resValues, nil
}
