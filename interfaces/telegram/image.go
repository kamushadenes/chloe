package telegram

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs"
	"strings"
)

func aiAction(ctx context.Context, msg *memory.Message) error {
	fields := strings.Fields(promptFromMessage(msg))

	req := structs.NewActionRequest()
	req.Message = msg
	req.Context = ctx
	req.Action = fields[0]
	req.Params["text"] = strings.Join(fields[1:], " ")
	req.Thought = fmt.Sprintf("User wants to run action %s", fields[0])
	req.Writer = NewTelegramWriter(ctx, req, false)
	req.Count = config.Telegram.ImageCount

	return channels.RunAction(req)
}

func aiGenerate(ctx context.Context, msg *memory.Message) error {
	req := structs.NewActionRequest()
	req.Context = ctx
	req.Action = "generate"
	req.Params["prompt"] = promptFromMessage(msg)
	req.Message = msg
	req.Writer = NewTelegramWriter(ctx, req, false)
	req.Count = config.Telegram.ImageCount

	return channels.RunAction(req)
}

func aiImage(ctx context.Context, msg *memory.Message) error {
	for _, path := range msg.GetImages() {
		req := structs.NewActionRequest()
		req.Message = msg
		req.Context = ctx
		req.Action = "variation"
		req.Params["path"] = path
		req.Writer = NewTelegramWriter(ctx, req, false)
		req.Count = config.Telegram.ImageCount

		if err := channels.RunAction(req); err != nil {
			return err
		}
	}

	return nil
}
