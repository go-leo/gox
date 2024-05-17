package errorx

// Break 函数是一个高阶函数，用于处理错误并决定是否继续执行另一个函数。
// 如果pre参数不为nil，它会返回一个零值和pre错误，从而中断后续的执行；
// 如果pre为nil，则会调用并返回f()的结果。
func Break[T any](pre error) func(f func() (T, error)) (T, error) {
	return func(f func() (T, error)) (T, error) {
		if pre != nil {
			var v T       // Declare a zero value of type T
			return v, pre // If pre is not nil, return the zero value of T and the error
		}
		return f() // If pre is nil, execute the wrapped function and return its result
	}
}

// Continue 函数允许在执行f()之前检查pre错误，如果f()返回错误，它会将这两个错误合并并返回
func Continue[T any](pre error) func(f func() (T, error)) (T, error) {
	return func(f func() (T, error)) (T, error) {
		v, err := f()
		if err != nil {
			var v T                  // Declare a zero value of type T
			return v, Join(pre, err) // If there's an error from f, combine the previous error pre with the new error using Join, and return the zero value of T

		}
		return v, pre // If f executes without error, return the value and the original error pre	}
	}
}
