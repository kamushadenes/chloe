package utils

import "testing"

import (
	"github.com/stretchr/testify/assert"
)

func TestSubtractInt(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Subtract single number",
			input:    []int{5},
			expected: -5,
		},
		{
			name:     "Subtract multiple numbers",
			input:    []int{5, 10, 15},
			expected: -30,
		},
		{
			name:     "Subtract negative numbers",
			input:    []int{-5, -10, -15},
			expected: 30,
		},
		{
			name:     "Subtract mixed numbers",
			input:    []int{5, -10, 15, -20},
			expected: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := SubtractInt(tc.input...)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestSubtractIntWithMinimum(t *testing.T) {
	testCases := []struct {
		name     string
		min      int
		input    []int
		expected int
	}{
		{
			name:     "Subtract with result above minimum",
			min:      -10,
			input:    []int{5, 10},
			expected: -10,
		},
		{
			name:     "Subtract with result below minimum",
			min:      -10,
			input:    []int{5, 20},
			expected: -10,
		},
		{
			name:     "Subtract with result equal to minimum",
			min:      -15,
			input:    []int{5, 10},
			expected: -15,
		},
		{
			name:     "Subtract with negative minimum",
			min:      -50,
			input:    []int{-5, -10, -15},
			expected: 30,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := SubtractIntWithMinimum(tc.min, tc.input...)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
