package errorx

import "reflect"

var ErrorType = reflect.TypeOf((*error)(nil)).Elem()
