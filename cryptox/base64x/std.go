package base64x

import (
	"github.com/go-leo/gox/encodingx/base64x"
)

// Deprecated: use github.com/go-leo/gox/encodingx/base64x
func StdEncode(src []byte) string {
	return base64x.StdEncode(src)
}

// Deprecated: use github.com/go-leo/gox/encodingx/base64x
func StdDecode(s string) ([]byte, error) {
	return base64x.StdDecode(s)
}
