package mapx

import "github.com/go-leo/gox/containerx"

func FromMapRange[M ~map[K]V, K comparable, V any](mi containerx.MapInterface) M {
	m := make(M)
	mi.Range(func(k any, v any) bool {
		m[k.(K)] = v.(V)
		return true
	})
	return m
}