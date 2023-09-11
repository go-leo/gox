package stringx

import "regexp"

func Match(str string, pattern string) []string {
	return regexp.MustCompile(pattern).FindAllString(str, -1)
}
