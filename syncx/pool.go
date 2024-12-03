package syncx

import "sync"

type Pool[T any] struct {
	p sync.Pool
}

func NewPool[T any](f func() T) *Pool[T] {
	if f == nil {
		panic("syncx: new function is nil")
	}
	return &Pool[T]{p: sync.Pool{New: func() any { return f() }}}
}

func (p *Pool[T]) Put(o T) {
	p.p.Put(o)
}

func (p *Pool[T]) Get() T {
	return p.p.Get().(T)
}
