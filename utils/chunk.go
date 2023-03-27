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

	lines := strings.Split(s, "\n")

	var strs [][]string
	var idx = 0

	for k := range lines {
		if len(strs) <= idx {
			strs = append(strs, make([]string, 0))
		}

		line := lines[k]

		if len(strings.Join(strs[idx], ""))+len(line) <= chunkSize {
			strs[idx] = append(strs[idx], line)
			continue
		} else {
			idx++
			if len(strs) <= idx {
				strs = append(strs, make([]string, 0))
			}
			strs[idx] = append(strs[idx], line)
		}
	}

	var respStrs []string

	for _, str := range strs {
		respStrs = append(respStrs, strings.Join(str, "\n"))
	}

	return respStrs
}
