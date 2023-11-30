package operator

func Indirect[P *E, E any](p P) (E, bool) {
	if p == nil {
		var e E
		return e, false
	}
	return *p, true
}

func IndirectOrZero[P *E, E any](p P) E {
	if p == nil {
		var e E
		return e
	}
	return *p
}
