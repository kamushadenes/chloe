package telegram

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
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
	request := &structs.GenerationRequest{}

	request.User = msg.User

	request.Prompt = promptFromMessage(msg)

	w := NewImageWriter(ctx, msg, false)

	var ws []io.WriteCloser
	for k := 0; k < 4; k++ {
		ws = append(ws, w.(*TelegramWriter).Subwriter())
	}

	request.Context = ctx
	request.Writers = ws

	channels.GenerationRequestsCh <- request

	return nil
}

func aiImage(ctx context.Context, msg *memory.Message) error {
	for _, path := range msg.GetImages() {
		w := NewImageWriter(ctx, msg, true)

		var ws []io.WriteCloser
		for k := 0; k < 4; k++ {
			ws = append(ws, w.(*TelegramWriter).Subwriter())
		}

		if msg.Source.Telegram.Update.Message.Caption == "" {
			channels.VariationRequestsCh <- &structs.VariationRequest{
				Context:   ctx,
				ImagePath: path,
				User:      msg.User,
				Writers:   ws,
			}
		} else {
			channels.EditRequestsCh <- &structs.GenerationRequest{
				Context:   ctx,
				ImagePath: path,
				Prompt:    msg.Source.Telegram.Update.Message.Caption,
				User:      msg.User,
				Writers:   ws,
			}
		}
	}

	return nil
}
