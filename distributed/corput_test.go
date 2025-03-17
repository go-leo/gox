package distributed

import (
	"testing"
)

func TestCorput(t *testing.T) {
	testCases := []struct {
		n        int
		base     int
		expected float64
	}{
		{10, 2, 0.7213475218575672},
		{100, 10, 0.5849624838248488},
		{1000, 10, 0.5440211108893698},
		{10000, 2, 0.4150375161751343},
	}

	for _, tc := range testCases {
		actual := VanDerCorputSequence(tc.n, tc.base)
		if actual != tc.expected {
			t.Errorf("VanDerCorputSequence(%d, %d) = %f, expected %f", tc.n, tc.base, actual, tc.expected)
		}
	}
}
