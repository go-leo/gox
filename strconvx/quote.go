package strconvx

import "strconv"

// Quote quotes the string.
func Quote[E ~string](e E) E {
	return E(strconv.Quote(string(e)))
}

// QuoteSlice quotes each string in the slice.
func QuoteSlice[S ~[]E, E ~string](s S) S {
	if s == nil {
		return s
	}
	r := make(S, 0, len(s))
	for _, e := range s {
		r = append(r, Quote(e))
	}
	return r
}
