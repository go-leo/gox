package slicex

import "golang.org/x/exp/slices"

func FindFunc[E any](s []E, f func(E) bool) (E, bool) {
	if i := slices.IndexFunc(s, f); i != -1 {
		return s[i], true
	}
	var e E
	return e, false
}
