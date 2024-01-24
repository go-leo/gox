//go:build sonic_fastest

package jsonx

import (
	"github.com/bytedance/sonic"
	"github.com/go-leo/gox/encodingx"
	"io"
)

var json = sonic.ConfigFastest

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func MarshalString(v any) (string, error) {
	return json.MarshalToString(v)
}

func UnmarshalString(str string, v interface{}) error {
	return json.UnmarshalFromString(str, v)
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

func MarshalNoEscape(v any) ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
