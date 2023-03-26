package telegram

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
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
	req := structs.NewActionRequest()
	req.Message = msg
	req.Context = ctx
	req.Action = "image"
	req.Params = promptFromMessage(msg)
	req.Writers = append(req.Writers, NewImageWriter(ctx, req, false))

	channels.ActionRequestsCh <- req

	return nil
}

func aiImage(ctx context.Context, msg *memory.Message) error {
	for _, path := range msg.GetImages() {
		req := structs.NewActionRequest()
		req.Message = msg
		req.Context = ctx
		req.Action = "variation"
		req.Params = path
		req.Writers = append(req.Writers, NewImageWriter(ctx, req, false))

		channels.ActionRequestsCh <- req
	}

	return nil
}
