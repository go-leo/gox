package slicex

import "golang.org/x/exp/slices"

func appendFirst[S ~[]E, E any](s S, e E) S {
	if len(s) == 0 {
		return append(s, e)
	}
	return Insert(s, e, 0)
}

func AppendIfNotContains[S ~[]E, E comparable](s S, v E) S {
	if slices.Contains(s, v) {
		return s
	}
	return append(s, v)
}
