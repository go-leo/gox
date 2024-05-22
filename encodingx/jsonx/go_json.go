//go:build go_json

package jsonx

import (
	"github.com/go-leo/gox/encodingx"
	json "github.com/goccy/go-json"
	"io"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func NewEncoder(w io.Writer) encodingx.JSONEncoder {
	return json.NewEncoder(w)
}

func NewDecoder(r io.Reader) encodingx.Decoder {
	return json.NewDecoder(r)
}

func MarshalNoEscape(v any) ([]byte, error) {
	return json.MarshalNoEscape(v)
}
