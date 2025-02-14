package urlx

import (
	"github.com/go-leo/gox/protox"
	"github.com/go-leo/gox/strconvx"
	"golang.org/x/exp/constraints"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/url"
)

func GetUint[Unsigned constraints.Unsigned](form url.Values, key string) (Unsigned, error) {
	if _, ok := form[key]; !ok {
		var v Unsigned
		return v, nil
	}
	return strconvx.ParseUint[Unsigned](form.Get(key), 10, 64)
}

func GetUintPtr[Unsigned constraints.Unsigned](form url.Values, key string) (*Unsigned, error) {
	v, err := GetUint[Unsigned](form, key)
	return &v, err
}

func GetUintSlice[Unsigned constraints.Unsigned](form url.Values, key string) ([]Unsigned, error) {
	if _, ok := form[key]; !ok {
		var v []Unsigned
		return v, nil
	}
	return strconvx.ParseUintSlice[Unsigned](form[key], 10, 64)
}

func GetUint32Value(form url.Values, key string) (*wrapperspb.UInt32Value, error) {
	v, err := GetUint[uint32](form, key)
	return wrapperspb.UInt32(v), err
}

func GetUint32ValueSlice(form url.Values, key string) ([]*wrapperspb.UInt32Value, error) {
	v, err := GetUintSlice[uint32](form, key)
	return protox.WrapUint32Slice(v), err
}

func GetUint64Value(form url.Values, key string) (*wrapperspb.UInt64Value, error) {
	v, err := GetUint[uint64](form, key)
	return wrapperspb.UInt64(v), err
}

func GetUint64ValueSlice(form url.Values, key string) ([]*wrapperspb.UInt64Value, error) {
	v, err := GetUintSlice[uint64](form, key)
	return protox.WrapUint64Slice(v), err
}
