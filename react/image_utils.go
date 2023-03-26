package react

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/interfaces/discord"
	"github.com/kamushadenes/chloe/interfaces/telegram"
	"github.com/kamushadenes/chloe/structs"
	"io"
)

func imagePreActions(a Action, request *structs.ActionRequest) error {
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
	default:
		ws = append(ws, request.GetWriters()[0])
	}

	a.SetWriters(ws)

	return nil
}
