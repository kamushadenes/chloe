package tokenizer

import (
	tiktoken "github.com/pkoukk/tiktoken-go"
)

func CountTokens(model string, text string) int {
	tkm, err := tiktoken.EncodingForModel(model)
	if err != nil {
		return -1
	}

	return len(tkm.Encode(text, nil, nil))
}
