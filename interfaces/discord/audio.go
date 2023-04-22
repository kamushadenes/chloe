package discord

import (
	"context"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/langchain/tts/common"
	"github.com/kamushadenes/chloe/langchain/tts/google"
	"github.com/kamushadenes/chloe/structs"
)

func tts(ctx context.Context, msg *memory.Message) error {
	t := google.NewTTSGoogle()

	res, err := t.TTSWithContext(ctx, common.TTSMessage{Text: promptFromMessage(msg)})
	if err != nil {
		return err
	}

	req := structs.NewActionRequest()
	req.Message = msg
	req.Context = ctx
	req.Writer = NewDiscordWriter(ctx, req, false)

	_, err = req.Writer.Write(res.Audio)
	if err != nil {
		return err
	}

	return req.Writer.Close()
}
