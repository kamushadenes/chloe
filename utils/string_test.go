package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTruncate(t *testing.T) {
	assert.Equal(t, "abcd", Truncate("abcdefg", 4))
	assert.Equal(t, "abcdefgh", Truncate("abcdefgh", 10))
}
