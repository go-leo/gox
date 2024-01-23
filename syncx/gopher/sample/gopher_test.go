package sample

import (
	"testing"
)

func TestGopher(t *testing.T) {
	gopher := Gopher{}
	_ = gopher.Go(func() {
		t.Log("hello")
	})
}

func BenchmarkGopher(b *testing.B) {
	gopher := Gopher{}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			gopher.Go(func() {})
		}
	})
}

func BenchmarkGo(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			go func() {}()
		}
	})
}
