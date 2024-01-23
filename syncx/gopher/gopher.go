package gopher

type Gopher interface {
	Go(f func()) error
}
