package textx

import (
	"bytes"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"io"
)

// EUCKRToUtf8 GBK 转 UTF-8
func EUCKRToUtf8(s []byte) ([]byte, error) {
	return io.ReadAll(transform.NewReader(bytes.NewReader(s), korean.EUCKR.NewDecoder()))
}

// Utf8ToEUCKR UTF-8 转 GBK
func Utf8ToEUCKR(s []byte) ([]byte, error) {
	return io.ReadAll(transform.NewReader(bytes.NewReader(s), korean.EUCKR.NewEncoder()))
}
