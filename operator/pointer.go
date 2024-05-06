package operator

// Pointer stores v in a new E value and returns a pointer to it.
func Pointer[E any](v E) *E {
	return &v
}

// Indirect returns the value pointed to by p.
func Indirect[P *E, E any](p P) (E, bool) {
	if p == nil {
		var e E
		return e, false
	}
	return *p, true
}
