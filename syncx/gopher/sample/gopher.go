package sample

import "github.com/go-leo/gox/syncx/brave"

type Gopher struct {
	Recover func(p any)
}

func (g Gopher) Go(f func()) error {
	brave.Go(f, g.Recover)
	return nil
}
