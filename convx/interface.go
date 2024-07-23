package convx

import (
	"database/sql/driver"
	"encoding"
	"fmt"
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
	emptyInt64er       = (*int64er)(nil)
	emptyFloat64er     = (*float64er)(nil)
	emptyValuer        = (*driver.Valuer)(nil)
	emptyAsDurationer  = (*asDurationer)(nil)
	empryAsTimeer      = (*asTimeer)(nil)
	emptyErrorer       = (*error)(nil)
	emptyStringer      = (*fmt.Stringer)(nil)
	emptyTextMarshaler = (*encoding.TextMarshaler)(nil)
)
