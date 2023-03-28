package cli

import (
	"context"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/providers/openai"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/utils/colors"
	"os"
	"time"
)

func Complete(ctx context.Context, text string) error {
	s := spinner.New(spinner.CharSets[40], 100*time.Millisecond)
	s.Prefix = colors.BoldCyan("Assistant: ")
	s.Start()

	startCh := make(chan bool)
	continueCh := make(chan bool)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-startCh:
				s.Stop()
				fmt.Print(s.Prefix)
				continueCh <- true
				return
			}
		}
	}()

	msg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), "cli")
	msg.Role = "user"
	msg.User = user
	msg.SetContent(text)

	channels.IncomingMessagesCh <- msg
	if err := <-msg.ErrorCh; err != nil {
		return err
	}

	req := structs.NewCompletionRequest()
	req.Context = ctx
	req.Writer = os.Stdout
	req.SkipClose = true
	req.StartChannel = startCh
	req.ContinueChannel = continueCh
	req.Mode = "default"
	req.Message = msg

	return openai.Complete(req)
}
