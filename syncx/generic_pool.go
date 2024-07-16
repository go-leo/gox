package syncx

import "sync"

type GenericPool[T any] struct {
	p sync.Pool
}

func NewGenericPool[T any](f func() T) *GenericPool[T] {
	if f == nil {
		panic("syncx: new function is nil")
	}
	return &GenericPool[T]{p: sync.Pool{New: func() any { return f() }}}
}

func (p *GenericPool[T]) Put(o T) {
	p.p.Put(o)
}

func (p *GenericPool[T]) Get() T {
	return p.p.Get().(T)
}
