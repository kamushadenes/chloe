package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kamushadenes/chloe/colors"
)

type CompleteCmd struct {
	Prompt []string `arg:"" optional:"" help:"Prompt to complete"`
	Model  string   `short:"m" long:"model" help:"Model to use for completion" default:"gpt-3.5-turbo"`
}

func (c *CompleteCmd) Run(globals *Globals) error {
	if len(c.Prompt) > 0 {
		return Complete(globals.Context, strings.Join(c.Prompt, " "), NewCLIWriter())
	}

	fmt.Println(colors.BoldGreen("Welcome to Chloe CLI"))
	fmt.Println(colors.BoldGreen("Type 'quit' to exit"))
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

		if err := Complete(globals.Context, text, NewCLIWriter()); err != nil {
			return err
		}
		fmt.Println()
		fmt.Println()
	}

	return nil
}
