package chanx

// AllMatch 检查通道中所有元素是否满足给定条件，若全部满足则返回true，否则返回false。
func AllMatch[T any](in <-chan T, predicate func(value T) bool) bool {
	for value := range in {
		if !predicate(value) {
			return false
		}
	}
	return true
}

// AnyMatch 检查通道中是否有元素满足给定条件，找到即返回 true，否则返回 false。
func AnyMatch[T any](in <-chan T, predicate func(value T) bool) bool {
	for value := range in {
		if predicate(value) {
			return true
		}
	}
	return false
}

// NoneMatch 检查通道中的所有元素是否都不满足给定条件，若所有元素都不满足则返回true，否则返回false。
func NoneMatch[T any](in <-chan T, predicate func(value T) bool) bool {
	for value := range in {
		if predicate(value) {
			return false
		}
	}
	return true
}
