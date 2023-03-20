package telegram

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/messages"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rs/zerolog"
	"io"
	"os/exec"
)

func convertAudioToMp3(ctx context.Context, filePath string) (string, error) {
	logger := zerolog.Ctx(ctx).With().Str("filePath", filePath).Logger()

	logger.Info().Msg("converting audio to mp3")

	npath := filePath + ".mp3"

	cmd := exec.Command("ffmpeg", "-i", filePath, npath)
	err := cmd.Run()

	return npath, err
}
func aiTranscribe(ctx context.Context, msg *messages.Message, ch chan interface{}) error {
	for _, path := range msg.GetAudios() {
		request := &structs.TranscriptionRequest{
			FilePath:      path,
			ResultChannel: ch,
		}

		request.User = msg.User

		request.Context = ctx
		request.Writer = NewTextWriter(ctx, msg, true)

		channels.TranscribeRequestsCh <- request
	}

	return nil
}

func aiTTS(ctx context.Context, msg *messages.Message) error {
	request := &structs.TTSRequest{}

	request.User = msg.User

	request.Content = promptFromMessage(msg)

	w := NewAudioWriter(ctx, msg, false)

	request.Context = ctx
	request.Writers = []io.WriteCloser{w}

	channels.TTSRequestsCh <- request

	return nil
}
