package slicex

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func Insert[S ~[]E, E any](index int, s S, es ...E) S {
	if len(es) <= 0 {
		return slices.Clone(s)
	}
	if index < 0 || index > len(s) {
		panic(fmt.Errorf("index: %d, length: %d", index, len(s)))
	}
	r := make(S, len(s)+len(es))
	if index > 0 {
		copy(r[:index], s[:index])
	}
	copy(r[index:index+len(es)], es)
	if index < len(s) {
		copy(r[index+len(es):], s[index:])
	}
	return r
}
