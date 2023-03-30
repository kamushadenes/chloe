package telegram

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"os/exec"
	"strings"
)

func convertImageToPng(filePath string) (string, error) {
	if _, err := exec.LookPath("convert"); err != nil {
		return "", fmt.Errorf("unable to locate `convert`: %w", err)
	}

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

func aiAction(ctx context.Context, msg *memory.Message) error {
	req := structs.NewActionRequest()
	req.Message = msg
	req.Context = ctx
	req.Action = "image"
	req.Params = promptFromMessage(msg)
	req.Writers = append(req.Writers, NewImageWriter(ctx, req, false))

	channels.ActionRequestsCh <- req

	return nil
}

func aiGenerate(ctx context.Context, msg *memory.Message) error {
	fields := strings.Fields(msg.Content)
	req := structs.NewActionRequest()
	req.Context = ctx
	req.Action = fields[0]
	req.Params = strings.Join(fields[1:], " ")
	req.Thought = fmt.Sprintf("User wants to run action %s", fields[0])
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
