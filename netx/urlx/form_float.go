package urlx

import (
	"github.com/go-leo/gox/protox"
	"github.com/go-leo/gox/strconvx"
	"golang.org/x/exp/constraints"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/url"
)

func GetFloat[Float constraints.Float](form url.Values, key string) (Float, error) {
	if _, ok := form[key]; !ok {
		var v Float
		return v, nil
	}
	return strconvx.ParseFloat[Float](form.Get(key), 64)
}

func GetFloatPtr[Float constraints.Float](form url.Values, key string) (*Float, error) {
	v, err := GetFloat[Float](form, key)
	return &v, err
}

func GetFloatSlice[Float constraints.Float](form url.Values, key string) ([]Float, error) {
	if _, ok := form[key]; !ok {
		var v []Float
		return v, nil
	}
	return strconvx.ParseFloatSlice[Float](form[key], 64)
}

func GetFloat32Value(form url.Values, key string) (*wrapperspb.FloatValue, error) {
	v, err := GetFloat[float32](form, key)
	return wrapperspb.Float(v), err
}

func GetFloat32ValueSlice(form url.Values, key string) ([]*wrapperspb.FloatValue, error) {
	v, err := GetFloatSlice[float32](form, key)
	return protox.WrapFloat32Slice(v), err
}

func GetFloat64Value(form url.Values, key string) (*wrapperspb.DoubleValue, error) {
	v, err := GetFloat[float64](form, key)
	return wrapperspb.Double(v), err
}

func GetFloat64ValueSlice(form url.Values, key string) ([]*wrapperspb.DoubleValue, error) {
	v, err := GetFloatSlice[float64](form, key)
	return protox.WrapFloat64Slice(v), err
}
