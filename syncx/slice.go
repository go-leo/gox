package syncx

import (
	"sync"
)

type Slice[S ~[]E, E any] struct {
	mu    sync.RWMutex
	slice S
}

func NewSlice[S ~[]E, E any]() *Slice[S, E] {
	return &Slice[S, E]{}
}

func WrapSlice[S ~[]E, E any](slice []E) *Slice[S, E] {
	return &Slice[S, E]{slice: slice}
}

func (s *Slice[S, E]) Range(f func(index int, elem E) bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for index, elem := range s.slice {
		if !f(index, elem) {
			break
		}
	}
}

func (s *Slice[S, E]) Append(elems ...E) *Slice[S, E] {
	s.mu.Lock()
	defer s.mu.Unlock()
	ns := new(Slice[S, E])
	ns.slice = append(s.slice, elems...)
	return ns
}

func (s *Slice[S, E]) Prepend(elems ...E) *Slice[S, E] {
	s.mu.Lock()
	defer s.mu.Unlock()
	ns := new(Slice[S, E])
	ns.slice = make([]E, len(elems)+len(s.slice))
	copy(ns.slice, elems)
	copy(ns.slice[len(elems):], s.slice)
	return ns
}

func (s *Slice[S, E]) Slice(low int, high int, max ...int) *Slice[S, E] {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ns := new(Slice[S, E])
	if len(max) == 0 {
		ns.slice = s.slice[low:high]
		return ns
	}
	if len(max) == 1 {
		ns.slice = s.slice[low:high:max[0]]
		return ns
	}
	panic("invalid argument")
}

func (s *Slice[S, E]) Index(x int) E {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.slice[x]
}

func (s *Slice[S, E]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.slice)
}

func (s *Slice[S, E]) Cap() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return cap(s.slice)
}

func (s *Slice[S, E]) Unwrap() []E {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.slice
}
