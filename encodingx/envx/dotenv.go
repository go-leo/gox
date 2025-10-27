package envx

import (
	"bytes"
	"errors"
	"io"

	"github.com/go-leo/gonv"

	"github.com/joho/godotenv"
	"golang.org/x/exp/maps"

	"github.com/go-leo/gox/encodingx"
)

func Marshal(val any) ([]byte, error) {
	var buf bytes.Buffer
	err := NewEncoder(&buf).Encode(val)
	return buf.Bytes(), err
}

func Unmarshal(data []byte, val any) error {
	return NewDecoder(bytes.NewReader(data)).Decode(val)
}

func NewEncoder(w io.Writer) encodingx.Encoder {
	return &encoder{w: w}
}

func NewDecoder(r io.Reader) encodingx.Decoder {
	return &decoder{r: r}
}

type encoder struct {
	w io.Writer
}

func (e *encoder) Encode(val any) error {
	envMap, ok := val.(map[string]string)
	if !ok {
		envMapPtr, ok := val.(*map[string]string)
		if !ok {
			return errors.New("envx: value not convert to map[string]string")
		}
		envMap = *envMapPtr
	}
	data, err := godotenv.Marshal(envMap)
	if err != nil {
		return err
	}
	_, err = e.w.Write(gonv.StringToBytes(data))
	return err
}

type decoder struct {
	r io.Reader
}

func (d *decoder) Decode(val any) error {
	envMap, ok := val.(map[string]string)
	if !ok {
		envMapPtr, ok := val.(*map[string]string)
		if !ok {
			return errors.New("any not convert to map[string]string")
		}
		envMap = *envMapPtr
	}
	m, err := godotenv.Parse(d.r)
	if err != nil {
		return err
	}
	maps.Copy(envMap, m)
	return nil
}
