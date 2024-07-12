package slicex

import (
	"golang.org/x/exp/slices"
	"sync"
)

type SyncSlice[S ~[]E, E any] struct {
	mu    sync.RWMutex
	slice S
}

func NewSyncSlice[S ~[]E, E any]() *SyncSlice[S, E] {
	return &SyncSlice[S, E]{}
}

func WrapSlice[S ~[]E, E any](slice []E) *SyncSlice[S, E] {
	return &SyncSlice[S, E]{slice: slice}
}

func (s *SyncSlice[S, E]) Range(f func(index int, elem E) bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for index, elem := range s.slice {
		if !f(index, elem) {
			break
		}
	}
}

func (s *SyncSlice[S, E]) Append(elems ...E) *SyncSlice[S, E] {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.slice = append(s.slice, elems...)
	return s
}

func (s *SyncSlice[S, E]) Prepend(elems ...E) *SyncSlice[S, E] {
	s.mu.Lock()
	defer s.mu.Unlock()
	slice := make([]E, len(elems)+len(s.slice))
	copy(slice, elems)
	copy(slice[len(elems):], s.slice)
	s.slice = slice
	return s
}

func (s *SyncSlice[S, E]) Slice(low int, high int, max ...int) *SyncSlice[S, E] {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ns := new(SyncSlice[S, E])
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

func (s *SyncSlice[S, E]) Index(x int) E {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.slice[x]
}

func (s *SyncSlice[S, E]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.slice)
}

func (s *SyncSlice[S, E]) Cap() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return cap(s.slice)
}

func (s *SyncSlice[S, E]) Unwrap() []E {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return slices.Clone(s.slice)
}
