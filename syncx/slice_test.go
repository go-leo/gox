package syncx

import "testing"

func TestSlice(t *testing.T) {
	s := NewSlice[int]()
	t.Log("len:", s.Len())
	t.Log("cap:", s.Cap())
	s = s.Append(1, 2, 3, 4, 5)
	t.Log("len:", s.Len())
	t.Log("cap:", s.Cap())
	t.Log(s.Unwrap())
	s = s.Prepend(0, 9, 8, 7, 6)
	t.Log("len:", s.Len())
	t.Log("cap:", s.Cap())
	t.Log(s.Unwrap())
	s.Range(func(index int, elem int) bool {
		t.Log(index, elem)
		return true
	})
	t.Log(s.Index(1))

	s1 := s.Slice(1, 7)
	t.Log("s1 len:", s1.Len())
	t.Log("s1 cap:", s1.Cap())
	t.Log(s1.Unwrap())

	s2 := s.Slice(1, 7, 9)
	t.Log("s2 len:", s2.Len())
	t.Log("s2 cap:", s2.Cap())
	t.Log(s2.Unwrap())
}
