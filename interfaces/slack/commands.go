package slack

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rs/zerolog"
	"io"
	"os"
	"strings"
)

func complete(ctx context.Context, msg *memory.Message) {
	logger := zerolog.Ctx(ctx)

	if err := aiComplete(ctx, msg); err != nil {
		logger.Error().Err(err).Msg("error generating image")
	}
}

func generate(ctx context.Context, msg *memory.Message) {
	logger := zerolog.Ctx(ctx)

	if err := aiGenerate(ctx, msg); err != nil {
		logger.Error().Err(err).Msg("error generating image")
	}
}

func tts(ctx context.Context, msg *memory.Message) {
	logger := zerolog.Ctx(ctx)

	if err := aiTTS(ctx, msg); err != nil {
		logger.Error().Err(err).Msg("error generating audio")
	}
}

func forgetUser(ctx context.Context, msg *memory.Message) error {
	return msg.User.DeleteMessages(ctx)
}

func action(ctx context.Context, msg *memory.Message) {
	fields := strings.Fields(msg.Content)

	req := structs.NewActionRequest()
	req.Context = ctx
	req.Action = fields[0]
	req.Params = strings.Join(fields[1:], " ")
	req.Thought = fmt.Sprintf("User wants to run action %s", fields[0])
	req.Writers = []io.WriteCloser{os.Stdout}

	channels.ActionRequestsCh <- req
}
