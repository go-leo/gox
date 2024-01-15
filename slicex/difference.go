package slicex

import "golang.org/x/exp/slices"

// Difference 返回差集
func Difference[S ~[]E, E comparable](s1 S, s2 S) S {
	if len(s1) >= len(s2) {
		return difference(s1, s2)
	}
	return difference(s2, s1)
}

func difference[S ~[]E, E comparable](a S, b S) S {
	var r S
	for _, v := range a {
		if !slices.Contains(b, v) {
			r = append(r, v)
		}
	}
	return r
}
