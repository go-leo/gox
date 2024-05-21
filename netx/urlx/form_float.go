package urlx

import (
	"github.com/go-leo/gox/protox"
	"github.com/go-leo/gox/strconvx"
	"golang.org/x/exp/constraints"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/url"
)

func GetFloat[Float constraints.Float](queries url.Values, key string) func() (Float, error) {
	return func() (Float, error) {
		if _, ok := queries[key]; !ok {
			var v Float
			return v, nil
		}
		return strconvx.ParseFloat[Float](queries.Get(key), 64)
	}
}

func GetFloatPtr[Float constraints.Float](queries url.Values, key string) func() (*Float, error) {
	return func() (*Float, error) {
		v, err := GetFloat[Float](queries, key)()
		return &v, err
	}
}

func GetFloatSlice[Float constraints.Float](queries url.Values, key string) func() ([]Float, error) {
	return func() ([]Float, error) {
		if _, ok := queries[key]; !ok {
			var v []Float
			return v, nil
		}
		return strconvx.ParseFloatSlice[Float](queries[key], 64)
	}
}

func GetFloat32Value(queries url.Values, key string) func() (*wrapperspb.FloatValue, error) {
	return func() (*wrapperspb.FloatValue, error) {
		v, err := GetFloat[float32](queries, key)()
		return wrapperspb.Float(v), err
	}
}

func GetFloat32ValueSlice(queries url.Values, key string) func() ([]*wrapperspb.FloatValue, error) {
	return func() ([]*wrapperspb.FloatValue, error) {
		v, err := GetFloatSlice[float32](queries, key)()
		return protox.WrapFloat32Slice(v), err
	}
}

func GetFloat64Value(queries url.Values, key string) func() (*wrapperspb.DoubleValue, error) {
	return func() (*wrapperspb.DoubleValue, error) {
		v, err := GetFloat[float64](queries, key)()
		return wrapperspb.Double(v), err
	}
}

func GetFloat64ValueSlice(queries url.Values, key string) func() ([]*wrapperspb.DoubleValue, error) {
	return func() ([]*wrapperspb.DoubleValue, error) {
		v, err := GetFloatSlice[float64](queries, key)()
		return protox.WrapFloat64Slice(v), err
	}
}
