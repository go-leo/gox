package condx

import (
	"github.com/go-leo/gox/strconvx"
	"reflect"
)

// DecodeHookFuncType is a DecodeHookFunc which has complete information about
// the source and target types.
type DecodeHookFuncType func(reflect.Type, reflect.Type, interface{}) (interface{}, error)

// DecodeHookFuncKind is a DecodeHookFunc which knows only the Kinds of the
// source and target types.
type DecodeHookFuncKind func(reflect.Kind, reflect.Kind, interface{}) (interface{}, error)

// DecodeHookFuncValue is a DecodeHookFunc which has complete access to both the source and target
// values.
type DecodeHookFuncValue func(from reflect.Value, to reflect.Value) (interface{}, error)

type Stringfy interface {
	ToString(value any) string
}

var _ Stringfy = (*BoolStringfy)(nil)

type BoolStringfy struct{}

func (s BoolStringfy) ToString(value any) string {
	return strconvx.FormatBool(value.(bool))
}

//Bool
//Int
//Int8
//Int16
//Int32
//Int64
//Uint
//Uint8
//Uint16
//Uint32
//Uint64
//Uintptr
//Float32
//Float64
//Complex64
//Complex128
//Array
//Chan
//Func
//Interface
//Map
//Pointer
//Slice
//String
//Struct
//UnsafePointer
