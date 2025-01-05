package chanx

func Once[M any](m M) <-chan M {
	ch := make(chan M, 1)
	ch <- m
	close(ch)
	return ch
}

func Nothing[M any]() <-chan M {
	ch := make(chan M)
	close(ch)
	return ch
}
