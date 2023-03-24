package envx

import (
	"github.com/caarlos0/env/v7"

	"github.com/go-leo/gox/encodingx"
)

func Unmarshal(prefix string, v any) error {
	return env.Parse(v, env.Options{Prefix: prefix})
}

func NewDecoder(prefix string) encodingx.Decoder {
	return &decoder{prefix: prefix}
}

type decoder struct {
	prefix string
}

func (d *decoder) Decode(obj interface{}) error {
	return env.Parse(obj, env.Options{Prefix: d.prefix})
}
