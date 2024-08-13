package slicex

import (
	"fmt"
)

type Stream[T any] interface {
	// AllMatch returns whether all elements of this stream match the provided predicate.
	AllMatch(predicate func(value T) bool) bool

	//AnyMatch Returns whether any elements of this stream match the provided predicate.
	AnyMatch(predicate func(value T) bool) bool

	// NoneMatch returns whether no elements of this stream match the provided predicate.
	NoneMatch(predicate func(value T) bool) bool

	// AsSlice returns a slice containing the elements of this stream.
	AsSlice() []T

	// Count returns the count of elements in this stream.
	Count() int

	// Distinct	returns a stream consisting of the distinct elements (according to cmp function) of this stream.
	Distinct(cmp func(a, b T) bool) Stream[T]

	// Filter returns a stream consisting of the elements of this stream that match the given predicate.
	Filter(predicate func(value T) bool) Stream[T]

	// FindFirst returns the first element of this stream, or false if the stream is empty.
	FindFirst() (T, bool)

	// Range performs an action for each element of this stream.
	Range(func(index int, value T))

	// Limit returns a stream consisting of the elements of this stream, truncated to be no longer than maxSize in length.
	Limit(maxSize int) Stream[T]

	// Max returns the maximum element of this stream according to the provided Comparator.
	Max(cmp func(a, b T) int) (T, bool)

	// Min returns the minimum element of this stream according to the provided Comparator.
	Min(cmp func(a, b T) int) (T, bool)

	// Peek	returns a stream consisting of the elements of this stream, additionally performing the provided action on
	// each element as elements are consumed from the resulting stream.
	Peek(action func(T)) Stream[T]
}

type stream[T any] struct {
	ch            chan T
	interrupt     chan struct{}
	inputFinished chan struct{}
}

func (s *stream[T]) AllMatch(predicate func(value T) bool) bool {
	for value := range s.ch {
		if !predicate(value) {
			return false
		}
	}
	return true
}

func (s *stream[T]) AnyMatch(predicate func(value T) bool) bool {
	for value := range s.ch {
		if predicate(value) {
			return true
		}
	}
	return false
}

func (s *stream[T]) NoneMatch(predicate func(value T) bool) bool {
	for value := range s.ch {
		if predicate(value) {
			return false
		}
	}
	return true
}

func (s *stream[T]) AsSlice() []T {
	var r []T
	for value := range s.ch {
		r = append(r, value)
	}
	return r
}

func (s *stream[T]) Count() int {
	<-s.inputFinished
	return len(s.ch)
}

func (s *stream[T]) Distinct(cmp func(value, exists T) bool) Stream[T] {
	r := &stream[T]{
		ch:            make(chan T, cap(s.ch)),
		interrupt:     make(chan struct{}),
		inputFinished: make(chan struct{}),
	}
	go func() {
		defer close(r.inputFinished)
		defer close(r.ch)
		values := make([]T, 0, cap(s.ch))
	outer:
		for value := range s.ch {
			// 判断是否已经发送过
			for _, exists := range values {
				if cmp(value, exists) {
					continue outer
				}
			}
			select {
			case r.ch <- value:
				// 发送value
				// 记录已发送的value
				values = append(values, value)
			case <-r.interrupt:
				// 被下游打断
				// 打断上游stream
				close(s.interrupt)
				return
			}
		}
	}()
	return r
}

func (s *stream[T]) Filter(predicate func(value T) bool) Stream[T] {
	r := &stream[T]{
		ch:            make(chan T, cap(s.ch)),
		interrupt:     make(chan struct{}),
		inputFinished: make(chan struct{}),
	}
	go func() {
		defer close(r.inputFinished)
		defer close(r.ch)
		for value := range s.ch {
			if !predicate(value) {
				continue
			}
			select {
			case r.ch <- value:
				// 发送value
			case <-r.interrupt:
				// 被下游打断
				// 打断上游stream
				close(s.interrupt)
				return
			}
		}
	}()
	return r
}

func (s *stream[T]) FindFirst() (T, bool) {
	value, ok := <-s.ch
	return value, ok
}

func (s *stream[T]) Range(f func(index int, value T)) {
	var index int
	for value := range s.ch {
		f(index, value)
		index++
	}
}

func (s *stream[T]) Limit(maxSize int) Stream[T] {
	r := &stream[T]{
		ch:            make(chan T, maxSize),
		interrupt:     make(chan struct{}),
		inputFinished: make(chan struct{}),
	}
	go func() {
		defer close(r.inputFinished)
		defer close(r.ch)
		for i := 0; i < maxSize; i++ {
			select {
			case value, ok := <-s.ch:
				if !ok {
					// 上游stream已关闭， 返回
					return
				}
				// 发送value
				r.ch <- value
			case <-r.interrupt:
				// 被下游打断
				// 打断上游stream
				close(s.interrupt)
				return
			}
		}
		// 达到最大次数限制
		// 打断上游stream
		close(s.interrupt)
	}()
	return r
}

