package protox

import "google.golang.org/protobuf/types/known/wrapperspb"

func WrapBoolSlice(s []bool) []*wrapperspb.BoolValue {
	if s == nil {
		return nil
	}
	r := make([]*wrapperspb.BoolValue, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Bool(v))
	}
	return r
}

func UnwrapBoolSlice(s []*wrapperspb.BoolValue) []bool {
	if s == nil {
		return nil
	}
	r := make([]bool, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}

func WrapInt32Slice(s []int32) []*wrapperspb.Int32Value {
	if s == nil {
		return nil
	}
	r := make([]*wrapperspb.Int32Value, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Int32(v))
	}
	return r
}

func UnwrapInt32Slice(s []*wrapperspb.Int32Value) []int32 {
	if s == nil {
		return nil
	}
	r := make([]int32, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}

func WrapInt64Slice(s []int64) []*wrapperspb.Int64Value {
	if s == nil {
		return nil
	}
	r := make([]*wrapperspb.Int64Value, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Int64(v))
	}
	return r
}

func UnwrapInt64Slice(s []*wrapperspb.Int64Value) []int64 {
	if s == nil {
		return nil
	}
	r := make([]int64, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}

func WrapFloat32Slice(s []float32) []*wrapperspb.FloatValue {
	if s == nil {
		return nil
	}
	r := make([]*wrapperspb.FloatValue, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Float(v))
	}
	return r
}

func UnwrapFloat32Slice(s []*wrapperspb.FloatValue) []float32 {
	if s == nil {
		return nil
	}
	r := make([]float32, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}

func WrapFloat64Slice(s []float64) []*wrapperspb.DoubleValue {
	if s == nil {
		return nil
	}
	r := make([]*wrapperspb.DoubleValue, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Double(v))
	}
	return r
}

func UnwrapFloat64Slice(s []*wrapperspb.DoubleValue) []float64 {
	if s == nil {
		return nil
	}
	r := make([]float64, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}

func WrapUint32Slice(s []uint32) []*wrapperspb.UInt32Value {
	if s == nil {
		return nil
	}
	r := make([]*wrapperspb.UInt32Value, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.UInt32(v))
	}
	return r
}

func UnwrapUint32Slice(s []*wrapperspb.UInt32Value) []uint32 {
	if s == nil {
		return nil
	}
	r := make([]uint32, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}

func WrapUint64Slice(s []uint64) []*wrapperspb.UInt64Value {
	if s == nil {
		return nil
	}
	r := make([]*wrapperspb.UInt64Value, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.UInt64(v))
	}
	return r
}

func UnwrapUint64Slice(s []*wrapperspb.UInt64Value) []uint64 {
	if s == nil {
		return nil
	}
	r := make([]uint64, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}

func WrapStringSlice(s []string) []*wrapperspb.StringValue {
	if s == nil {
		return nil
	}
	r := make([]*wrapperspb.StringValue, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.String(v))
	}
	return r
}

func UnwrapStringSlice(s []*wrapperspb.StringValue) []string {
	if s == nil {
		return nil
	}
	r := make([]string, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}

func WrapBytesSlice(s [][]byte) []*wrapperspb.BytesValue {
	if s == nil {
		return nil
	}
	r := make([]*wrapperspb.BytesValue, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Bytes(v))
	}
	return r
}

func UnwrapBytesSlice(s []*wrapperspb.BytesValue) [][]byte {
	if s == nil {
		return nil
	}
	r := make([][]byte, 0, len(s))
	for _, v := range s {
		r = append(r, v.GetValue())
	}
	return r
}
