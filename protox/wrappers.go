package protox

import "google.golang.org/protobuf/types/known/wrapperspb"

func BoolSlice(s []bool) []*wrapperspb.BoolValue {
	r := make([]*wrapperspb.BoolValue, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Bool(v))
	}
	return r
}

func Int32Slice(s []int32) []*wrapperspb.Int32Value {
	r := make([]*wrapperspb.Int32Value, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Int32(v))
	}
	return r
}

func Int64Slice(s []int64) []*wrapperspb.Int64Value {
	r := make([]*wrapperspb.Int64Value, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Int64(v))
	}
	return r
}

func Float32Slice(s []float32) []*wrapperspb.FloatValue {
	r := make([]*wrapperspb.FloatValue, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Float(v))
	}
	return r
}

func Float64Slice(s []float64) []*wrapperspb.DoubleValue {
	r := make([]*wrapperspb.DoubleValue, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Double(v))
	}
	return r
}

func Uint32Slice(s []uint32) []*wrapperspb.UInt32Value {
	r := make([]*wrapperspb.UInt32Value, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.UInt32(v))
	}
	return r
}

func Uint64Slice(s []uint64) []*wrapperspb.UInt64Value {
	r := make([]*wrapperspb.UInt64Value, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.UInt64(v))
	}
	return r
}

func StringSlice(s []string) []*wrapperspb.StringValue {
	r := make([]*wrapperspb.StringValue, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.String(v))
	}
	return r
}

func BytesSlice(s [][]byte) []*wrapperspb.BytesValue {
	r := make([]*wrapperspb.BytesValue, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Bytes(v))
	}
	return r
}