func (s *stream[T]) Max(cmp func(a, b T) int) (T, bool) {
	var maxValue T
	maxValue, ok := <-s.ch
	if !ok {
		return maxValue, false
	}
	for value := range s.ch {
		if cmp(value, maxValue) > 0 {
			maxValue = value
		}
	}
	return maxValue, true
}

func (s *stream[T]) Min(cmp func(a, b T) int) (T, bool) {
	var minValue T
	minValue, ok := <-s.ch
	if !ok {
		return minValue, false
	}
	for value := range s.ch {
		if cmp(value, minValue) < 0 {
			minValue = value
		}
	}
	return minValue, true
}

func (s *stream[T]) Peek(action func(T)) Stream[T] {
	r := &stream[T]{
		ch:            make(chan T, cap(s.ch)),
		interrupt:     make(chan struct{}),
		inputFinished: make(chan struct{}),
	}
	go func() {
		defer close(r.inputFinished)
		defer close(r.ch)
		for value := range s.ch {
			action(value)
			select {
			case r.ch <- value:
				// 发送value
			case <-r.interrupt:
				// 被下游打断
				// 打断上游stream
				close(s.interrupt)
				return
			}
		}
	}()
	return r
}

func AsStream[T any](values ...T) Stream[T] {
	r := &stream[T]{
		ch:            make(chan T, len(values)),
		interrupt:     make(chan struct{}),
		inputFinished: make(chan struct{}),
	}
	go func() {
		defer close(r.inputFinished)
		defer close(r.ch)
		for _, value := range values {
			select {
			case r.ch <- value:
				// 发送value
			case <-r.interrupt:
				// 被下游打断
				return
			}
		}
	}()
	return r
}

// GenerateStream returns an infinite sequential unordered stream where each element is generated by the provided supplier.
func GenerateStream[T any](supplier func() T) Stream[T] {
	r := &stream[T]{
		ch:            make(chan T, 1),
		interrupt:     make(chan struct{}),
		inputFinished: make(chan struct{}),
	}
	go func() {
		defer close(r.inputFinished)
		defer close(r.ch)
		for {
			value := supplier()
			select {
			case r.ch <- value:
				// 发送value
			case <-r.interrupt:
				// 被下游打断
				return
			}
		}
	}()
	return r
}

// IterateStream returns an infinite sequential ordered Stream produced by iterative application of a function f to an
// initial element seed, producing a Stream consisting of seed, f(seed), f(f(seed)), etc.
func IterateStream[T any](seed T, f func(T) T) Stream[T] {
	r := &stream[T]{
		ch:            make(chan T, 1),
		interrupt:     make(chan struct{}),
		inputFinished: make(chan struct{}),
	}
	go func() {
		defer close(r.inputFinished)
		defer close(r.ch)
		pre := seed
		for {
			value := f(pre)
			select {
			case r.ch <- value:
				// 发送value
			case <-r.interrupt:
				// 被下游打断
				return
			}
			pre = value
		}
	}()
	return r
}

// StreamMap returns a stream consisting of the results of applying the given function to the elements of this stream.
func StreamMap[T any, R any](src Stream[T], mapper func(T) R) Stream[R] {
	s, ok := src.(*stream[T])
	if !ok {
		panic(fmt.Errorf("slicex: failed to convert %T to %T", src, new(stream[T])))
	}
	r := &stream[R]{
		ch:            make(chan R, cap(s.ch)),
		interrupt:     make(chan struct{}),
		inputFinished: make(chan struct{}),
	}
	go func() {
		defer close(r.inputFinished)
		defer close(r.ch)
		for value := range s.ch {
			mappedValue := mapper(value)
			select {
			case r.ch <- mappedValue:
				// 发送value
			case <-r.interrupt:
				// 被下游打断
				// 打断上游stream
				close(s.interrupt)
				return
			}
		}
	}()
	return r
}

// StreamReduce Performs a reduction on the elements of this stream, using an associative accumulation function, and
// returns the reduced value.
func StreamReduce[T any, R any](src Stream[T], identity R, accumulator func(R, T) R, combiner func(R, R) R) R {
	s, ok := src.(*stream[T])
	if !ok {
		panic(fmt.Errorf("slicex: failed to convert %T to %T", src, new(stream[T])))
	}

}

// StreamFlatMap returns a stream consisting of the results of replacing each element of this stream with the contents
// of a mapped stream produced by applying the provided mapping function to each element.
func StreamFlatMap[T any, R Stream[E], E any](src Stream[T], mapper func(T) R) Stream[R] {
	s, ok := src.(*stream[T])
	var length int
	if !ok {
		panic(fmt.Errorf("slicex: failed to convert %T to %T", src, new(stream[T])))
	}
	length = cap(s.ch)
	r := &stream[R]{
		ch:            make(chan R, length),
		interrupt:     make(chan struct{}),
		inputFinished: make(chan struct{}),
	}
	go func() {
		defer close(r.inputFinished)
		defer close(r.ch)
		for value := range s.ch {
			mappedValue := mapper(value)
			select {
			case r.ch <- mappedValue:
				// 发送value
			case <-r.interrupt:
				// 被下游打断
				// 打断上游stream
				close(s.interrupt)
				return
			}
		}
	}()
	return r
}
