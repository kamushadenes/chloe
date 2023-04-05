package midjourney_prompt_generator

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	reactOpenAI "github.com/kamushadenes/chloe/react/openai"
	"github.com/kamushadenes/chloe/structs"
	"strings"
)

type MidjourneyPromptGeneratorAction struct {
	Name   string
	Params string
}

func NewMidjourneyPromptGeneratorAction() structs.Action {
	return &MidjourneyPromptGeneratorAction{
		Name: "image",
	}
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

func (a *MidjourneyPromptGeneratorAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Text)

	resp, err := reactOpenAI.SimpleCompletionRequest(request.Context, "midjourney_prompt_generator", a.Params)
	if err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	content := strings.TrimSpace(resp.Choices[0].Message.Content)

	if _, err := obj.Write([]byte(content)); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*structs.ResponseObject{obj}, nil
}

func (a *MidjourneyPromptGeneratorAction) RunPreActions(request *structs.ActionRequest) error {
	return nil
}

func (a *MidjourneyPromptGeneratorAction) RunPostActions(request *structs.ActionRequest) error {
	return nil
}
