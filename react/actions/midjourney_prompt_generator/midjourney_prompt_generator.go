package midjourney_prompt_generator

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	reactOpenAI "github.com/kamushadenes/chloe/react/openai"
	"github.com/kamushadenes/chloe/structs"
	"strings"
)

type MidjourneyPromptGeneratorAction struct {
	Name   string
	Params map[string]string
}

func (a *MidjourneyPromptGeneratorAction) GetNotification() string {
	return fmt.Sprintf("🖼️ Improving prompt: **%s**", a.Params["prompt"])
}

func (a *MidjourneyPromptGeneratorAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Text)

	resp, err := reactOpenAI.SimpleCompletionRequest(request.Context, "midjourney_prompt_generator", a.Params["prompt"])
	if err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	content := strings.TrimSpace(resp.Choices[0].Message.Content)

	if _, err := obj.Write([]byte(content)); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*structs.ResponseObject{obj}, nil
}
