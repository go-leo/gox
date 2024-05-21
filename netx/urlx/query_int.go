package urlx

import (
	"github.com/go-leo/gox/protox"
	"github.com/go-leo/gox/strconvx"
	"golang.org/x/exp/constraints"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/url"
)

func GetInt[Signed constraints.Signed](queries url.Values, key string) func() (Signed, error) {
	return func() (Signed, error) {
		if _, ok := queries[key]; !ok {
			var v Signed
			return v, nil
		}
		return strconvx.ParseInt[Signed](queries.Get(key), 10, 64)
	}
}

func GetIntPtr[Signed constraints.Signed](queries url.Values, key string) func() (*Signed, error) {
	return func() (*Signed, error) {
		v, err := GetInt[Signed](queries, key)()
		return &v, err
	}
}

func GetIntSlice[Signed constraints.Signed](queries url.Values, key string) func() ([]Signed, error) {
	return func() ([]Signed, error) {
		if _, ok := queries[key]; !ok {
			var v []Signed
			return v, nil
		}
		return strconvx.ParseIntSlice[Signed](queries[key], 10, 64)
	}
}

func GetInt32Value(queries url.Values, key string) func() (*wrapperspb.Int32Value, error) {
	return func() (*wrapperspb.Int32Value, error) {
		v, err := GetInt[int32](queries, key)()
		return wrapperspb.Int32(v), err
	}
}

func GetInt32ValueSlice(queries url.Values, key string) func() ([]*wrapperspb.Int32Value, error) {
	return func() ([]*wrapperspb.Int32Value, error) {
		v, err := GetIntSlice[int32](queries, key)()
		return protox.WrapInt32Slice(v), err
	}
}

func GetInt64Value(queries url.Values, key string) func() (*wrapperspb.Int64Value, error) {
	return func() (*wrapperspb.Int64Value, error) {
		v, err := GetInt[int64](queries, key)()
		return wrapperspb.Int64(v), err
	}
}

func GetInt64ValueSlice(queries url.Values, key string) func() ([]*wrapperspb.Int64Value, error) {
	return func() ([]*wrapperspb.Int64Value, error) {
		v, err := GetIntSlice[int64](queries, key)()
		return protox.WrapInt64Slice(v), err
	}
}
