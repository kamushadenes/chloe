package telegram

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"os/exec"
)

func convertAudioToMp3(ctx context.Context, filePath string) (string, error) {
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		return "", fmt.Errorf("unable to locate `ffmpeg`: %w", err)
	}

	logger := logging.FromContext(ctx).With().Str("filePath", filePath).Logger()

	logger.Info().Msg("converting audio to mp3")

	npath := filePath + ".mp3"

	cmd := exec.Command("ffmpeg", "-i", filePath, npath)
	err := cmd.Run()

	if err != nil {
		return npath, errors.Wrap(errors.ErrFFmpegError, err)
	}

	return npath, nil
}
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
	req.Action = "tts"
	req.Params["text"] = promptFromMessage(msg)
	req.Writer = NewTelegramWriter(ctx, req, true)

	return channels.RunAction(req)
}
