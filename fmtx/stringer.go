package fmtx

import (
	"fmt"
	"reflect"
)

var StringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
