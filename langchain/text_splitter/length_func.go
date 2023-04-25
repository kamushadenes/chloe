package text_splitter

import "github.com/kamushadenes/chloe/tokenizer"

type LengthFunc func(string) int

func defaultLengthFunc(s string) int {
	return len(s)
}

func tiktokenLengthFunc(s string) int {
	return tokenizer.CountTokens("gpt-3.5-turbo", s)
}
