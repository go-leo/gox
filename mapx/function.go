package mapx

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

// FromRanger 从一个实现了 Range 方法的接口中创建一个 map
func FromRanger[M ~map[K]V, K comparable, V any](ranger interface {
	Range(func(key, value any) (shouldContinue bool))
}) M {
	m := make(M)
	ranger.Range(func(k any, v any) bool {
		m[k.(K)] = v.(V)
		return true
	})
	return m
}
