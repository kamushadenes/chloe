package common

type Embedding interface {
	Embed([]string) (EmbeddingResult, error)
}
