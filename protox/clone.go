package protox

import "google.golang.org/protobuf/proto"

func Clone[M proto.Message](m M) M {
	return proto.Clone(m).(M)
}
