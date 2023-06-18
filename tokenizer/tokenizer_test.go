package tokenizer

import (
	"testing"
)

func TestGetCompletionMaxTokens(t *testing.T) {
	count := CountTokens("gpt-3.5-turbo", "hello world")
	if count != 2 {
		t.Errorf("GetCompletionMaxTokens() = %v, want %v", count, 2)
	}
}
