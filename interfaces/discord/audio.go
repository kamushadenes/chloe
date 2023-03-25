package discord

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"io"
)

func aiTranscribe(ctx context.Context, msg *memory.Message, ch chan interface{}) error {
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
