package listx

import "errors"

var ErrEmptyQueue = errors.New("queue is empty")

var ErrFullQueue = errors.New("queue is full")

// Queue is a FIFO list.
type Queue interface {

	// Enqueue inserts the specified element into this queue.
	// If this queue is full, returns ErrFullQueue error.
	Enqueue(v any) error

	// Dequeue retrieves and removes the head of this queue.
	// If this queue is empty. returns ErrEmptyQueue error.
	Dequeue() (any, error)

	// Peek retrieves, but does not remove, the head of this queue.
	// If this queue is empty, returns ErrEmptyQueue error.
	Peek() (any, error)

	// Empty tests if this queue is empty.
	Empty()

	// Search returns the 1-based position where an object is on this queue.
	Search(v any) int
}
