package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rs/zerolog"
)

func (w *DiscordWriter) closeAudio() error {
	logger := zerolog.Ctx(w.Context).With().Str("requestID", w.Request.GetID()).Logger()

	logger.Debug().Str("chatID", w.ChatID).Msg("replying with audio")

	var files []*discordgo.File

	for k := range w.objs {
		obj := w.objs[k]

		if obj.Type == structs.Audio {
			files = append(files, &discordgo.File{
				Name:        fmt.Sprintf("generated-%d.mp3", k),
				ContentType: "audio/mpeg",
				Reader:      obj,
			})
		}
	}
	if len(files) == 0 {
		return nil
	}

	_, err := w.Bot.ChannelMessageSendComplex(w.ChatID, &discordgo.MessageSend{
		Files:   files,
		Content: fmt.Sprintf("Prompt: %s", w.Prompt),
	})
	return err
}
