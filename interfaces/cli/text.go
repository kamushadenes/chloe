package cli

import (
	"context"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/flags"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/providers/openai"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/utils/colors"
	"os"
	"time"
)

func Complete(ctx context.Context, text string) error {
	s := spinner.New(spinner.CharSets[40], 100*time.Millisecond)
	if flags.InteractiveCLI {
		s.Prefix = colors.BoldCyan("Chloe: ")
		s.Start()
		s.Disable()
	}

	startCh := make(chan bool)
	continueCh := make(chan bool)
	errorCh := make(chan error)

	pauseSpinnerCh := make(chan bool)
	resumeSpinnerCh := make(chan bool)

	go func() {
		for {
			select {
			case <-pauseSpinnerCh:
				if flags.InteractiveCLI {
					s.Stop()
				}
			case <-resumeSpinnerCh:
				if flags.InteractiveCLI {
					s.Start()
				}
			case <-ctx.Done():
				return
			case <-startCh:
				if flags.InteractiveCLI {
					s.Stop()
					fmt.Println()
					fmt.Print(s.Prefix)
				}
				continueCh <- true
				return
			case err := <-errorCh:
				if flags.InteractiveCLI {
					s.Stop()
					fmt.Println()
					fmt.Print(s.Prefix)
				}
				fmt.Println(colors.BoldRed(err.Error()))
				return
			}
		}
	}()

	msg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), "cli")
	msg.Role = "user"
	msg.User = user
	msg.Source = &memory.MessageSource{
		CLI: &memory.CLIMessageSource{
			PauseSpinnerCh:  pauseSpinnerCh,
			ResumeSpinnerCh: resumeSpinnerCh,
		},
	}

	msg.SetContent(text)

	if err := channels.RegisterIncomingMessage(msg); err != nil {
		return err
	}

	req := structs.NewCompletionRequest()
	req.Context = ctx
	req.Writer = os.Stdout
	req.SkipClose = true
	req.StartChannel = startCh
	req.ContinueChannel = continueCh
	req.ErrorChannel = errorCh
	req.Mode = "default"
	req.Message = msg

	return openai.Complete(req)
}
