package stringx

import (
	"github.com/go-leo/gox/slicex"
	"strings"
	"unicode"
)

// CamelCase Converts `string` to [camel case](https://en.wikipedia.org/wiki/CamelCase).
func CamelCase(str string) string {
	initValue := ""
	return slicex.Reduce[[]string, string](
		strings.FieldsFunc(str, func(r rune) bool { return unicode.IsSpace(r) || unicode.IsPunct(r) }),
		initValue,
		func(result string, word string, index int, s []string) string {
			word = strings.ToLower(word)
			if index == 0 {
				return result + word
			}
			return result + UpperFirst(word)
		},
	)
}
