package utils

import (
	"strings"
)

// StringToChunks splits a string into chunks of a given size
// Not very efficient, and very ugly, but it works
func StringToChunks(s string, chunkSize int) []string {
	if len(s) <= chunkSize {
		return []string{s}
	}

	var chunks []string
	start := 0
	for start < len(s) {
		end := start + chunkSize
		if end > len(s) {
			end = len(s)
		}

		// Move the end index back to the previous newline, if possible
		if end < len(s) && s[end-1] != '\n' {
			newEnd := strings.LastIndex(s[start:end], "\n")
			if newEnd != -1 {
				end = start + newEnd + 1
			}
		}

		chunks = append(chunks, s[start:end])
		start = end
	}

	return chunks
}
