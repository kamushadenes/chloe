package embedding

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/embeddings/common"
	"github.com/kamushadenes/chloe/langchain/embeddings/openai"
)

func NewEmbedding(model *common.EmbeddingModel) common.Embedding {
	return openai.NewEmbeddingOpenAI(config.OpenAI.APIKey, model)
}

func NewEmbeddingWithDefaultModel() common.Embedding {
	return openai.NewEmbeddingOpenAIWithDefaultModel(config.OpenAI.APIKey)
}
