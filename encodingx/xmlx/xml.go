package xmlx

import (
	"encoding/xml"
	"io"

	"github.com/go-leo/gox/encodingx"
)

func Marshal(v any) ([]byte, error) {
	return xml.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return xml.Unmarshal(data, v)
}

func NewEncoder(w io.Writer) encodingx.Encoder {
	return xml.NewEncoder(w)
}

func NewDecoder(r io.Reader) encodingx.Decoder {
	return xml.NewDecoder(r)
}
