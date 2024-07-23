package reflectx

import (
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

func IndirectToInterface(a any, ifaces ...any) any {
	if a == nil {
		return nil
	}
	v := reflect.ValueOf(a)
	ifaceTypes := make([]reflect.Type, 0, len(ifaces))
	for _, iface := range ifaces {
		ifaceTypes = append(ifaceTypes, reflect.TypeOf(iface).Elem())
	}
	for {
		for _, ifaceType := range ifaceTypes {
			if v.Type().Implements(ifaceType) {
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
