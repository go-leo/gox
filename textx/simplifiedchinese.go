package textx

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
)

// GBKToUtf8 GBK 转 UTF-8
func GBKToUtf8(s []byte) ([]byte, error) {
	return io.ReadAll(transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder()))
}

// Utf8ToGBK UTF-8 转 GBK
func Utf8ToGBK(s []byte) ([]byte, error) {
	return io.ReadAll(transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder()))
}

// GB18030ToUtf8 GBK 转 UTF-8
func GB18030ToUtf8(s []byte) ([]byte, error) {
	return io.ReadAll(transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewDecoder()))
}

// Utf8ToGB18030 UTF-8 转 GB18030
func Utf8ToGB18030(s []byte) ([]byte, error) {
	return io.ReadAll(transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewEncoder()))
}

// HZGB2312ToUtf8 GBK 转 UTF-8
func HZGB2312ToUtf8(s []byte) ([]byte, error) {
	return io.ReadAll(transform.NewReader(bytes.NewReader(s), simplifiedchinese.HZGB2312.NewDecoder()))
}

// Utf8ToHZGB2312 UTF-8 转 GB18030
func Utf8ToHZGB2312(s []byte) ([]byte, error) {
	return io.ReadAll(transform.NewReader(bytes.NewReader(s), simplifiedchinese.HZGB2312.NewEncoder()))
}
