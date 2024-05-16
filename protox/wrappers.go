package protox

import "google.golang.org/protobuf/types/known/wrapperspb"

func BoolSlice(s []bool) []*wrapperspb.BoolValue {
	r := make([]*wrapperspb.BoolValue, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Bool(v))
	}
	return r
}

func UnwrapBoolSlice(s []*wrapperspb.BoolValue) []bool {
	r := make([]bool, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
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

func UnwrapInt32Slice(s []*wrapperspb.Int32Value) []int32 {
	r := make([]int32, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
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

func UnwrapInt64Slice(s []*wrapperspb.Int64Value) []int64 {
	r := make([]int64, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
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

func UnwrapFloat32Slice(s []*wrapperspb.FloatValue) []float32 {
	r := make([]float32, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
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

func UnwrapFloat64Slice(s []*wrapperspb.DoubleValue) []float64 {
	r := make([]float64, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}

func Uint32Slice(s []*wrapperspb.UInt32Value) []uint32 {
	r := make([]uint32, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
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

func UnwrapUint64Slice(s []*wrapperspb.UInt64Value) []uint64 {
	r := make([]uint64, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
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

func UnwrapStringSlice(s []*wrapperspb.StringValue) []string {
	r := make([]string, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
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

func UnwrapBytesSlice(s []*wrapperspb.BytesValue) [][]byte {
	r := make([][]byte, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}
