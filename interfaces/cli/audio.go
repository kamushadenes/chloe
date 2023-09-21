package cli

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/tts/base"
	"github.com/kamushadenes/chloe/langchain/tts/common"
	"github.com/kamushadenes/chloe/structs/writer_structs"
)

func TTS(ctx context.Context, text string, writer writer_structs.ChloeWriter) error {
	t := base.NewTTSWithDefaultModel(config.TTS.Provider)

	res, err := t.TTSWithContext(ctx, common.TTSMessage{Text: text})
	if err != nil {
		return err
	}

	_, err = writer.Write(res.Audio)

	return err
}
