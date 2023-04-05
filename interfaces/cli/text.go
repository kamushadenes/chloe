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
	"time"
)

func Complete(ctx context.Context, text string, writer structs.ChloeWriter) error {
	s := spinner.New(spinner.CharSets[40], 100*time.Millisecond)

	if flags.InteractiveCLI {
		s.Prefix = colors.BoldCyan("Chloe: ")
		s.Start()
	}

	msg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), "cli")
	msg.Role = "user"
	msg.User = user
	msg.Source = &memory.MessageSource{
		CLI: &memory.CLIMessageSource{
			PauseSpinnerCh:  make(chan bool),
			ResumeSpinnerCh: make(chan bool),
		},
	}

	msg.SetContent(text)

	if err := channels.RegisterIncomingMessage(msg); err != nil {
		return err
	}

	req := structs.NewCompletionRequest()
	req.Context = ctx
	req.Writer = writer
	req.SkipClose = true
	req.StartChannel = make(chan bool)
	req.ContinueChannel = make(chan bool)
	req.ErrorChannel = make(chan error)
	req.Mode = "default"
	req.Message = msg

	go func() {
		for {
			select {
			case <-msg.Source.CLI.PauseSpinnerCh:
				if flags.InteractiveCLI {
					s.Stop()
				}
			case <-msg.Source.CLI.ResumeSpinnerCh:
				if flags.InteractiveCLI {
					s.Start()
				}
			case <-ctx.Done():
				return
			case <-req.StartChannel:
				if flags.InteractiveCLI {
					s.Stop()
					fmt.Println()
					fmt.Print(s.Prefix)
				}
				req.ContinueChannel <- true
				return
			case err := <-req.ErrorChannel:
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

	return openai.Complete(req)
}
