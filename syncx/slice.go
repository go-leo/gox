package syncx

import "sync"

type Slice[E any] struct {
	mu    sync.RWMutex
	slice []E
}

func WrapSlice[E any](slice []E) *Slice[E] {
	return &Slice[E]{slice: slice}
}

func (s *Slice[E]) Range(f func(index int, elem E) bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for index, elem := range s.slice {
		if !f(index, elem) {
			break
		}
	}
}

func (s *Slice[E]) Append(elems ...E) *Slice[E] {
	s.mu.Lock()
	defer s.mu.Unlock()
	ns := new(Slice[E])
	ns.slice = append(s.slice, elems...)
	return ns
}

func (s *Slice[E]) Slice(low int, high int, max ...int) *Slice[E] {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ns := new(Slice[E])
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

func (s *Slice[E]) Index(x int) E {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.slice[x]
}

func (s *Slice[E]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.slice)
}

func (s *Slice[E]) Cap() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return cap(s.slice)
}

func (s *Slice[E]) Unwrap() []E {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.slice
}
