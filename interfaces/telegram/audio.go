package telegram

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/i18n"
	"github.com/kamushadenes/chloe/memory"
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
func aiTranscribe(ctx context.Context, msg *memory.Message, ch chan interface{}) error {
	_ = msg.SendText(i18n.GetTranscriptionText(), msg.Source.Telegram.Update.Message.MessageID)

	for _, path := range msg.GetAudios() {
		request := structs.NewTranscriptionRequest()

		request.FilePath = path
		request.ResultChannel = ch

		request.User = msg.User
		request.Message = msg

		request.Context = ctx
		request.Writer = NewTextWriter(ctx, request, true)

		channels.TranscribeRequestsCh <- request
	}

	return nil
}

func aiTTS(ctx context.Context, msg *memory.Message) error {
	_ = msg.SendText(i18n.GetTextToSpeechText(), msg.Source.Telegram.Update.Message.MessageID)

	request := structs.NewTTSRequest()

	request.User = msg.User

	request.Content = promptFromMessage(msg)
	request.Message = msg
	request.Context = ctx

	w := NewAudioWriter(ctx, request, false)
	request.Writers = append(request.Writers, w.(io.WriteCloser))

	channels.TTSRequestsCh <- request

	return nil
}
