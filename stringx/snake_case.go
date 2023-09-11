package stringx

import (
	"github.com/go-leo/gox/slicex"
	"strings"
	"unicode"
)

// SnakeCase Converts `string` to [snake case](https://en.wikipedia.org/wiki/Snake_case).
func SnakeCase(str string) string {
	initValue := ""
	return slicex.Reduce[[]string, string](
		strings.FieldsFunc(str, func(r rune) bool { return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsUpper(r) }),
		initValue,
		func(result string, word string, index int, s []string) string {
			word = strings.ToLower(word)
			if index == 0 {
				return result + word
			}
			return result + "_" + word
		},
	)
}
