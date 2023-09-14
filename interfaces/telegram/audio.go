package telegram

import (
	"context"
	"github.com/kamushadenes/chloe/langchain/actions"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/writer_structs"
)

func aiTranscribe(ctx context.Context, msg *memory.Message) error {
	for _, path := range msg.GetAudios() {

		req := action_structs.NewActionRequest()
		req.Message = msg
		req.Context = ctx
		req.Action = "transcribe"
		req.Params["path"] = path

		w := writer_structs.NewMockWriter()
		req.Writer = w

		if err := actions.HandleAction(req); err != nil {
			return err
		}

		tw := NewTelegramWriter(ctx, req, true)
		if _, err := tw.Write(w.GetObjects()[0].Bytes()); err != nil {
			return err
		}
		if err := tw.Close(); err != nil {
			return err
		}

		msg.ID = 0
		msg.Content = w.GetObjects()[0].String()
		msg.Source.Telegram.Update.Message.Text = w.GetObjects()[0].String()
		if err := msg.Save(ctx); err != nil {
			return err
		}

		return processText(ctx, msg, nil)
	}

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
