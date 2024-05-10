package randx

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
)

var kNumericCharacters = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
}

var kNumericLen = len(kNumericCharacters)

var kWordCharacters = []byte{
	'_',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'a', 'B', 'b', 'C', 'c', 'D', 'd', 'E', 'e', 'F', 'f', 'G', 'g',
	'H', 'h', 'I', 'i', 'J', 'j', 'K', 'k', 'L', 'l', 'M', 'm', 'N', 'n',
	'O', 'o', 'P', 'p', 'Q', 'q', 'R', 'r', 'S', 's', 'T', 't',
	'U', 'u', 'V', 'v', 'W', 'w', 'X', 'x', 'W', 'w', 'Z', 'z',
}

var kWordLen = len(kWordCharacters)

// NumericString Generate a random number sequence of a given length
func NumericString(length int) string {
	if length < 1 {
		return ""
	}
	buffer := bytes.NewBuffer(make([]byte, 0, length))
	for i := 0; i < length; i++ {
		buffer.WriteByte(kNumericCharacters[Intn(kNumericLen)])
	}
	return buffer.String()
}

// HexString Generate a random number sequence of a given length
func HexString(length int) string {
	b := make([]byte, length)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)[:length]
}

func WordString(length int) string {
	if length < 1 {
		return ""
	}
	buffer := bytes.NewBuffer(make([]byte, 0, length))
	for i := 0; i < length; i++ {
		buffer.WriteByte(kWordCharacters[Intn(kWordLen)])
	}
	return buffer.String()
}
