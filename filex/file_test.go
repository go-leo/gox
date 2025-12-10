package filex

import (
	"testing"
)

func TestHumanReadableSize(t *testing.T) {
	t.Log(HumanReadableSize(1))
	t.Log(HumanReadableSize(10))
	t.Log(HumanReadableSize(100))
	t.Log(HumanReadableSize(1000))
	t.Log(HumanReadableSize(10000))
	t.Log(HumanReadableSize(100000))
	t.Log(HumanReadableSize(1000000))
	t.Log(HumanReadableSize(10000000))
	t.Log(HumanReadableSize(100000000))
	t.Log(HumanReadableSize(1000000000))
	t.Log(HumanReadableSize(10000000000))
	t.Log(HumanReadableSize(100000000000))
	t.Log(HumanReadableSize(1000000000000))
	t.Log(HumanReadableSize(10000000000000))
	t.Log(HumanReadableSize(100000000000000))
}
