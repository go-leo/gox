package sample

type Gopher struct {
	Pool interface{ Submit(task func()) error }
}

func (g Gopher) Go(f func()) error {
	return g.Pool.Submit(f)
}
