package stringx

func ReplaceFunc(str string, f func(rune) string) string {
	var b Builder
	for _, r := range str {
		_, _ = b.WriteString(f(r))
	}
	return b.String()
}
