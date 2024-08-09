package slicex

import (
	"github.com/go-leo/gox/syncx/chanx"
)

type Stream[T any] interface {
	// AllMatch returns whether all elements of this stream match the provided predicate.
	AllMatch(predicate func(value T) bool) bool

	//AnyMatch Returns whether any elements of this stream match the provided predicate.
	AnyMatch(predicate func(value T) bool) bool

	// AsSlice returns a slice containing the elements of this stream.
	AsSlice() []T

	// Count returns the count of elements in this stream.
	Count() int

	// Distinct	returns a stream consisting of the distinct elements (according to cmp function) of this stream.
	Distinct(cmp func(a, b T) bool) Stream[T]

	Range(func(index int, value T))
	//Filter()

}

type stream[T any] struct {
	ch <-chan T
}

func (s stream[T]) AllMatch(predicate func(value T) bool) bool {
	for value := range s.ch {
		if !predicate(value) {
			return false
		}
	}
	return true
}

func (s stream[T]) AnyMatch(predicate func(value T) bool) bool {
	for value := range s.ch {
		if predicate(value) {
			return true
		}
	}
	return false
}

func (s stream[T]) AsSlice() []T {
	r := make([]T, 0, s.Count())

	return r
}

func (s stream[T]) Count() int {
	return len(s.ch)
}

func (s stream[T]) Distinct(cmp func(a, b T) bool) Stream[T] {
	return nil
}

func (s stream[T]) Range(f func(index int, value T)) {
	go func() {
		var index int
		for value := range s.ch {
			f(index, value)
			index++
		}
	}()
}

func AsStream[T any](values ...T) Stream[T] {
	return stream[T]{ch: chanx.Emit(values...)}
}
