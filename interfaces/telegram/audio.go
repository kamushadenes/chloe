package telegram

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rs/zerolog"
	"os/exec"
)

func convertAudioToMp3(ctx context.Context, filePath string) (string, error) {
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		return "", fmt.Errorf("unable to locate `ffmpeg`: %w", err)
	}

	logger := zerolog.Ctx(ctx).With().Str("filePath", filePath).Logger()

	logger.Info().Msg("converting audio to mp3")

	npath := filePath + ".mp3"

	cmd := exec.Command("ffmpeg", "-i", filePath, npath)
	err := cmd.Run()

	return npath, err
}
func aiTranscribe(ctx context.Context, msg *memory.Message, ch chan interface{}) error {
	for _, path := range msg.GetAudios() {
		req := structs.NewActionRequest()
		req.Message = msg
		req.Context = ctx
		req.Action = "transcribe"
		req.Params = path
		req.Writers = append(req.Writers, NewTextWriter(ctx, req, true))

		channels.ActionRequestsCh <- req
	}

	return nil
}

func aiTTS(ctx context.Context, msg *memory.Message) error {
	req := structs.NewActionRequest()
	req.Message = msg
	req.Context = ctx
	req.Action = "tts"
	req.Params = promptFromMessage(msg)
	req.Writers = append(req.Writers, NewAudioWriter(ctx, req, false))

	channels.ActionRequestsCh <- req

	return nil
}
