package textx

import (
	"bytes"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
	"io"
)

// Big5ToUtf8 GBK 转 UTF-8
func Big5ToUtf8(s []byte) ([]byte, error) {
	return io.ReadAll(transform.NewReader(bytes.NewReader(s), traditionalchinese.Big5.NewDecoder()))
}

// Utf8ToBig5 UTF-8 转 GBK
func Utf8ToBig5(s []byte) ([]byte, error) {
	return io.ReadAll(transform.NewReader(bytes.NewReader(s), traditionalchinese.Big5.NewEncoder()))
}
