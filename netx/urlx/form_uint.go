package urlx

import (
	"github.com/go-leo/gox/protox"
	"github.com/go-leo/gox/strconvx"
	"golang.org/x/exp/constraints"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/url"
)

func GetUint[Unsigned constraints.Unsigned](queries url.Values, key string) func() (Unsigned, error) {
	return func() (Unsigned, error) {
		if _, ok := queries[key]; !ok {
			var v Unsigned
			return v, nil
		}
		return strconvx.ParseUint[Unsigned](queries.Get(key), 10, 64)
	}
}

func GetUintPtr[Unsigned constraints.Unsigned](queries url.Values, key string) func() (*Unsigned, error) {
	return func() (*Unsigned, error) {
		v, err := GetUint[Unsigned](queries, key)()
		return &v, err
	}
}

func GetUintSlice[Unsigned constraints.Unsigned](queries url.Values, key string) func() ([]Unsigned, error) {
	return func() ([]Unsigned, error) {
		if _, ok := queries[key]; !ok {
			var v []Unsigned
			return v, nil
		}
		return strconvx.ParseUintSlice[Unsigned](queries[key], 10, 64)
	}
}

func GetUint32Value(queries url.Values, key string) func() (*wrapperspb.UInt32Value, error) {
	return func() (*wrapperspb.UInt32Value, error) {
		v, err := GetUint[uint32](queries, key)()
		return wrapperspb.UInt32(v), err
	}
}

func GetUint32ValueSlice(queries url.Values, key string) func() ([]*wrapperspb.UInt32Value, error) {
	return func() ([]*wrapperspb.UInt32Value, error) {
		v, err := GetUintSlice[uint32](queries, key)()
		return protox.WrapUint32Slice(v), err
	}
}

func GetUint64Value(queries url.Values, key string) func() (*wrapperspb.UInt64Value, error) {
	return func() (*wrapperspb.UInt64Value, error) {
		v, err := GetUint[uint64](queries, key)()
		return wrapperspb.UInt64(v), err
	}
}

func GetUint64ValueSlice(queries url.Values, key string) func() ([]*wrapperspb.UInt64Value, error) {
	return func() ([]*wrapperspb.UInt64Value, error) {
		v, err := GetUintSlice[uint64](queries, key)()
		return protox.WrapUint64Slice(v), err
	}
}
