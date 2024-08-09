package tunnyx

import "github.com/Jeffail/tunny"

type Gopher struct {
	Pool *tunny.Pool
}

func (g Gopher) Go(f func()) error {
	g.Pool.Process(f)
	return nil
}
