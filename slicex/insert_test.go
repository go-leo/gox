package slicex_test

import (
	"testing"

	"github.com/go-leo/gox/slicex"
)

func TestInsert(t *testing.T) {
	t.Log(slicex.Insert(0, []int{1, 2, 3, 4, 5}, 0))
	t.Log(slicex.Insert(0, []int{1, 2, 3, 4, 5}, 0, 10, 20))
	t.Log(slicex.Insert(1, []int{1, 2, 3, 4, 5}, 11, 111, 1111))
	t.Log(slicex.Insert(2, []int{1, 2, 3, 4, 5}, 22, 222, 2222))
	t.Log(slicex.Insert(3, []int{1, 2, 3, 4, 5}, 33, 333, 3333))
	t.Log(slicex.Insert(4, []int{1, 2, 3, 4, 5}, 44, 444, 4444))
	t.Log(slicex.Insert(5, []int{1, 2, 3, 4, 5}, 55, 555, 5555))

	// t.Log(slicex.Insert(-1, []int{1, 2, 3, 4, 5}, 0))
	// t.Log(slicex.Insert(6, []int{1, 2, 3, 4, 5}, 66, 666, 6666))
}
