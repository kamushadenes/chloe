package latex

import (
	"github.com/kamushadenes/chloe/interfaces/discord"
	"github.com/kamushadenes/chloe/interfaces/slack"
	"github.com/kamushadenes/chloe/interfaces/telegram"
	"github.com/kamushadenes/chloe/structs"
	"io"
)

func latexPreActions(a structs.Action, request *structs.ActionRequest) error {
	var ws []io.WriteCloser

	switch request.Message.Interface {
	case "telegram":
		w := request.GetWriters()[0].(*telegram.TelegramWriter)
		iw := w.ToImageWriter()
		siw := iw.Subwriter()
		siw.SetPrompt(request.Params)
		ws = append(ws, siw)
	case "discord":
		w := request.GetWriters()[0].(*discord.DiscordWriter)
		iw := w.ToImageWriter()
		siw := iw.Subwriter()
		siw.SetPrompt(request.Params)
		ws = append(ws, siw)
	case "slack":
		w := request.GetWriters()[0].(*slack.SlackWriter)
		iw := w.ToImageWriter()
		siw := iw.Subwriter()
		siw.SetPrompt(request.Params)
		ws = append(ws, siw)
	default:
		ws = append(ws, request.GetWriters()[0])
	}

	a.SetWriters(ws)

	return nil
}
