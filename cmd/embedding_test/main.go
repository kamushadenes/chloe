package main

import (
	"fmt"

	"github.com/kamushadenes/chloe/langchain/embeddings/embedding"
)

func main() {
	embed := embedding.NewEmbeddingWithDefaultModel()

	embeddings, err := embed.Embed([]string{"Hello world!"})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", embeddings)
}
