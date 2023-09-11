package stringx

import (
	"strings"
	"unicode"
)

func Words(str string) []string {
	words := strings.FieldsFunc(str, func(r rune) bool {
		return unicode.IsSpace(r) || unicode.IsControl(r) || unicode.IsMark(r) || unicode.IsPunct(r)
	})
	strings.Map()
	return words
}
