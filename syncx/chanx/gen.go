package chanx

// Gen creates and returns a read-only channel that sequentially sends all the passed `values` through the channel.
// Once all `values` have been sent, it closes the channel.
// The function uses generics `[T any]`, making it applicable to values of any type.
// Internally, it starts a goroutine where it iterates over the `values`, sending each one through the channel.
// Finally, the function returns this channel.
//
// See: [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
func Gen[T any](values ...T) <-chan T {
	out := make(chan T, len(values))
	for _, value := range values {
		out <- value
	}
	close(out)
	return out
}
