package cli

import (
	"context"
	"fmt"
	"github.com/briandowns/spinner"
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

	return openai.Complete(ctx, &structs.CompletionRequest{
		Context:         ctx,
		Writer:          os.Stdout,
		SkipClose:       true,
		StartChannel:    startCh,
		ContinueChannel: continueCh,
		User:            user,
		Mode:            "default",
		Content:         text,
	})
}
