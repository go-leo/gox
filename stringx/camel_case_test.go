package stringx

import (
	"testing"
)

func TestCamelCase(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello world", "helloWorld"},
		{"foo bar", "fooBar"},
		{"lorem ipsum dolor", "loremIpsumDolor"},
		{"hello_world", "helloWorld"},
		{"snake_case", "snakeCase"},
	}

	for _, testCase := range testCases {
		result := CamelCase(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %s, but got %s for input %s", testCase.expected, result, testCase.input)
		}
	}
}
