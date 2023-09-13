package telegram

import (
	"context"

	"github.com/kamushadenes/chloe/langchain/actions"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
)

func aiTranscribe(ctx context.Context, msg *memory.Message) error {
	/*
		for _, path := range msg.GetAudios() {
			tts := google.NewTTSGoogle()

			res, err := tts.TTSWithContext(ctx, common.TTSMessage{Text: promptFromMessage(msg)})
			if err != nil {
				return err
			}

			req := action_structs.NewActionRequest()
			req.Message = msg
			req.Context = ctx
			req.Writer = NewTelegramWriter(ctx, req, true)

			_, err = req.Writer.Write(res.Audio)
			if err != nil {
				return err
			}
		}
	*/

	return nil
}

func aiTTS(ctx context.Context, msg *memory.Message) error {
	req := action_structs.NewActionRequest()
	req.Message = msg
	req.Context = ctx
	req.Action = "tts"
	req.Params["text"] = promptFromMessage(msg)
	req.Writer = NewTelegramWriter(ctx, req, false)

	return actions.HandleAction(req)
}
