package cli

import (
	"context"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/diffusion_models/common"
	"github.com/kamushadenes/chloe/langchain/diffusion_models/openai"
	"github.com/kamushadenes/chloe/structs/writer_structs"
)

func Generate(ctx context.Context, text string, writer writer_structs.ChloeWriter) error {
	dif := openai.NewDiffusionOpenAI(config.OpenAI.APIKey)

	res, err := dif.GenerateWithContext(ctx, common.DiffusionMessage{Prompt: text})
	if err != nil {
		return err
	}

	for k := range res.Images {
		_, err = writer.Write([]byte(res.Images[k]))
		if err != nil {
			return err
		}
	}

	return nil
}
