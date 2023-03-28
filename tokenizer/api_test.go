package tokenizer

import (
	"github.com/kamushadenes/chloe/models"
	"testing"
)

// https://platform.openai.com/tokenizer

func TestGetCompletionMaxTokens(t *testing.T) {
	count := CountTokens(models.GPT35Turbo, "hello world")
	if count != 2 {
		t.Errorf("GetCompletionMaxTokens() = %v, want %v", count, 2)
	}
}

func TestGetContextSize(t *testing.T) {
	count := GetContextSize(models.GPT35Turbo)
	if count != 4096 {
		t.Errorf("GetContextSize() = %v, want %v", count, 4096)
	}
}
