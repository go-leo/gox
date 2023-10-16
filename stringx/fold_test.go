// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stringx

import (
	"bytes"
	"strings"
	"testing"
	"unicode/utf8"
)

var foldTests = []struct {
	fn   func(s, t []byte) bool
	s, t string
	want bool
}{
	{EqualFoldRight, "", "", true},
	{EqualFoldRight, "a", "a", true},
	{EqualFoldRight, "", "a", false},
	{EqualFoldRight, "a", "", false},
	{EqualFoldRight, "a", "A", true},
	{EqualFoldRight, "AB", "ab", true},
	{EqualFoldRight, "AB", "ac", false},
	{EqualFoldRight, "sbkKc", "ſbKKc", true},
	{EqualFoldRight, "SbKkc", "ſbKKc", true},
	{EqualFoldRight, "SbKkc", "ſbKK", false},
	{EqualFoldRight, "e", "é", false},
	{EqualFoldRight, "s", "S", true},

	{SimpleLetterEqualFold, "", "", true},
	{SimpleLetterEqualFold, "abc", "abc", true},
	{SimpleLetterEqualFold, "abc", "ABC", true},
	{SimpleLetterEqualFold, "abc", "ABCD", false},
	{SimpleLetterEqualFold, "abc", "xxx", false},

	{AsciiEqualFold, "a_B", "A_b", true},
	{AsciiEqualFold, "aa@", "aa`", false}, // verify 0x40 and 0x60 aren't case-equivalent
}

func TestFold(t *testing.T) {
	for i, tt := range foldTests {
		if got := tt.fn([]byte(tt.s), []byte(tt.t)); got != tt.want {
			t.Errorf("%d. %q, %q = %v; want %v", i, tt.s, tt.t, got, tt.want)
		}
		truth := strings.EqualFold(tt.s, tt.t)
		if truth != tt.want {
			t.Errorf("strings.EqualFold doesn't agree with case %d", i)
		}
	}
}

func TestFoldAgainstUnicode(t *testing.T) {
	var buf1, buf2 []byte
	var runes []rune
	for i := 0x20; i <= 0x7f; i++ {
		runes = append(runes, rune(i))
	}
	runes = append(runes, kelvin, smallLongEss)

	funcs := []struct {
		name   string
		fold   func(s, t []byte) bool
		letter bool // must be ASCII letter
		simple bool // must be simple ASCII letter (not 'S' or 'K')
	}{
		{
			name: "EqualFoldRight",
			fold: EqualFoldRight,
		},
		{
			name:   "AsciiEqualFold",
			fold:   AsciiEqualFold,
			simple: true,
		},
		{
			name:   "SimpleLetterEqualFold",
			fold:   SimpleLetterEqualFold,
			simple: true,
			letter: true,
		},
	}

	for _, ff := range funcs {
		for _, r := range runes {
			if r >= utf8.RuneSelf {
				continue
			}
			if ff.letter && !isASCIILetter(byte(r)) {
				continue
			}
			if ff.simple && (r == 's' || r == 'S' || r == 'k' || r == 'K') {
				continue
			}
			for _, r2 := range runes {
				buf1 = append(utf8.AppendRune(append(buf1[:0], 'x'), r), 'x')
				buf2 = append(utf8.AppendRune(append(buf2[:0], 'x'), r2), 'x')
				want := bytes.EqualFold(buf1, buf2)
				if got := ff.fold(buf1, buf2); got != want {
					t.Errorf("%s(%q, %q) = %v; want %v", ff.name, buf1, buf2, got, want)
				}
			}
		}
	}
}

func isASCIILetter(b byte) bool {
	return ('A' <= b && b <= 'Z') || ('a' <= b && b <= 'z')
}
