package slack

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"strings"
)

func userFromMessage(ctx context.Context, msg *memory.Message) (*memory.User, error) {
	var user *memory.User
	var err error

	if msg.Source.Slack.Message != nil {
		user, err = memory.GetUserByExternalID(ctx, msg.Source.Slack.Message.User, "slack")
	} else {
		user, err = memory.GetUserByExternalID(ctx, msg.Source.Slack.SlashCommand.UserID, "slack")
	}

	if err != nil {
		var userInfo *slack.User
		userInfo, err = msg.Source.Slack.API.GetUserInfo(msg.Source.Slack.Message.User)
		if err != nil {
			return nil, err
		}

		names := strings.Fields(userInfo.Profile.DisplayNameNormalized)
		surname := ""
		if len(names) > 1 {
			surname = strings.Join(names[1:], " ")
		}

		user, err = memory.CreateUser(ctx, names[0], surname, msg.Source.Slack.Message.Username)
		if err != nil {
			return nil, err
		}
		err = user.AddExternalID(ctx, msg.Source.Slack.Message.User, "slack")
		if err != nil {
			return nil, err
		}
	}

	return user, err
}

func promptFromMessage(msg *memory.Message) string {
	if msg.Source.Slack.Message != nil {
		return msg.Source.Slack.Message.Text
	}
	return msg.Source.Slack.SlashCommand.Text
}

func isMentioned(auth *slack.AuthTestResponse, m *slackevents.MessageEvent) bool {
	return !config.Slack.OnlyMention || strings.Contains(m.Text, auth.UserID) || m.ChannelType == "im"
}
