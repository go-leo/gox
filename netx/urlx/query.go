package urlx

import (
	"github.com/go-leo/gox/protox"
	"github.com/go-leo/gox/strconvx"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/url"
)

func GetBool(queries url.Values, key string) func() (bool, error) {
	return func() (bool, error) {
		return strconvx.ParseBool(queries.Get(key))
	}
}

func GetBoolPtr(queries url.Values, key string) func() (*bool, error) {
	return func() (*bool, error) {
		v, err := strconvx.ParseBool(queries.Get(key))
		return proto.Bool(v), err
	}
}

func GetBoolSlice(queries url.Values, key string) func() ([]bool, error) {
	return func() ([]bool, error) {
		return strconvx.ParseBoolSlice(queries[key])
	}
}

func GetBoolValue(queries url.Values, key string) func() (*wrapperspb.BoolValue, error) {
	return func() (*wrapperspb.BoolValue, error) {
		v, err := strconvx.ParseBool(queries.Get(key))
		return wrapperspb.Bool(v), err
	}
}

func GetBoolValueSlice(queries url.Values, key string) func() ([]*wrapperspb.BoolValue, error) {
	return func() ([]*wrapperspb.BoolValue, error) {
		v, err := strconvx.ParseBoolSlice(queries[key])
		return protox.WrapBoolSlice(v), err
	}
}
