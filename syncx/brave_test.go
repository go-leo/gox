package syncx

import (
	"testing"
	"time"
)

func TestBraveGo(t *testing.T) {
	BraveGo(func() {
		panic("this is a panic")
	}, func(p any) {
		t.Log(p)
	})
	time.Sleep(time.Second)
}

func TestBraveDo(t *testing.T) {
	BraveDo(func() {
		panic("this is a panic")
	}, func(p any) {
		t.Log(p)
	})
}
