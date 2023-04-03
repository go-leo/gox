package envx

import (
	"errors"
	"io"

	"github.com/joho/godotenv"
	"golang.org/x/exp/maps"

	"github.com/go-leo/gox/bytesconvx"
	"github.com/go-leo/gox/encodingx"
)

func Marshal(envMap map[string]string) ([]byte, error) {
	data, err := godotenv.Marshal(envMap)
	if err != nil {
		return nil, err
	}
	return bytesconvx.StringToBytes(data), nil
}

func Unmarshal(data []byte, envMap map[string]string) error {
	m, err := godotenv.UnmarshalBytes(data)
	if err != nil {
		return err
	}
	maps.Copy(envMap, m)
	return nil
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
	m, ok := val.(map[string]string)
	if !ok {
		return errors.New("any not convert to map[string]string")
	}
	data, err := godotenv.Marshal(m)
	if err != nil {
		return err
	}
	_, err = e.w.Write(bytesconvx.StringToBytes(data))
	return err
}

type decoder struct {
	r io.Reader
}

func (d *decoder) Decode(obj any) error {
	m, ok := obj.(map[string]string)
	if !ok {
		return errors.New("any not convert to map[string]string")
	}
	envMap, err := godotenv.Parse(d.r)
	if err != nil {
		return err
	}
	maps.Copy(m, envMap)
	return nil
}
