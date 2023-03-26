package react

import (
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/resources"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/utils"
	"github.com/sashabaranov/go-openai"
	"io"
	"strings"
)

type MidjourneyPromptGeneratorAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
}

func NewMidjourneyPromptGeneratorAction() Action {
	return &MidjourneyPromptGeneratorAction{
		Name: "image",
	}
}

func (a *MidjourneyPromptGeneratorAction) SetWriters(writers []io.WriteCloser) {
	a.Writers = writers
}

func (a *MidjourneyPromptGeneratorAction) GetWriters() []io.WriteCloser {
	return a.Writers
}
func (a *MidjourneyPromptGeneratorAction) GetName() string {
	return a.Name
}

func (a *MidjourneyPromptGeneratorAction) GetNotification() string {
	return fmt.Sprintf("🖼️ Improving prompt: **%s**", a.Params)
}

func (a *MidjourneyPromptGeneratorAction) SetParams(params string) {
	a.Params = params
}

func (a *MidjourneyPromptGeneratorAction) GetParams() string {
	return a.Params
}

func (a *MidjourneyPromptGeneratorAction) SetMessage(message *memory.Message) {}

func (a *MidjourneyPromptGeneratorAction) Execute(request *structs.ActionRequest) error {
	logger := logging.GetLogger()

	prompt, err := resources.GetPrompt("midjourney_prompt_generator", &resources.PromptArgs{
		Args: map[string]interface{}{},
		Mode: "midjourney_prompt_generator",
	})

	req := openai.ChatCompletionRequest{
		Model: config.OpenAI.DefaultModel.ChainOfThought,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: prompt,
			},
			{
				Role:    "user",
				Content: a.Params,
			},
		},
	}

	var resp openai.ChatCompletionResponse

	respi, err := utils.WaitTimeout(request.Context, config.Timeouts.ChainOfThought, func(ch chan interface{}, errCh chan error) {
		resp, err := openAIClient.CreateChatCompletion(request.Context, req)
		if err != nil {
			logger.Error().Err(err).Msg("error requesting prompt improvement")
			errCh <- err
		}
		ch <- resp
	})
	if err != nil {
		return err
	}

	resp = respi.(openai.ChatCompletionResponse)

	content := strings.TrimSpace(resp.Choices[0].Message.Content)

	for _, w := range a.Writers {
		if _, err := w.Write([]byte(content)); err != nil {
			return err
		}

	}

	return nil
}

func (a *MidjourneyPromptGeneratorAction) RunPreActions(request *structs.ActionRequest) error {
	return nil
}

func (a *MidjourneyPromptGeneratorAction) RunPostActions(request *structs.ActionRequest) error {
	return nil
}
