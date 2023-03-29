package sortx

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Asc sorts a slice of any ordered type in ascending order.
func Asc[E constraints.Ordered](x []E) {
	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})
}

// Desc sorts a slice of any ordered type in descending order.
func Desc[E constraints.Ordered](x []E) {
	sort.Slice(x, func(i, j int) bool {
		return x[i] > x[j]
	})
}

// IsAsc reports whether x is sorted in ascending order.
func IsAsc[E constraints.Ordered](x []E) bool {
	for i := len(x) - 1; i > 0; i-- {
		if x[i] < x[i-1] {
			return false
		}
	}
	return true
}

// IsDesc reports whether x is sorted in descending order.
func IsDesc[E constraints.Ordered](x []E) bool {
	for i := len(x) - 1; i > 0; i-- {
		if x[i] > x[i-1] {
			return false
		}
	}
	return true
}

func BubbleSort[E constraints.Ordered](x []E) {
	for i := 0; i < len(x)-1; i++ {
		for j := 1; j < len(x)-i; j++ {
			if x[j] < x[j-1] {
				x[j], x[j-1] = x[j-1], x[j]
			}
		}
	}
}

func SelectSort[E constraints.Ordered](x []E) {
	for i := 0; i < len(x); i++ {
		min := i
		for j := i + 1; j < len(x); j++ {
			if x[min] > x[j] {
				min = j
			}
		}
		x[i], x[min] = x[min], x[i]
	}
}

func InsertSort[E constraints.Ordered](x []E) {
	for i := 0; i < len(x); i++ {
		for j := i - 1; j >= 0; j-- {
			if x[j+1] < x[j] {
				x[j+1], x[j] = x[j], x[j+1]
			} else {
				break
			}
		}
	}
}
