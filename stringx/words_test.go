package stringx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWords(t *testing.T) {
	assert.Equal(t, []string{"12", "ft"}, Words("12ft"))
	assert.Equal(t, []string{"aeiou", "Are", "Vowels"}, Words("aeiouAreVowels"))
	assert.Equal(t, []string{"enable", "6", "h", "format"}, Words("enable 6h format"))
	assert.Equal(t, []string{"enable", "24", "H", "format"}, Words("enable 24H format"))
	assert.Equal(t, []string{"is", "ISO", "8601"}, Words("isISO8601"))
	assert.Equal(t, []string{"LETTERS", "Aeiou", "Are", "Vowels"}, Words("LETTERSAeiouAreVowels"))
	assert.Equal(t, []string{"too", "Legit", "2", "Quit"}, Words("tooLegit2Quit"))
	assert.Equal(t, []string{"walk", "500", "Miles"}, Words("walk500Miles"))
	assert.Equal(t, []string{"xhr", "2", "Request"}, Words("xhr2Request"))
	assert.Equal(t, []string{"XML", "Http"}, Words("XMLHttp"))
	assert.Equal(t, []string{"Xml", "HTTP"}, Words("XmlHTTP"))
	assert.Equal(t, []string{"Xml", "Http"}, Words("XmlHttp"))
	assert.Equal(t, []string{"LETTERS", "Æiou", "Are", "Vowels"}, Words("LETTERSÆiouAreVowels"))
	assert.Equal(t, []string{"æiou", "Are", "Vowels"}, Words("æiouAreVowels"))
	assert.Equal(t, []string{"æiou", "2", "Consonants"}, Words("æiou2Consonants"))
}
