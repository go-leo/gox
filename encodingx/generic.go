package encodingx

func GenericUnmarshal[V any](data []byte, unmarshal func(data []byte, v any) error) (V, error) {
	var v V
	err := unmarshal(data, &v)
	if err != nil {
		return v, err
	}
	return v, nil
}
