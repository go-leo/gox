package slicex_test

import (
	"github.com/go-leo/gox/slicex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelete(t *testing.T) {
	t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, 0))
	t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, 1))
	t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, 2))
	t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, 3))
	t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, 4))
	t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, 5))
	t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, 6))
	// t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, -1))
	// t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, 7))
}

func TestDeleteAll(t *testing.T) {
	t.Log(slicex.DeleteAll([]int{0, 1, 2, 3, 4, 5, 6}, 0, 1))
	t.Log(slicex.DeleteAll([]int{0, 1, 2, 3, 4, 5, 6}, 1, 3))
	t.Log(slicex.DeleteAll([]int{0, 1, 2, 3, 4, 5, 6}, 2, 5))
	t.Log(slicex.DeleteAll([]int{0, 1, 2, 3, 4, 5, 6}, 3, 1))
	t.Log(slicex.DeleteAll([]int{0, 1, 2, 3, 4, 5, 6}, 4, 0))
	t.Log(slicex.DeleteAll([]int{0, 1, 2, 3, 4, 5, 6}, 5, 2))
	t.Log(slicex.DeleteAll([]int{0, 1, 2, 3, 4, 5, 6}, 6, 4))
	// t.Log(slicex.DeleteAll([]int{0, 1, 2, 3, 4, 5, 6}, -1))
	// t.Log(slicex.DeleteAll([]int{0, 1, 2, 3, 4, 5, 6}, 7))
}

func TestRemove(t *testing.T) {
	ints := slicex.Remove([]int{1, 2, 3, 4, 5}, 2)
	t.Log(ints)
}

func TestReverse(t *testing.T) {
	assert.Equal(t, []int{5, 4, 3, 2, 1}, slicex.Reverse([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, []string{"5", "4", "3", "2", "1"}, slicex.Reverse([]string{"1", "2", "3", "4", "5"}))
}

func TestToSet(t *testing.T) {
	strings := []string{"1", "2", "3"}
	set := slicex.ToSet[[]string, map[string]struct{}](strings)
	t.Log(set)
}

func TestConcat(t *testing.T) {
	t.Log(slicex.Concat([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 0}))
	t.Log(slicex.Concat([]float32{1, 2, 3, 4, 5}, []float32{6, 7, 8, 9, 0}))
	t.Log(slicex.Concat([]string{"1", "2", "3", "4", "5"}, []string{"6", "7", "8", "9", "0"}))
}

func TestChunk(t *testing.T) {
	tests := []struct {
		array          []int
		size           int
		expectedChunks [][]int
	}{
		{
			array:          []int{},
			size:           2,
			expectedChunks: [][]int{},
		},
		{
			array:          []int{0},
			size:           2,
			expectedChunks: [][]int{{0}},
		},
		{
			array:          []int{0, 1},
			size:           2,
			expectedChunks: [][]int{{0, 1}},
		},
		{
			array:          []int{0, 1, 2},
			size:           2,
			expectedChunks: [][]int{{0, 1}, {2}},
		},
		{
			array:          []int{0, 1, 2, 3},
			size:           2,
			expectedChunks: [][]int{{0, 1}, {2, 3}},
		},
		{
			array:          []int{0, 1, 2, 3, 4},
			size:           2,
			expectedChunks: [][]int{{0, 1}, {2, 3}, {4}},
		},
		{
			array:          []int{0, 1, 2, 3, 4, 5},
			size:           2,
			expectedChunks: [][]int{{0, 1}, {2, 3}, {4, 5}},
		},
	}
	for _, test := range tests {
		chunks := slicex.Chunk(test.array, test.size)
		if len(chunks) != len(test.expectedChunks) {
			t.Fatalf("%v expected chunks is %v, but is %v", test.array, test.expectedChunks, chunks)
		}
		for i := range chunks {
			if len(chunks[i]) != len(test.expectedChunks[i]) {
				t.Fatalf("%v expected chunks is %v, but is %v", test.array, test.expectedChunks, chunks)
			}
			for j := range chunks[i] {
				if chunks[i][j] != test.expectedChunks[i][j] {
					t.Fatalf("%v expected chunks is %v, but is %v", test.array, test.expectedChunks, chunks)
				}
			}
		}
		t.Log(chunks)
	}

}
