package chanx

func SendChannels[T any](channels ...chan T) []<-chan T {
	c := make([]<-chan T, len(channels))
	for _, ch := range channels {
		c = append(c, ch)
	}
	return c
}

func ReceiveChannels[T any](channels ...chan T) []chan<- T {
	c := make([]chan<- T, len(channels))
	for _, ch := range channels {
		c = append(c, ch)
	}
	return c
}
