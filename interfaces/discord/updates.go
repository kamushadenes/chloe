package discord

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/i18n"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/memory"
	"strings"
)

func isMentioned(s *discordgo.Session, m *discordgo.Message) bool {
	if !config.Discord.OnlyMention {
		return true
	}

	if m.GuildID == "" {
		return true
	}

	for k := range m.Mentions {
		user := m.Mentions[k]
		if user.ID == s.State.User.ID {
			return true
		}
	}

	return false
}

func handleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !isMentioned(s, m.Message) {
		return
	}

	ctx := context.Background()
	logger := logging.GetLogger().With().Str("interface", "discord").Logger()

	msg := memory.NewMessage(m.Message.ID, "discord")
	msg.Source.Discord = &memory.DiscordMessageSource{
		API:         s,
		Message:     m.Message,
		Interaction: false,
	}
	msg.Content = strings.TrimSpace(strings.TrimLeft(m.Content, fmt.Sprintf("<@!%s>", s.State.User.ID)))
	msg.Role = "user"

	user, err := userFromMessage(ctx, msg)
	if err != nil {
		logger.Error().Err(err).Msg("error getting user from message")
		return
	}
	msg.User = user

	_ = msg.Save(ctx)

	channels.IncomingMessagesCh <- msg

	complete(ctx, msg)
}

func handleCommandInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	ctx := context.Background()
	logger := logging.GetLogger().With().Str("interface", "discord").Logger()

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

	switch i.ApplicationCommandData().Name {
	case "generate":
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		prompt := optionMap["prompt"].StringValue()

		msg.Content = prompt

		_ = msg.Save(ctx)

		generate(ctx, msg)

		reply = i18n.GetImageGenerationText()
	case "tts":
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		text := optionMap["text"].StringValue()

		msg.Content = text

		_ = msg.Save(ctx)

		tts(ctx, msg)

		reply = i18n.GetTextToSpeechText()
	case "forget":
		_ = forgetUser(ctx, msg)
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
