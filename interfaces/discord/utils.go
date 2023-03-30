package discord

import (
	"context"
	"github.com/kamushadenes/chloe/memory"
	"strings"
)

func userFromMessage(ctx context.Context, msg *memory.Message) (*memory.User, error) {
	user, err := memory.GetUserByExternalID(ctx, msg.Source.Discord.Message.Author.ID, "discord")

	names := strings.Fields(msg.Source.Discord.Message.Author.Username)
	surname := ""
	if len(names) > 1 {
		surname = strings.Join(names[1:], " ")
	}

	if err != nil {
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
