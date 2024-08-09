package antsx

import "github.com/panjf2000/ants/v2"

type Gopher struct {
	Pool *ants.MultiPool
}

func (g Gopher) Go(f func()) error {
	return g.Pool.Submit(f)
}
