package slicex

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func Sort[S ~[]E, E constraints.Ordered](s S) {
	slices.Sort([]E(s))
}

func SortFunc[S ~[]E, E any](s S, less func(a, b E) bool) {
	slices.SortFunc([]E(s), less)
}

func SortStableFunc[S ~[]E, E any](s S, less func(a, b E) bool) {
	slices.SortStableFunc([]E(s), less)
}

func IsSorted[S ~[]E, E constraints.Ordered](s S) bool {
	return slices.IsSorted([]E(s))
}

func IsSortedFunc[S ~[]E, E any](s S, less func(a, b E) bool) bool {
	return slices.IsSortedFunc([]E(s), less)
}
