//go:build protojson

package jsonx

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-leo/gox/encodingx"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
)

var ErrConvertProtoMessage = errors.New("jsonx: failed convert to proto.Message")

func Marshal(v any) ([]byte, error) {
	m, ok := v.(proto.Message)
	if !ok {
		return nil, ErrConvertProtoMessage
	}
	return protojson.Marshal(m)
}

func Unmarshal(data []byte, v any) error {
	m, ok := v.(proto.Message)
	if !ok {
		return ErrConvertProtoMessage
	}
	return protojson.Unmarshal(data, m)
}

func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	data, err := Marshal(v)
	if err != nil {
		return nil, err
	}
	var dest bytes.Buffer
	if err := json.Indent(&dest, data, prefix, indent); err != nil {
		return nil, err
	}
	return dest.Bytes(), nil
}

func NewEncoder(w io.Writer) encodingx.JSONEncoder {
	return &encoder{w: w}
}

func NewDecoder(r io.Reader) encodingx.Decoder {
	return &decoder{r: r}
}

func MarshalNoEscape(v any) ([]byte, error) {
	return Marshal(v)
}

type encoder struct {
	MarshalOptions protojson.MarshalOptions
	w              io.Writer
}

func (e *encoder) SetIndent(prefix, indent string) {
	e.MarshalOptions.Indent = indent
}

func (e *encoder) SetEscapeHTML(escapeHTML bool) {

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
	UnmarshalOptions protojson.UnmarshalOptions
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
