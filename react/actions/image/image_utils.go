package image

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/react/actions/midjourney_prompt_generator"
	"github.com/kamushadenes/chloe/structs"
)

func imagePreActions(a structs.Action, request *structs.ActionRequest) error {

	if config.React.ImproveImagePrompts {
		na := midjourney_prompt_generator.NewMidjourneyPromptGeneratorAction()
		na.SetParams(a.GetParams())
		request.Message.NotifyAction(na.GetNotification())
		objs, err := na.Execute(request)
		if err == nil {
			a.SetParams(string(objs[0].Data))
		}
	}

	return nil
}
