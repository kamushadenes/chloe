package discord

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)

func aiGenerate(ctx context.Context, msg *memory.Message) error {
	request := structs.NewGenerationRequest()

	request.User = msg.User

	request.Prompt = promptFromMessage(msg)
	request.Message = msg
	request.Context = ctx

	w := NewImageWriter(ctx, request, false, request.Prompt)

	for k := 0; k < config.Discord.ImageCount; k++ {
		request.Writers = append(request.Writers, w.(*DiscordWriter).Subwriter())
	}

	channels.GenerationRequestsCh <- request

	return nil
}
