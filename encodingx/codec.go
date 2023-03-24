package encodingx

import "io"

type Encoder interface {
	Encode(val interface{}) error
}

type Decoder interface {
	Decode(obj interface{}) error
}

type NewDecoder func(r io.Reader) Decoder
