package telegram

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"io"
	"os/exec"
)

func convertImageToPng(filePath string) (string, error) {
	npath := filePath + ".png"

	cmd := exec.Command("convert",
		"-background", "none",
		"-gravity", "center",
		"-resize", "1024x1024>",
		"-extent", "1:1>",
		filePath, npath)
	err := cmd.Run()

	return npath, err
}

func aiGenerate(ctx context.Context, msg *memory.Message) error {
	request := structs.NewGenerationRequest()

	request.User = msg.User

	request.Prompt = promptFromMessage(msg)
	request.Message = msg
	request.Context = ctx

	w := NewImageWriter(ctx, request, false)

	var ws []io.WriteCloser
	for k := 0; k < config.Telegram.ImageCount; k++ {
		ws = append(ws, w.(*TelegramWriter).Subwriter())
	}

	request.Writers = ws

	channels.GenerationRequestsCh <- request

	return nil
}

func aiImage(ctx context.Context, msg *memory.Message) error {
	for _, path := range msg.GetImages() {

		if msg.Source.Telegram.Update.Message.Caption == "" {
			req := structs.NewVariationRequest()
			req.Context = ctx
			req.ImagePath = path
			req.User = msg.User
			req.Message = msg

			w := NewImageWriter(ctx, req, true)

			var ws []io.WriteCloser
			for k := 0; k < config.Telegram.ImageCount; k++ {
				ws = append(ws, w.(*TelegramWriter).Subwriter())
			}

			req.Writers = ws

			channels.VariationRequestsCh <- req
		} else {
			req := structs.NewGenerationRequest()
			req.Context = ctx
			req.ImagePath = path
			req.Prompt = msg.Source.Telegram.Update.Message.Caption
			req.User = msg.User
			req.Message = msg

			w := NewImageWriter(ctx, req, true)

			var ws []io.WriteCloser
			for k := 0; k < config.Telegram.ImageCount; k++ {
				ws = append(ws, w.(*TelegramWriter).Subwriter())
			}
			req.Writers = ws

			channels.EditRequestsCh <- req
		}
	}

	return nil
}
