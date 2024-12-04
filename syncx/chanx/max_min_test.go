package chanx

import (
	"context"
	"testing"
)

func TestMax(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	in := make(chan int, 10)
	in <- 1
	in <- 3
	in <- 5
	in <- 2
	in <- 8
	in <- 7
	close(in)

	out := Max(ctx, in, func(a, b int) int {
		return a - b
	})

	select {
	case <-ctx.Done():
		t.Errorf("context done")
	case max := <-out:
		if max != 8 {
			t.Errorf("expected max value to be 8, got %d", max)
		}
	}
}

func TestMin(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	in := make(chan int, 10)
	in <- 1
	in <- 3
	in <- 5
	in <- 2
	in <- 8
	in <- 7
	close(in)

	out := Min(ctx, in, func(a, b int) int {
		return a - b
	})

	select {
	case <-ctx.Done():
		t.Errorf("context done")
	case min := <-out:
		if min != 1 {
			t.Errorf("expected min value to be 1, got %d", min)
		}
	}
}
