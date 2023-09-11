package stringx

import (
	"testing"
)

func TestSnakeCase(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello world", "hello_world"},
		{"foo bar", "foo_bar"},
		{"lorem ipsum dolor", "lorem_ipsum_dolor"},
		{"HelloWorld", "hello_world"},
		{"snakeCase", "snake_case"},
	}

	for _, testCase := range testCases {
		result := SnakeCase(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %s, but got %s for input %s", testCase.expected, result, testCase.input)
		}
	}
}
