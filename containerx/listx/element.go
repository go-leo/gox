package listx

// Element is an element of a linked list.
type Element interface {
	Next() Element
	Prev() Element
	Value() any
}

type element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev Element

	// The list to which this element belongs.
	list List

	// The value stored with this element.
	value any
}

// Next returns the next list element or nil.
func (e *element) Next() Element {
	if p := e.next; e.list != nil && p != e.list.Root() {
		return p
	}
	return nil
}

// Prev returns the previous list element or nil.
func (e *element) Prev() Element {
	if p := e.prev; e.list != nil && p != e.list.Root() {
		return p
	}
	return nil
}

func (e *element) Value() any {
	return e.value
}
