package telegram

import (
	"context"
	"strings"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/actions"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
)

func aiAction(ctx context.Context, msg *memory.Message) error {
	fields := strings.Fields(promptFromMessage(msg))

	req := action_structs.NewActionRequest()
	req.Message = msg
	req.Context = ctx
	req.Action = fields[0]
	req.Params["prompt"] = strings.Join(fields[1:], " ")
	req.Writer = NewTelegramWriter(ctx, msg, false)
	req.Count = config.Discord.ImageCount

	return actions.HandleAction(req)
}

func aiGenerate(ctx context.Context, msg *memory.Message) error {
	req := action_structs.NewActionRequest()
	req.Context = ctx
	req.Message = msg
	req.Action = "generate"
	req.Params["prompt"] = promptFromMessage(msg)
	req.Writer = NewTelegramWriter(ctx, msg, false)
	req.Count = config.Telegram.ImageCount

	return actions.HandleAction(req)
}

func aiImage(ctx context.Context, msg *memory.Message) error {
	for _, path := range msg.GetImages() {
		req := action_structs.NewActionRequest()
		req.Message = msg
		req.Context = ctx
		req.Action = "variation"
		req.Params["path"] = path
		req.Writer = NewTelegramWriter(ctx, msg, false)
		req.Count = config.Telegram.ImageCount

		/*if err := structs.RunAction(req); err != nil {
			return err
		}*/
	}

	return nil
}
