package chanx

func Pipe[T any](src <-chan T, dest chan<- T) {
	for v := range src {
		dest <- v
	}
}

func AsyncPipe[T any](src <-chan T, dest chan<- T) {
	go func() { Pipe(src, dest) }()
}
