package slicex

import "testing"

func TestName(t *testing.T) {
	strings := []string{"1", "2", "3"}
	set := ToSet[[]string, map[string]struct{}](strings)
	t.Log(set)
}
