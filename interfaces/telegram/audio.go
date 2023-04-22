package telegram

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/langchain/tts/common"
	"github.com/kamushadenes/chloe/langchain/tts/google"
	"github.com/kamushadenes/chloe/structs"
)

func aiTranscribe(ctx context.Context, msg *memory.Message) error {
	for _, path := range msg.GetAudios() {
		req := structs.NewActionRequest()
		req.Message = msg
		req.Context = ctx
		req.Action = "transcribe"
		req.Params["path"] = path
		req.Writer = NewTelegramWriter(ctx, req, true)

		if err := channels.RunAction(req); err != nil {
			return err
		}
	}

	return nil
}

func aiTTS(ctx context.Context, msg *memory.Message) error {
	req := structs.NewActionRequest()
	req.Message = msg
	req.Context = ctx
	req.Writer = NewTelegramWriter(ctx, req, true)

	t := google.NewTTSGoogle()

	res, err := t.TTSWithContext(ctx, common.TTSMessage{Text: promptFromMessage(msg)})
	if err != nil {
		return err
	}

	_, _ = req.Writer.Write(res.Audio)

	return req.Writer.Close()
}
