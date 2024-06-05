package protox

import "google.golang.org/protobuf/proto"

func Clone[M proto.Message](m M) M {
	var cloned M
	if m == nil {
		return cloned
	}
	cloned = proto.Clone(m).(M)
	return cloned
}
