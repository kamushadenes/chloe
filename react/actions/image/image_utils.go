package image

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/interfaces/discord"
	"github.com/kamushadenes/chloe/interfaces/slack"
	"github.com/kamushadenes/chloe/interfaces/telegram"
	"github.com/kamushadenes/chloe/react/actions/midjourney_prompt_generator"
	structs2 "github.com/kamushadenes/chloe/react/actions/structs"
	utils2 "github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/structs"
	"io"
)

func imagePreActions(a structs2.Action, request *structs.ActionRequest) error {
	var ws []io.WriteCloser

	switch request.Message.Interface {
	case "telegram":
		w := request.GetWriters()[0].(*telegram.TelegramWriter)
		iw := w.ToImageWriter()
		for k := 0; k < config.Telegram.ImageCount; k++ {
			siw := iw.Subwriter()
			siw.SetPrompt(request.Params)
			ws = append(ws, siw)
		}
	case "discord":
		w := request.GetWriters()[0].(*discord.DiscordWriter)
		iw := w.ToImageWriter()
		for k := 0; k < config.Discord.ImageCount; k++ {
			siw := iw.Subwriter()
			siw.SetPrompt(request.Params)
			ws = append(ws, siw)
		}
	case "slack":
		w := request.GetWriters()[0].(*slack.SlackWriter)
		iw := w.ToImageWriter()
		for k := 0; k < config.Slack.ImageCount; k++ {
			siw := iw.Subwriter()
			siw.SetPrompt(request.Params)
			ws = append(ws, siw)
		}
	default:
		ws = append(ws, request.GetWriters()[0])
	}

	a.SetWriters(ws)

	if config.React.ImproveImagePrompts {
		b := &utils2.BytesWriter{}

		na := midjourney_prompt_generator.NewMidjourneyPromptGeneratorAction()
		na.SetParams(a.GetParams())
		na.SetWriters([]io.WriteCloser{b})
		request.Message.NotifyAction(na.GetNotification())
		if err := na.Execute(request); err == nil {
			a.SetParams(string(b.Bytes))
		}
	}

	return nil
}
