package discord

import (
	"github.com/bwmarrin/discordgo"
)

func getCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "action",
			Description: "Perform an action",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "action",
					Description: "Action to execute",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "params",
					Description: "Parameters for the action",
					Required:    true,
				},
			},
		},
		{
			Name:        "forget",
			Description: "Wipe all context and reset the conversation with the bot",
		},
		{
			Name:        "generate",
			Description: "Generate an image using DALL-E",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "prompt",
					Description: "Prompt for the image",
					Required:    true,
				},
			},
		},
		{
			Name:        "tts",
			Description: "Converts text to speech",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "text",
					Description: "Text to convert",
					Required:    true,
				},
			},
		},
	}
}

func registerCommands(s *discordgo.Session) error {
	for _, command := range getCommands() {
		if _, err := s.ApplicationCommandCreate(s.State.User.ID, "", command); err != nil {
			return err
		}
	}

	return nil
}
