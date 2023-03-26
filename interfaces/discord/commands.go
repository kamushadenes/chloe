package discord

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/kamushadenes/chloe/memory"
	"github.com/rs/zerolog"
)

func getCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
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
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command)
		if err != nil {
			return err
		}
	}

	return nil
}

func complete(ctx context.Context, msg *memory.Message) {
	logger := zerolog.Ctx(ctx)

	_ = msg.Source.Discord.API.ChannelTyping(msg.Source.Discord.Message.ChannelID)

	if err := aiComplete(ctx, msg); err != nil {
		logger.Error().Err(err).Msg("error generating image")
	}
}

func generate(ctx context.Context, msg *memory.Message) {
	logger := zerolog.Ctx(ctx)

	_ = msg.Source.Discord.API.ChannelTyping(msg.Source.Discord.Message.ChannelID)

	if err := aiGenerate(ctx, msg); err != nil {
		logger.Error().Err(err).Msg("error generating image")
	}
}

func tts(ctx context.Context, msg *memory.Message) {
	logger := zerolog.Ctx(ctx)

	_ = msg.Source.Discord.API.ChannelTyping(msg.Source.Discord.Message.ChannelID)

	if err := aiTTS(ctx, msg); err != nil {
		logger.Error().Err(err).Msg("error generating audio")
	}
}

func forgetUser(ctx context.Context, msg *memory.Message) error {
	return msg.User.DeleteMessages(ctx)
}
