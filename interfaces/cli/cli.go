package cli

import (
	"bufio"
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/flags"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/utils/colors"
	"os"
	"strings"
)

var user *memory.User

func Handle(ctx context.Context) error {
	var err error

	user, err = memory.GetUserByExternalID(ctx, "cli", "cli")
	if err != nil {
		user, err = memory.NewUser(ctx, "User", "CLI", "cli")
		if err != nil {
			return err
		}
		err = user.AddExternalID(ctx, "cli", "cli")
		if err != nil {
			return err
		}
	}

	switch os.Args[1] {
	case "complete":
		if len(os.Args) > 2 {
			if err := Complete(ctx, strings.Join(os.Args[2:], " ")); err != nil {
				fmt.Println(err)
				return err
			}
		} else {
			flags.InteractiveCLI = true

			fmt.Println("Welcome to Chloe CLI")
			fmt.Println("Type 'quit' to exit")
			fmt.Println()

			scanner := bufio.NewScanner(os.Stdin)
			for {
				fmt.Print(colors.Bold("User: "))
				scanner.Scan()
				text := scanner.Text()

				if text == "quit" {
					break
				}
				fmt.Println()

				if err := Complete(ctx, text); err != nil {
					fmt.Println(err)
					return err
				}
				fmt.Println()
				fmt.Println()
			}
		}

	case "generate":
		if err := Generate(ctx, strings.Join(os.Args[2:], " ")); err != nil {
			fmt.Println(err)
			return err
		}
	case "forget":
		if err := Forget(ctx); err != nil {
			fmt.Println(err)
			return err
		}
	case "tts":
		if err := TTS(ctx, strings.Join(os.Args[2:], " ")); err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}
