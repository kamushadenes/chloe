package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToChunks(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		chunkSize int
		expected  []string
	}{
		{
			name:      "Short string with no newlines",
			input:     "Hello, world!",
			chunkSize: 10,
			expected: []string{
				"Hello, wor",
				"ld!",
			},
		},
		{
			name:      "String with newlines",
			input:     "Hello, world!\nWelcome to the test\n",
			chunkSize: 30,
			expected: []string{
				"Hello, world!\n", "Welcome to the test\n",
			},
		},
		{
			name:      "String with newlines and chunk break required",
			input:     "Hello, world!\nWelcome to the test",
			chunkSize: 15,
			expected: []string{
				"Hello, world!\n",
				"Welcome to the ",
				"test",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := StringToChunks(tc.input, tc.chunkSize)
			assert.Equal(t, tc.expected, actual)

			// Check that none of the chunks exceed the chunkSize
			for _, chunk := range actual {
				assert.True(t, len(chunk) <= tc.chunkSize)

				// Check that no chunks break within words (if possible)
				if tc.chunkSize > len(chunk) {
					indexOfLastSpace := strings.LastIndex(chunk, " ")
					if indexOfLastSpace > 0 {
						assert.True(t, tc.chunkSize-len(chunk) <= indexOfLastSpace+len(chunk))
					}
				}
			}
		})
	}
}
