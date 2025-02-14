package urlx

import (
	"github.com/go-leo/gox/protox"
	"github.com/go-leo/gox/strconvx"
	"golang.org/x/exp/constraints"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/url"
)

func GetInt[Signed constraints.Signed](form url.Values, key string) (Signed, error) {
	if _, ok := form[key]; !ok {
		var v Signed
		return v, nil
	}
	return strconvx.ParseInt[Signed](form.Get(key), 10, 64)
}

func GetIntPtr[Signed constraints.Signed](form url.Values, key string) (*Signed, error) {
	v, err := GetInt[Signed](form, key)
	return &v, err
}

func GetIntSlice[Signed constraints.Signed](form url.Values, key string) ([]Signed, error) {
	if _, ok := form[key]; !ok {
		var v []Signed
		return v, nil
	}
	return strconvx.ParseIntSlice[Signed](form[key], 10, 64)
}

func GetInt32Value(form url.Values, key string) (*wrapperspb.Int32Value, error) {
	v, err := GetInt[int32](form, key)
	return wrapperspb.Int32(v), err
}

func GetInt32ValueSlice(form url.Values, key string) ([]*wrapperspb.Int32Value, error) {
	v, err := GetIntSlice[int32](form, key)
	return protox.WrapInt32Slice(v), err
}

func GetInt64Value(form url.Values, key string) (*wrapperspb.Int64Value, error) {
	v, err := GetInt[int64](form, key)
	return wrapperspb.Int64(v), err
}

func GetInt64ValueSlice(form url.Values, key string) ([]*wrapperspb.Int64Value, error) {
	v, err := GetIntSlice[int64](form, key)
	return protox.WrapInt64Slice(v), err
}
