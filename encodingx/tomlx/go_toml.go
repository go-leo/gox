//go:build go_toml

package tomlx

import (
	"io"

	"github.com/pelletier/go-toml/v2"

	"github.com/go-leo/gox/encodingx"
)

func Marshal(v any) ([]byte, error) {
	return toml.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return toml.Unmarshal(data, v)
}

func NewEncoder(w io.Writer) encodingx.Encoder {
	return toml.NewEncoder(w)
}

func NewDecoder(r io.Reader) encodingx.Decoder {
	return toml.NewDecoder(r)
}
