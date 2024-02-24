package sample

// import "github.com/panjf2000/ants/v2"

type Gopher struct {
	//Pool *ants.Pool
	//Pool ants.MultiPool
	Pool interface{ Submit(task func()) error }
}

func (g Gopher) Go(f func()) error {
	return g.Pool.Submit(f)
}
