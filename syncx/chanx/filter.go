package chanx

func Filter[T any](in <-chan T, predicate func(value T) bool) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for value := range in {
			if !predicate(value) {
				continue
			}
			out <- value
		}
	}()
	return out
}
