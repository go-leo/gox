package md5x

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func MD5(data []byte) []byte {
	hash := md5.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func MD5Hex(data []byte) string {
	return hex.EncodeToString(MD5(data))
}

func TextMD5(text string) []byte {
	hash := md5.New()
	_, _ = io.WriteString(hash, text)
	return hash.Sum(nil)
}

func TextMD5Hex(text string) string {
	return hex.EncodeToString(TextMD5(text))
}
