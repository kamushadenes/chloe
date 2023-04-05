package midjourney_prompt_generator

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	reactOpenAI "github.com/kamushadenes/chloe/react/openai"
	"github.com/kamushadenes/chloe/structs"
	"io"
	"strings"
)

type MidjourneyPromptGeneratorAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
}

func NewMidjourneyPromptGeneratorAction() structs.Action {
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
	resp, err := reactOpenAI.SimpleCompletionRequest(request.Context, "midjourney_prompt_generator", a.Params)
	if err != nil {
		return errors.Wrap(errors.ErrActionFailed, err)
	}

	content := strings.TrimSpace(resp.Choices[0].Message.Content)

	for _, w := range a.Writers {
		if _, err := w.Write([]byte(content)); err != nil {
			return errors.Wrap(errors.ErrActionFailed, err)
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
