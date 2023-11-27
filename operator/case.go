package operator

func SwitchCases[K comparable, V any](key K, keys []K, values []V) (V, bool) {
	var v V
	if len(keys) != len(values) {
		return v, false
	}
	for i := range keys {
		if keys[i] == key {
			return values[i], true
		}
	}
	return v, false
}
