//go:build toml
// +build toml

package tomlx

import (
	"bytes"
	"io"

	"github.com/BurntSushi/toml"

	"github.com/go-leo/gox/encodingx"
)

func Marshal(v any) ([]byte, error) {
	w := &bytes.Buffer{}
	err := toml.NewEncoder(w).Encode(v)
	return w.Bytes(), err
}

func Unmarshal(data []byte, v any) error {
	return toml.Unmarshal(data, v)
}

func NewEncoder(w io.Writer) encodingx.Encoder {
	return toml.NewEncoder(w)
}

func NewDecoder(r io.Reader) encodingx.Decoder {
	return &decoder{Decoder: toml.NewDecoder(r)}
}

type decoder struct {
	Decoder *toml.Decoder
}

func (d *decoder) Decode(obj interface{}) error {
	_, err := d.Decoder.Decode(obj)
	return err
}
