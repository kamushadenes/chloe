package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractJSON(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string input",
			input:    "",
			expected: "",
		},
		{
			name:     "Valid JSON string",
			input:    `{"name": "Alice", "age": 30}`,
			expected: `{"name":"Alice","age":30}`,
		},
		{
			name:     "Invalid JSON string",
			input:    `{"name": "Alice",}`,
			expected: "",
		},
		{
			name:     "String with embedded JSON",
			input:    `Something {"name": "Bob", "age": 25} else`,
			expected: `{"name":"Bob","age":25}`,
		},
		{
			name:     "String with multiple JSON",
			input:    `{"a": 1, "b": 2} Something {"name": "Eve", "age": 22}`,
			expected: `{"a":1,"b":2}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ExtractJSON(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
