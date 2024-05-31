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
