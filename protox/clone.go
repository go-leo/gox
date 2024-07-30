package protox

import "google.golang.org/protobuf/proto"

func Clone[M proto.Message](m M) M {
	return proto.Clone(m).(M)
}

func CloneSlice[S ~[]M, M proto.Message](s S) S {
	var zero S
	if s == nil {
		return zero
	}
	r := make(S, 0, len(s))
	for _, m := range s {
		r = append(r, Clone(m))
	}
	return r
}
