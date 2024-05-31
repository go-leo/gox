package protox

import "google.golang.org/protobuf/proto"

func MessageSlice[S []E, E proto.Message](s S) []proto.Message {
	if s == nil {
		return nil
	}
	r := make([]proto.Message, 0, len(s))
	for _, e := range s {
		r = append(r, e)
	}
	return r
}

func ProtoSlice[S []E, E proto.Message](s []proto.Message) S {
	if s == nil {
		return nil
	}
	r := make(S, 0, len(s))
	for _, e := range s {
		p, ok := e.(E)
		if !ok {
			continue
		}
		r = append(r, p)
	}
	return r
}
