package gpool

import (
	"context"
	"github.com/sherifabdlnaby/gpool"
)

type Gopher struct {
	Pool *gpool.Pool
}

func (g Gopher) Go(f func()) error {
	return g.Pool.Enqueue(context.Background(), f)
}
