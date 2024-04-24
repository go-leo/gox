package stringx

import (
	"strings"
	"unicode"
)

func SnackCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) && !unicode.IsUpper(rune(s[i-1])) {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}
