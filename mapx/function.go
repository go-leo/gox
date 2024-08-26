package mapx

import (
	"golang.org/x/sync/singleflight"
)

func AppendValue[M ~map[K]V, R ~map[K]S, S ~[]V, K comparable, V any](r R, ms ...M) R {
	if r == nil {
		r = make(R)
	}
	for _, elem := range ms {
		for k, v := range elem {
			r[k] = append(r[k], v)
		}
	}
	return r
}

func IsEmpty[M ~map[K]V, K comparable, V any](m M) bool {
	return len(m) <= 0
}

func IsNotEmpty[M ~map[K]V, K comparable, V any](m M) bool {
	return len(m) > 0
}

func Entries[M ~map[K]V, K comparable, V any](m M) []Entry[K, V] {
	r := make([]Entry[K, V], 0, len(m))
	for k, v := range m {
		r = append(r, Entry[K, V]{Key: k, Value: v})
	}
	return r
}

func KeySet[M ~map[K]V, R map[K]struct{}, K comparable, V any](m M) R {
	r := make(R, len(m))
	for k := range m {
		r[k] = struct{}{}
	}
	return r
}

func FromMapRange[M ~map[K]V, K comparable, V any](mi MapInterface) M {
	m := make(M)
	mi.Range(func(k any, v any) bool {
		m[k.(K)] = v.(V)
		return true
	})
	return m
}

// LoadOrCreate 尝试从给定的Map中加载键key的值，若不存在，则仅执行一次计算函数f并将结果存入Map，最后返回值及是否已存在的标志。
func LoadOrCreate(m MapInterface, sfg *singleflight.Group, key string, f func() (any, error)) (any, error, bool) {
	if value, ok := m.Load(key); ok {
		return value, nil, true
	}
	value, err, _ := sfg.Do(key, func() (interface{}, error) {
		if value, ok := m.Load(key); ok {
			return value, nil
		}
		value, err := f()
		if err != nil {
			return nil, err
		}
		m.Store(key, value)
		return value, nil
	})
	return value, err, false
}
