package condx

import "testing"

// TestAnd is a test function for the And() Operator
func TestAnd(t *testing.T) {
	// Create the And operator
	andOperator := And()

	// Define test cases
	testCases := []struct {
		name     string
		conds    []string
		expected string
	}{
		{"Empty conditions", []string{}, ""},
		{"Single condition", []string{"condition1"}, "condition1"},
		{"Multiple conditions", []string{"condition1", "condition2"}, "condition1 AND condition2"},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := andOperator.Apply(tc.conds)
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}

// TestOr function to test Or() Operator
func TestOr(t *testing.T) {
	orOperator := Or()

	tests := []struct {
		name     string
		conds    []string
		expected string
	}{
		{
			name:     "Empty conditions",
			conds:    []string{},
			expected: "",
		},
		{
			name:     "Single condition",
			conds:    []string{"condition1"},
			expected: "condition1",
		},
		{
			name:     "Multiple conditions",
			conds:    []string{"condition1", "condition2", "condition3"},
			expected: "condition1 OR condition2 OR condition3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := orOperator.Apply(tt.conds)
			if result != tt.expected {
				t.Errorf("Or() = %v, want %v", result, tt.expected)
			}
		})
	}
}
