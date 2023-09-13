package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

func (w *DiscordWriter) closeImage() error {
	logger := logging.FromContext(w.Context)

	var files []*discordgo.File
	for k := range w.objs {
		obj := w.objs[k]
		if obj.Type == response_object_structs.Image {
			files = append(files, &discordgo.File{
				Name:        fmt.Sprintf("generated-%d.png", k),
				ContentType: "image/png",
				Reader:      obj,
			})
		}
	}

	if len(files) == 0 {
		return nil
	}

	logger.Debug().Str("chatID", w.ChatID).Msg("replying with image")

	content := fmt.Sprintf("Prompt: %s", w.Prompt)

	_, err := w.Bot.ChannelMessageSendComplex(w.ChatID, &discordgo.MessageSend{
		Files:   files,
		Content: content,
	})
	return err
}
