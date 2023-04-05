package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs"
)

func (w *DiscordWriter) closeImage() error {
	logger := logging.GetLogger().With().Str("requestID", w.Request.GetID()).Logger()

	var files []*discordgo.File
	for k := range w.objs {
		obj := w.objs[k]
		fmt.Println(obj.Type, obj.Type == structs.Image)
		if obj.Type == structs.Image {
			files = append(files, &discordgo.File{
				Name:        fmt.Sprintf("generated-%d.png", k),
				ContentType: "image/png",
				Reader:      obj,
			})
		}
	}

	fmt.Println(files)

	if len(files) == 0 {
		return nil
	}

	logger.Debug().Str("chatID", w.ChatID).Msg("replying with image")

	content := fmt.Sprintf("Prompt: %s", w.Prompt)

	_, err := w.Bot.ChannelMessageSendComplex(w.ChatID, &discordgo.MessageSend{
		Files:   files,
		Content: content,
	})
	fmt.Println(err)
	fmt.Println("???")
	return err
}
