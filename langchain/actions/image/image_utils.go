package image

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/structs/action_structs"
)

func imagePreActions(a action_structs.Action, request *action_structs.ActionRequest) error {
	if config.React.ImproveImagePrompts {
		/*
			na := midjourney_prompt_generator.NewMidjourneyPromptGeneratorAction()

			p, err := a.GetParam("prompt")
			if err != nil {
				return err
			}

			na.SetParam("prompt", p)
			request.Message.NotifyAction(na.GetNotification())
			objs, err := na.Execute(request)
			if err == nil {
				a.SetParam("prompt", string(objs[0].Data))
			}
		*/
	}

	return nil
}
