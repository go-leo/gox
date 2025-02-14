package urlx

import (
	"github.com/go-leo/gox/protox"
	"github.com/go-leo/gox/strconvx"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/url"
)

func GetBool(form url.Values, key string) (bool, error) {
	if _, ok := form[key]; !ok {
		return false, nil
	}
	return strconvx.ParseBool(form.Get(key))
}

func GetBoolPtr(form url.Values, key string) (*bool, error) {
	v, err := GetBool(form, key)
	return &v, err
}

func GetBoolSlice(form url.Values, key string) ([]bool, error) {
	if _, ok := form[key]; !ok {
		return nil, nil
	}
	return strconvx.ParseBoolSlice(form[key])
}

func GetBoolValue(form url.Values, key string) (*wrapperspb.BoolValue, error) {
	v, err := strconvx.ParseBool(form.Get(key))
	return wrapperspb.Bool(v), err
}

func GetBoolValueSlice(form url.Values, key string) ([]*wrapperspb.BoolValue, error) {
	v, err := strconvx.ParseBoolSlice(form[key])
	return protox.WrapBoolSlice(v), err
}
