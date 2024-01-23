package pool

type Pool interface {
	Go(f func()) error
}
