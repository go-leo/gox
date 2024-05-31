package protox

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"testing"
)

func TestMessageSlice(t *testing.T) {
	messages := MessageSlice[[]*wrapperspb.BoolValue]([]*wrapperspb.BoolValue{
		wrapperspb.Bool(true),
		wrapperspb.Bool(false),
	})
	t.Log(messages)
}

func TestProtoSlice(t *testing.T) {
	protos := ProtoSlice[[]*wrapperspb.BoolValue]([]proto.Message{
		wrapperspb.Bool(true),
		wrapperspb.Bool(false),
		wrapperspb.String("true"),
	})
	t.Log(protos)
}
