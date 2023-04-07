package telegram

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
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

	if err != nil {
		return npath, errors.Wrap(errors.ErrImageMagickError, err)
	}

	return npath, nil
}

func aiAction(ctx context.Context, msg *memory.Message) error {
	fields := strings.Fields(promptFromMessage(msg))

	req := structs.NewActionRequest()
	req.Message = msg
	req.Context = ctx
	req.Action = fields[0]
	req.Params["text"] = strings.Join(fields[1:], " ")
	req.Thought = fmt.Sprintf("User wants to run action %s", fields[0])
	req.Writer = NewTelegramWriter(ctx, req, false)
	req.Count = config.Telegram.ImageCount

	return channels.RunAction(req)
}

func aiGenerate(ctx context.Context, msg *memory.Message) error {
	req := structs.NewActionRequest()
	req.Context = ctx
	req.Action = "generate"
	req.Params["prompt"] = promptFromMessage(msg)
	req.Message = msg
	req.Writer = NewTelegramWriter(ctx, req, false)
	req.Count = config.Telegram.ImageCount

	return channels.RunAction(req)
}

func aiImage(ctx context.Context, msg *memory.Message) error {
	for _, path := range msg.GetImages() {
		req := structs.NewActionRequest()
		req.Message = msg
		req.Context = ctx
		req.Action = "variation"
		req.Params["path"] = path
		req.Writer = NewTelegramWriter(ctx, req, false)
		req.Count = config.Telegram.ImageCount

		if err := channels.RunAction(req); err != nil {
			return err
		}
	}

	return nil
}
