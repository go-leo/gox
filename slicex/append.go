package slicex

import "golang.org/x/exp/slices"

func AppendFirst[S ~[]E, E any](s S, e E) S {
	return slices.Insert(s, 0, e)
}

func AppendIfNotContains[S ~[]E, E comparable](s S, v E) S {
	if slices.Contains(s, v) {
		return s
	}
	return append(s, v)
}

func Prepend[S ~[]E, E any](s S, elems ...E) S {
	return slices.Insert(s, 0, elems...)
}

// AppendUnique appends an element to a slice, if the element is not already in the slice
func AppendUnique[S ~[]E, E comparable](s S, v E) S {
	return AppendIfNotContains(s, v)
}
