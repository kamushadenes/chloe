package discord

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	"strings"
)

func userFromMessage(ctx context.Context, msg *memory.Message) (*memory.User, error) {
	user, err := memory.GetUserByExternalID(ctx, msg.Source.Discord.Message.Author.ID, "discord")

	if err != nil {
		names := strings.Fields(msg.Source.Discord.Message.Author.Username)
		surname := ""
		if len(names) > 1 {
			surname = strings.Join(names[1:], " ")
		}

		user, err = memory.CreateUser(ctx, names[0], surname, msg.Source.Discord.Message.Author.Username)
		if err != nil {
			return nil, err
		}
		err = user.AddExternalID(ctx, msg.Source.Discord.Message.Author.ID, "discord")
		if err != nil {
			return nil, err
		}
	}

	return user, err
}

func promptFromMessage(msg *memory.Message) string {
	if msg.Source.Discord.Interaction {
		return msg.Content
	}
	return msg.Source.Discord.Message.Content
}

func isMentioned(s *discordgo.Session, m *discordgo.Message) bool {
	if !config.Discord.OnlyMention || m.GuildID == "" {
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
