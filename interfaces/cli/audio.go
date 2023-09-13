package cli

import (
	"context"

	"github.com/kamushadenes/chloe/langchain/tts/common"
	"github.com/kamushadenes/chloe/langchain/tts/google"
	"github.com/kamushadenes/chloe/structs/writer_structs"
)

func TTS(ctx context.Context, text string, writer writer_structs.ChloeWriter) error {
	tts := google.NewTTSGoogle()

	res, err := tts.TTSWithContext(ctx, common.TTSMessage{Text: text})
	if err != nil {
		return err
	}

	_, err = writer.Write(res.Audio)

	return err
}
