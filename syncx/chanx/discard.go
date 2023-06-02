package chanx

func Discard[T any](c <-chan T) {
	for range c {
	}
}
