package protobufx

import (
	"errors"
	"io"

	"google.golang.org/protobuf/proto"

	"github.com/go-leo/gox/encodingx"
)

var ErrConvertProtoMessage = errors.New("protobufx: failed convert to proto.Message")

func Marshal(v any) ([]byte, error) {
	m, ok := v.(proto.Message)
	if !ok {
		return nil, ErrConvertProtoMessage
	}
	return proto.Marshal(m)
}

func Unmarshal(data []byte, v any) error {
	m, ok := v.(proto.Message)
	if !ok {
		return ErrConvertProtoMessage
	}
	return proto.Unmarshal(data, m)
}

func NewEncoder(w io.Writer) encodingx.Encoder {
	return &encoder{w: w}
}

func NewDecoder(r io.Reader) encodingx.Decoder {
	return &decoder{r: r}
}

type encoder struct {
	MarshalOptions proto.MarshalOptions
	w              io.Writer
}

func (e *encoder) Encode(val any) error {
	m, ok := val.(proto.Message)
	if !ok {
		return ErrConvertProtoMessage
	}
	data, err := e.MarshalOptions.Marshal(m)
	if err != nil {
		return err
	}
	_, err = e.w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

type decoder struct {
	UnmarshalOptions proto.UnmarshalOptions
	r                io.Reader
}

func (d *decoder) Decode(obj any) error {
	m, ok := obj.(proto.Message)
	if !ok {
		return ErrConvertProtoMessage
	}
	data, err := io.ReadAll(d.r)
	if err != nil {
		return err
	}
	return d.UnmarshalOptions.Unmarshal(data, m)
}
