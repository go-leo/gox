package internal

import "regexp"

var apostropheRegexp = regexp.MustCompile("['\u2019]")

func RemoveApostrophe(str string) string {
	return apostropheRegexp.ReplaceAllString(str, "")
}
