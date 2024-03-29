package discord

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/kamushadenes/chloe/i18n"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/logging"
)

func handleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !isMentioned(s, m.Message) {
		return
	}

	logger := logging.GetLogger().With().Str("interface", "discord").Str("requestID", m.Message.ID).Logger()
	ctx := logger.WithContext(context.Background())

	msg := memory.NewMessage(m.Message.ID, "discord")
	msg.Source.Discord = &memory.DiscordMessageSource{
		API:         s,
		Message:     m.Message,
		Interaction: false,
	}
	msg.SetContent(strings.TrimSpace(strings.TrimLeft(m.Content, fmt.Sprintf("<@!%s>", s.State.User.ID))))
	msg.Role = "user"

	user, err := userFromMessage(ctx, msg)
	if err != nil {
		logger.Error().Err(err).Msg("error getting user from message")
		_ = msg.SendError(err)
		return
	}
	msg.User = user

	if err := msg.Save(ctx); err != nil {
		_ = msg.SendError(err)
		return
	}

	_ = msg.Source.Discord.API.ChannelTyping(msg.Source.Discord.Message.ChannelID)

	if err := complete(ctx, msg); err != nil {
		_ = msg.SendError(err)
		return
	}
}

func handleCommandInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Message == nil {
		i.Message = &discordgo.Message{
			ID:        i.ID,
			ChannelID: i.ChannelID,
			Author:    i.User,
		}
		if i.Message.Author == nil {
			i.Message.Author = i.Member.User
		}
	}

	logger := logging.GetLogger().With().Str("interface", "discord").Str("requestID", i.Message.ID).Logger()
	ctx := logger.WithContext(context.Background())

	msg := memory.NewMessage(i.Message.ID, "discord")
	msg.Source.Discord = &memory.DiscordMessageSource{
		API:         s,
		Message:     i.Message,
		Interaction: true,
	}
	msg.Role = "user"

	user, err := userFromMessage(ctx, msg)
	if err != nil {
		logger.Error().Err(err).Msg("error getting user from message")
		return
	}
	msg.User = user

	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))

	var reply string

	_ = msg.Source.Discord.API.ChannelTyping(msg.Source.Discord.Message.ChannelID)

	switch i.ApplicationCommandData().Name {
	case "action":
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		act := optionMap["action"].StringValue()
		params := optionMap["params"].StringValue()

		msg.SetContent(fmt.Sprintf("%s %s", act, params))

		if err := msg.Save(ctx); err != nil {
			_ = msg.SendError(err)
			return
		}

		go func() {
			if err := action(ctx, msg); err != nil {
				_ = msg.SendError(err)
				return
			}
		}()

		reply = "Running action..."
	case "generate":
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		prompt := optionMap["prompt"].StringValue()

		msg.SetContent(prompt)

		if err := msg.Save(ctx); err != nil {
			_ = msg.SendError(err)
			return
		}

		go func() {
			if err := generate(ctx, msg); err != nil {
				_ = msg.SendError(err)
				return
			}
		}()

		reply = i18n.GetImageGenerationText()
	case "tts":
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		text := optionMap["text"].StringValue()

		msg.SetContent(text)

		if err := msg.Save(ctx); err != nil {
			_ = msg.SendError(err)
			return
		}

		go func() {
			if err := tts(ctx, msg); err != nil {
				_ = msg.SendError(err)
				return
			}
		}()

		reply = i18n.GetTextToSpeechText()
	case "forget":
		if err := forgetUser(ctx, msg); err != nil {
			_ = msg.SendError(err)
			return
		}
		reply = i18n.GetForgetText()
	}

	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: reply,
		},
	}); err != nil {
		logger.Error().Err(err).Msg("error responding to interaction")
	}
}
