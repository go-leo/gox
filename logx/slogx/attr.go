package slogx

import (
	"encoding/json"
	"log/slog"
	"time"

	"golang.org/x/exp/constraints"
)

func Bool[Bool ~bool](key string, value Bool) slog.Attr {
	return slog.Attr{Key: key, Value: slog.BoolValue(bool(value))}
}

func Int[Int constraints.Signed](key string, value Int) slog.Attr {
	return slog.Attr{Key: key, Value: slog.Int64Value(int64(value))}
}

func Uint[Uint constraints.Unsigned](key string, value Uint) slog.Attr {
	return slog.Attr{Key: key, Value: slog.Uint64Value(uint64(value))}
}

func Duration[Duration time.Duration](key string, value Duration) slog.Attr {
	return slog.Attr{Key: key, Value: slog.DurationValue(time.Duration(value))}
}

func Float[Float constraints.Float](key string, value Float) slog.Attr {
	return slog.Attr{Key: key, Value: slog.Float64Value(float64(value))}
}

func String[String ~string](key string, value String) slog.Attr {
	return slog.Attr{Key: key, Value: slog.StringValue(string(value))}
}

func Time[Time time.Time](key string, value Time) slog.Attr {
	return slog.Attr{Key: key, Value: slog.TimeValue(time.Time(value))}
}

func Json(key string, value any) slog.Attr {
	data, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return slog.Attr{Key: key, Value: slog.StringValue(string(data))}
}

func Error(key string, value error) slog.Attr {
	return String(key, value.Error())
}

// func Valuer(key string, value slog.LogValuer) slog.Attr {
// 	return slog.Attr{Key: key, Value: value.LogValue()}
// }

// KindGroup
// KindLogValuer
