package convx

import (
	"database/sql/driver"
	"encoding"
	"fmt"
	"reflect"
	"time"
)

type (
	// int64er json.Number
	int64er interface{ Int64() (int64, error) }
	// float64er json.Number
	float64er interface{ Float64() (float64, error) }
	// asDurationer  google.golang.org/protobuf/types/known/durationpb.Duration
	asDurationer interface{ AsDuration() time.Duration }
	// asTimeer  google.golang.org/protobuf/types/known/timestamppb.Timestamp
	asTimeer interface{ AsTime() time.Time }
)

var (
	emptyInt64er       = reflect.TypeOf((*int64er)(nil)).Elem()
	emptyFloat64er     = reflect.TypeOf((*float64er)(nil)).Elem()
	emptyValuer        = reflect.TypeOf((*driver.Valuer)(nil)).Elem()
	emptyAsDurationer  = reflect.TypeOf((*asDurationer)(nil)).Elem()
	empryAsTimeer      = reflect.TypeOf((*asTimeer)(nil)).Elem()
	emptyErrorer       = reflect.TypeOf((*error)(nil)).Elem()
	emptyStringer      = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	emptyTextMarshaler = reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()
)
