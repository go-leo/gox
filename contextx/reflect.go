package contextx

import (
	"context"
	"reflect"
)

var ContextType = reflect.TypeOf((*context.Context)(nil)).Elem()
