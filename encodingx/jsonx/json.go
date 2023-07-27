//go:build !jsoniter && !go_json && !jsoniter_fastest && !sonic

package jsonx

import (
	"encoding/json"
	"github.com/go-leo/gox/encodingx"
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

func NewEncoder(w io.Writer) JSONEncoder {
	return json.NewEncoder(w)
}

func NewDecoder(r io.Reader) encodingx.Decoder {
	return json.NewDecoder(r)
}
