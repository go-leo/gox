package yamlx

import (
	"io"

	"gopkg.in/yaml.v3"

	"github.com/go-leo/gox/encodingx"
)

func Marshal(v any) ([]byte, error) {
	return yaml.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return yaml.Unmarshal(data, v)
}

func NewEncoder(w io.Writer) encodingx.Encoder {
	return yaml.NewEncoder(w)
}

func NewDecoder(r io.Reader) encodingx.Decoder {
	return yaml.NewDecoder(r)
}
