package listx

type List interface {
	Init() List
	Len() int
	Front() Element
	Back() Element
	Remove(e Element) any
	PushFront(v any) Element
	PushBack(v any) Element
	InsertBefore(v any, mark Element) Element
	InsertAfter(v any, mark Element) Element
	MoveToFront(e Element)
	MoveToBack(e Element)
	MoveBefore(e, mark Element)
	MoveAfter(e, mark Element)
	PushBackList(other List)
	PushFrontList(other List)
	Root() Element
}
