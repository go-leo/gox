package jsonx

import "github.com/go-leo/gox/encodingx"

type JSONEncoder interface {
	encodingx.Encoder
	SetIndent(prefix, indent string)
	SetEscapeHTML(escapeHTML bool)
}
