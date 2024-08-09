package gopgpoolx

import (
	"gopkg.in/go-playground/pool.v3"
)

type Gopher struct {
	Pool pool.Pool
}

func (g Gopher) Go(f func()) error {
	_ = g.Pool.Queue(func(wu pool.WorkUnit) (interface{}, error) {
		f()
		return nil, nil
	})
	return nil
}
