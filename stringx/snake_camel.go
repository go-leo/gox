package stringx

import (
	"strings"
	"unicode"
)

const lowLine = '_'

// Camel2Snake 驼峰转蛇形
func Camel2Snake(s string) string {
	output := strings.Builder{}
	var preLowLine bool
	sr := []rune(s)
	for i := 0; i < len(sr); i++ {
		r := sr[i]
		if unicode.IsLower(r) {
			if preLowLine {
				preLowLine = false
				continue
			}
			output.WriteRune(r)
			continue
		}
		if unicode.IsUpper(r) {
			if preLowLine {
				output.WriteRune(unicode.ToLower(r))
				preLowLine = false
				continue
			}
			if output.Len() == 0 {
				output.WriteRune(unicode.ToLower(r))
				continue
			}
			output.WriteRune(lowLine)
			output.WriteRune(unicode.ToLower(r))
			continue
		}
		if unicode.IsSpace(r) {
			if preLowLine {
				continue
			}
			output.WriteRune(lowLine)
			preLowLine = true
		}
		if r == lowLine {
			if preLowLine {
				continue
			}
			output.WriteRune(lowLine)
			preLowLine = true
			continue
		}
		output.WriteRune(r)
	}
	return output.String()
}

// Snake2Camel 蛇形转驼峰, cases首字母是否大小写转换， true->大写，false->小写
func Snake2Camel(s string, cases bool) string {
	output := strings.Builder{}
	wordStart := true
	sr := []rune(s)
	for i := 0; i < len(sr); i++ {
		r := sr[i]
		if unicode.IsSpace(r) {
			if wordStart {
				continue
			}
			wordStart = true
			continue
		}
		if r == lowLine {
			if wordStart {
				continue
			}
			wordStart = true
			continue
		}
		if !wordStart {
			output.WriteRune(r)
			continue
		}
		if output.Len() > 0 {
			output.WriteRune(unicode.ToUpper(r))
			wordStart = false
			continue
		}
		if cases {
			output.WriteRune(unicode.ToUpper(r))
			wordStart = false
			continue
		}
		output.WriteRune(unicode.ToLower(r))
		wordStart = false
		continue
	}
	return output.String()
}
