package chanx

func Max[T any](in <-chan T, cmp func(a, b T) int) (T, bool) {
	var maxValue T
	maxValue, ok := <-in
	if !ok {
		return maxValue, false
	}
	for value := range in {
		if cmp(value, maxValue) > 0 {
			maxValue = value
		}
	}
	return maxValue, true
}

func Min[T any](in <-chan T, cmp func(a, b T) int) (T, bool) {
	var minValue T
	minValue, ok := <-in
	if !ok {
		return minValue, false
	}
	for value := range in {
		if cmp(value, minValue) < 0 {
			minValue = value
		}
	}
	return minValue, true
}
