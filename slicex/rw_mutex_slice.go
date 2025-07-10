package slicex

import (
	"golang.org/x/exp/slices"
	"sync"
)

// RWMutexSlice 是一个线程安全的切片。
type RWMutexSlice[S ~[]E, E any] struct {
	// mu 是一个读写锁，用于保护切片的并发访问。
	mu sync.RWMutex
	// raw 是一个切片，存储切片中的元素。
	raw S
}

func MakeSlice[S ~[]E, E any](length, capacity int) *RWMutexSlice[S, E] {
	return &RWMutexSlice[S, E]{raw: make(S, length, capacity)}
}

func WrapSlice[S ~[]E, E any](raw []E) *RWMutexSlice[S, E] {
	return &RWMutexSlice[S, E]{raw: raw}
}

// Range 遍历切片中的每个元素，并调用 f 函数。如果 f 返回 false，则停止遍历。
func (s *RWMutexSlice[S, E]) Range(f func(index int, elem E) bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for index, elem := range s.raw {
		if !f(index, elem) {
			break
		}
	}
}

// Append 将一个或多个元素追加到切片的末尾。
func (s *RWMutexSlice[S, E]) Append(elems ...E) *RWMutexSlice[S, E] {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.raw = append(s.raw, elems...)
	return s
}

// Prepend 将一个或多个元素插入到切片的开头。
func (s *RWMutexSlice[S, E]) Prepend(elems ...E) *RWMutexSlice[S, E] {
	s.mu.Lock()
	defer s.mu.Unlock()
	slice := make([]E, len(elems)+len(s.raw))
	copy(slice, elems)
	copy(slice[len(elems):], s.raw)
	s.raw = slice
	return s
}

// Slice 返回一个新的切片，包含原切片中从 low 到 high 的元素。如果提供了 max，则返回的切片的容量为 max。
func (s *RWMutexSlice[S, E]) Slice(low int, high int, max ...int) *RWMutexSlice[S, E] {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if low < 0 || high > len(s.raw) || low > high {
		panic("slicex: raw bounds out of range")
	}
	if len(max) == 0 {
		return WrapSlice[S, E](slices.Clone(s.raw[low:high]))
	}
	if len(max) == 1 {
		if max[0] < high || max[0] > cap(s.raw) {
			panic("slicex: raw capacity out of range")
		}
		return WrapSlice[S, E](slices.Clone(s.raw[low:high:max[0]]))
	}
	panic("slicex: invalid argument")
}

// Index 返回切片中指定索引位置的元素。
func (s *RWMutexSlice[S, E]) Index(x int) E {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if x < 0 || x >= len(s.raw) {
		panic("slicex: index out of range")
	}
	return s.raw[x]
}

// Len 返回切片的长度。
func (s *RWMutexSlice[S, E]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.raw)
}

// Cap 返回切片的容量。
func (s *RWMutexSlice[S, E]) Cap() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return cap(s.raw)
}

// Unwrap 返回一个新的切片，包含原切片中的所有元素。
func (s *RWMutexSlice[S, E]) Unwrap() []E {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return slices.Clone(s.raw)
}

// ClearAndUnwrap 清空切片并返回原切片。
func (s *RWMutexSlice[S, E]) ClearAndUnwrap() []E {
	s.mu.Lock()
	defer s.mu.Unlock()
	raw := slices.Clone(s.raw)
	s.raw = s.raw[:]
	return raw
}
