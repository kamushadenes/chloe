package utils

import "io"

func WritersOrDefault(writers []io.WriteCloser, fallback io.WriteCloser) []io.WriteCloser {
	if len(writers) == 0 {
		return []io.WriteCloser{fallback}
	}

	return writers
}
