package chanx

func Copy[T any](dest chan<- T, src <-chan T) {
	for v := range src {
		dest <- v
	}
	close(dest)
}

func AsyncCopy[T any](dest chan<- T, src <-chan T) {
	go func() { Copy[T](dest, src) }()
}

// Pipe copies values from src to dest.
// Deprecated: Do not use. use Copy instead.
func Pipe[T any](src <-chan T, dest chan<- T) {
	Copy[T](dest, src)
}

// AsyncPipe copies values from src to dest asynchronously.
// Deprecated: Do not use. use AsyncCopy instead.
func AsyncPipe[T any](src <-chan T, dest chan<- T) {
	AsyncCopy[T](dest, src)
}
