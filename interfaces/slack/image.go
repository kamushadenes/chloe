package slack

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/diffusion_models/openai"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs"
)

func generate(ctx context.Context, msg *memory.Message) error {
	req := structs.NewActionRequest()
	req.Message = msg
	req.Context = ctx
	req.Writer = NewSlackWriter(ctx, req, false, req.Params["prompt"])
	req.Count = config.Slack.ImageCount

	dif := openai.NewDiffusionOpenAI(config.OpenAI.APIKey)

	opts := openai.NewDiffusionOptionsOpenAI()

	res, err := dif.GenerateWithOptions(ctx, opts.
		WithPrompt(promptFromMessage(msg)).
		WithCount(config.Slack.ImageCount))
	if err != nil {
		return err
	}

	for _, v := range res.Images {
		obj := structs.ResponseObject{}

		_, _ = obj.Write(v)
		err = req.Writer.WriteObject(&obj)
		if err != nil {
			return err
		}
	}

	return req.Writer.Close()
}
