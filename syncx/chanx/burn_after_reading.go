package chanx

func BurnAfterReading[M any](m M) <-chan M {
	ch := make(chan M, 1)
	ch <- m
	return ch
}
