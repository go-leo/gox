package formx

import (
	"github.com/go-leo/gox/encodingx"
	"github.com/go-playground/form/v4"
	"io"
	"net/url"
)

func Marshal(encoders ...*form.Encoder) func(v any) ([]byte, error) {
	return func(v any) ([]byte, error) {
		var encoder *form.Encoder
		if len(encoders) > 0 {
			encoder = encoders[0]
		} else {
			encoder = form.NewEncoder()
		}
		values, err := encoder.Encode(v)
		if err != nil {
			return nil, err
		}
		return []byte(values.Encode()), nil
	}
}

func Unmarshal(decoders ...*form.Decoder) func(data []byte, v any) error {
	return func(data []byte, v any) error {
		var decoder *form.Decoder
		if len(decoders) > 0 {
			decoder = decoders[0]
		} else {
			decoder = form.NewDecoder()
		}
		values, err := url.ParseQuery(string(data))
		if err != nil {
			return err
		}
		return decoder.Decode(v, values)
	}
}

func NewEncoder(encoders ...*form.Encoder) func(w io.Writer) encodingx.Encoder {
	var enc *form.Encoder
	if len(encoders) > 0 {
		enc = encoders[0]
	} else {
		enc = form.NewEncoder()
	}
	return func(w io.Writer) encodingx.Encoder {
		return &encoder{w: w, enc: enc}
	}
}

func NewDecoder(decoders ...*form.Decoder) func(r io.Reader) encodingx.Decoder {
	var dec *form.Decoder
	if len(decoders) > 0 {
		dec = decoders[0]
	} else {
		dec = form.NewDecoder()
	}
	return func(r io.Reader) encodingx.Decoder {
		return &decoder{r: r, dec: dec}
	}
}

type encoder struct {
	enc *form.Encoder
	w   io.Writer
}

func (e *encoder) Encode(val any) error {
	values, err := e.enc.Encode(val)
	if err != nil {
		return err
	}
	data := values.Encode()
	_, err = e.w.Write([]byte(data))
	return err
}

type decoder struct {
	r   io.Reader
	dec *form.Decoder
}

func (d *decoder) Decode(obj any) error {
	data, err := io.ReadAll(d.r)
	if err != nil {
		return nil
	}
	values, err := url.ParseQuery(string(data))
	if err != nil {
		return err
	}
	return d.dec.Decode(obj, values)
}
