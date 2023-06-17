package common

import "time"

type EmbeddingOptions interface {
	WithChunkSize(int) EmbeddingOptions
	WithMaxRetries(int) EmbeddingOptions
	WithTimeout(time.Duration) EmbeddingOptions
	GetTimeout() time.Duration
	GetRequest() interface{}
	WithText([]string) EmbeddingOptions
	GetText() []string
	WithModel(string) EmbeddingOptions
}
