package reflectx

import (
	"fmt"
	"reflect"
)

// Indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
func Indirect(a any) any {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Pointer {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	return IndirectValue(reflect.ValueOf(a)).Interface()
}

func IndirectValue(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Pointer && !v.IsNil() {
		v = v.Elem()
	}
	return v
}

func IndirectType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t
}

var (
	StringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	ErrorType    = reflect.TypeOf((*error)(nil)).Elem()
)

func IndirectToInterface(a any, ifaces ...reflect.Type) any {
	if a == nil {
		return nil
	}
	v := reflect.ValueOf(a)
	for {
		for _, iface := range ifaces {
			if v.Type().Implements(iface) {
				return v.Interface()
			}
		}
		if v.Kind() == reflect.Pointer && !v.IsNil() {
			v = v.Elem()
		} else {
			return v.Interface()
		}
	}
}
