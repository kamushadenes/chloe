package cli

import (
	"bufio"
	"fmt"
	"github.com/kamushadenes/chloe/utils/colors"
	"os"
	"strings"
)

type CompleteCmd struct {
	Prompt []string `arg:"" optional:"" help:"Prompt to complete"`
}

func (c *CompleteCmd) Run(globals *Globals) error {
	if len(c.Prompt) > 0 {
		return Complete(globals.Context, strings.Join(c.Prompt, " "))
	} else {
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

			if err := Complete(globals.Context, text); err != nil {
				return err
			}
			fmt.Println()
			fmt.Println()
		}
	}

	return nil
}
