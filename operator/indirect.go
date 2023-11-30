package operator

func Indirect[P *E, E any](p P) (E, bool) {
	var e E
	if p == nil {
		return e, false
	}
	return *p, true
}

func IndirectOrZero[P *E, E any](p P) (E, bool) {
	var e E
	if p == nil {
		return e, false
	}
	return *p, true
}
