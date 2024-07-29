package operator

// Ternary is akin to a ternary operator in Go,
// which based on the boolean `condition`, returns either `exprIfTrue` or `exprIfFalse`.
// It utilizes generics `T`, making it applicable for expressions of any type.
func Ternary[T any](condition bool, exprIfTrue, exprIfFalse T) T {
	if condition {
		return exprIfTrue
	}
	return exprIfFalse
}

// TernaryFunc is akin to a ternary operator in Go,
// which conditionally executes either `exprIfTrue` or `exprIfFalse` based on the value of `condition`,
// and returns the result of the executed function.
// It leverages Go's generics feature, making it applicable for any type `T`.
func TernaryFunc[T any](condition bool, exprIfTrue, exprIfFalse func() T) T {
	if condition {
		return exprIfTrue()
	}
	return exprIfFalse()
}
