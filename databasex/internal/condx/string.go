package condx

import (
	"github.com/go-leo/gox/strconvx"
	"reflect"
)

type DecodeHookFuncType func(reflect.Type, any) (string, error)

type DecodeHookFuncKind func(reflect.Kind, any) (string, error)

type DecodeHookFuncValue func(reflect.Value) (string, error)

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
