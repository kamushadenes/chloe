package openai

import "github.com/sashabaranov/go-openai"

var modelStringToEnum = map[string]openai.EmbeddingModel{
	"text-embedding-ada-002": openai.AdaEmbeddingV2,
}
