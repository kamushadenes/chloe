package tokenizer

import (
	"testing"
)

// https://platform.openai.com/tokenizer

func TestGetCompletionMaxTokens(t *testing.T) {
	count := CountTokens("gpt-3.5-turbo", "hello world")
	if count != 2 {
		t.Errorf("GetCompletionMaxTokens() = %v, want %v", count, 2)
	}
}

func TestGetContextSize(t *testing.T) {
	count := GetContextSize("gpt-3.5-turbo")
	if count != 4096 {
		t.Errorf("GetContextSize() = %v, want %v", count, 4096)
	}
}
