package slack

import (
	"context"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/langchain/tts/common"
	"github.com/kamushadenes/chloe/langchain/tts/google"
	"github.com/kamushadenes/chloe/structs"
)

func tts(ctx context.Context, msg *memory.Message) error {
	req := structs.NewActionRequest()
	req.Message = msg
	req.Context = ctx
	req.Writer = NewSlackWriter(ctx, req, false)

	t := google.NewTTSGoogle()

	res, err := t.TTSWithContext(ctx, common.TTSMessage{Text: promptFromMessage(msg)})
	if err != nil {
		return err
	}

	_, _ = req.Writer.Write(res.Audio)

	return req.Writer.Close()
}
