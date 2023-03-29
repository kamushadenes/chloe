package cli

import (
	"bufio"
	"context"
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/kamushadenes/chloe/flags"
	"github.com/kamushadenes/chloe/tokenizer"
	"github.com/kamushadenes/chloe/utils/colors"
	"os"
	"strings"
)

type Globals struct {
	Context context.Context
}

var CLIFlags struct {
	Complete    CompleteCmd    `cmd:"complete" short:"c" help:"Complete a prompt"`
	Generate    GenerateCmd    `cmd:"generate" short:"g" help:"Generate an prompt"`
	TTS         TTSCmd         `cmd:"tts" short:"t" help:"Generate an audio from a prompt"`
	Forget      ForgetCmd      `cmd:"forget" short:"f" help:"Forget all users"`
	CountTokens CountTokensCmd `cmd:"count-tokens" help:"Count tokens"`
	Version     VersionFlag    `name:"version" help:"Print version information and quit"`
}

func parseFlags() *kong.Context {
	return kong.Parse(&CLIFlags)
}

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

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

type GenerateCmd struct {
	Prompt     []string `arg:"" help:"Prompt to generate"`
	OutputPath string   `help:"Output path, if not specified, it will be printed to stdout if not a TTY, otherwise it will be saved to generated.png"`
}

func (c *GenerateCmd) Run(globals *Globals) error {
	if len(c.OutputPath) > 0 {
		f, err := os.Create(c.OutputPath)
		if err != nil {
			return err
		}

		return Generate(globals.Context, strings.Join(c.Prompt, " "), f)
	}

	if flags.InteractiveCLI {
		f, err := os.Create("generated.png")
		if err != nil {
			return err
		}
		return Generate(globals.Context, strings.Join(c.Prompt, " "), f)
	}
	return Generate(globals.Context, strings.Join(c.Prompt, " "))
}

type TTSCmd struct {
	Prompt     []string `arg:"" help:"Prompt to generate"`
	OutputPath string   `help:"Output path, if not specified, it will be printed to stdout if not a TTY, otherwise it will be saved to generated.mp3"`
}

func (c *TTSCmd) Run(globals *Globals) error {
	if len(c.OutputPath) > 0 {
		f, err := os.Create(c.OutputPath)
		if err != nil {
			return err
		}

		return TTS(globals.Context, strings.Join(c.Prompt, " "), f)
	}

	if flags.InteractiveCLI {
		f, err := os.Create("generated.mp3")
		if err != nil {
			return err
		}
		return TTS(globals.Context, strings.Join(c.Prompt, " "), f)
	}
	return TTS(globals.Context, strings.Join(c.Prompt, " "))
}

type ForgetCmd struct {
	All bool `help:"Forget all users, not just the CLI user"`
}

func (c *ForgetCmd) Run(globals *Globals) error {
	return Forget(globals.Context, c.All)
}

type CountTokensCmd struct {
	Prompt []string `arg:"" help:"Prompt to generate"`
	Model  string   `help:"Model to use" default:"gpt-3.5-turbo"`
}

func (c *CountTokensCmd) Run(globals *Globals) error {
	fmt.Println(tokenizer.CountTokens(c.Model, strings.Join(c.Prompt, " ")))

	return nil
}
