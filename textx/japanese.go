package textx

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
)

// EUCJPToUtf8 GBK 转 UTF-8
func EUCJPToUtf8(s []byte) ([]byte, error) {
	return io.ReadAll(transform.NewReader(bytes.NewReader(s), japanese.EUCJP.NewDecoder()))
}

// Utf8ToEUCJP UTF-8 转 GBK
func Utf8ToEUCJP(s []byte) ([]byte, error) {
	return io.ReadAll(transform.NewReader(bytes.NewReader(s), japanese.EUCJP.NewEncoder()))
}

// ISO2022JPToUtf8 GBK 转 UTF-8
func ISO2022JPToUtf8(s []byte) ([]byte, error) {
	return io.ReadAll(transform.NewReader(bytes.NewReader(s), japanese.ISO2022JP.NewDecoder()))
}

// Utf8ToISO2022JP UTF-8 转 GBK
func Utf8ToISO2022JP(s []byte) ([]byte, error) {
	return io.ReadAll(transform.NewReader(bytes.NewReader(s), japanese.ISO2022JP.NewEncoder()))
}

// ShiftJISToUtf8 GBK 转 UTF-8
func ShiftJISToUtf8(s []byte) ([]byte, error) {
	return io.ReadAll(transform.NewReader(bytes.NewReader(s), japanese.ShiftJIS.NewDecoder()))
}

// Utf8ToShiftJIS UTF-8 转 GBK
func Utf8ToShiftJIS(s []byte) ([]byte, error) {
	return io.ReadAll(transform.NewReader(bytes.NewReader(s), japanese.ShiftJIS.NewEncoder()))
}
