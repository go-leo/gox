package listx

import "errors"

var ErrEmptyStack = errors.New("stack is empty")

var ErrFullStack = errors.New("stack is full")

// Stack is a LIFO list.
type Stack interface {

	// Pop removes the object at the top of this stack and returns that object as the value of this function.
	// If stack is empty, returns ErrEmptyStack error.
	Pop() (any, error)

	// Push pushes an item onto the top of this stack.
	// If stack is full, returns ErrFullStack error.
	Push(v any) error

	// Peek looks at the object at the top of this stack without removing it from the stack.
	// If stack is empty, returns ErrEmptyStack error.
	Peek() (any, error)

	// Empty tests if this stack is empty.
	Empty()

	// Search returns the 1-based position where an object is on this stack.
	Search(v any) int
}
