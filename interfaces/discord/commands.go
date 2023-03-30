package discord

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rs/zerolog"
	"io"
	"strings"
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
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command)
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO: This is a mess, refactor it

func action(ctx context.Context, msg *memory.Message) {
	fields := strings.Fields(msg.Content)

	req := structs.NewActionRequest()
	req.Context = ctx
	req.Message = msg
	req.Action = fields[0]
	req.Params = strings.Join(fields[1:], " ")
	req.Thought = fmt.Sprintf("User wants to run action %s", fields[0])
	req.Writers = []io.WriteCloser{NewTextWriter(ctx, req, false)}

	channels.ActionRequestsCh <- req
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
