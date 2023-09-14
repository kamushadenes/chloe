package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundFloat(t *testing.T) {
	testCases := []struct {
		name      string
		input     float64
		precision uint
		expected  float64
	}{
		{
			name:      "Round with zero precision",
			input:     3.14159,
			precision: 0,
			expected:  3,
		},
		{
			name:      "Round with one decimal precision",
			input:     3.14159,
			precision: 1,
			expected:  3.1,
		},
		{
			name:      "Round with two decimal precision",
			input:     3.14159,
			precision: 2,
			expected:  3.14,
		},
		{
			name:      "Round with three decimal precision",
			input:     3.14159,
			precision: 3,
			expected:  3.142,
		},
		{
			name:      "Round a negative number with two decimal precision",
			input:     -3.14159,
			precision: 2,
			expected:  -3.14,
		},
		{
			name:      "Round a large number with precision",
			input:     123456789.98765,
			precision: 3,
			expected:  123456789.988,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := RoundFloat(tc.input, tc.precision)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
