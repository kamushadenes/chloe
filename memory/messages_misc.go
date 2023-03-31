package memory

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/i18n"
	"github.com/slack-go/slack"
	"strings"
)

// SendError TODO: this is in the wrong place
func (m *Message) SendError(err error) error {
	text := i18n.GetErrorText(err)

	switch m.Interface {
	case "telegram":
		msg := tgbotapi.NewMessage(m.Source.Telegram.Update.Message.Chat.ID, text)
		_, err := m.Source.Telegram.API.Send(msg)
		if err != nil {
			msg.ParseMode = ""
			_, err = m.Source.Telegram.API.Send(msg)
		}
		return err
	case "discord":
		_, err := m.Source.Discord.API.ChannelMessageSend(m.Source.Discord.Message.ChannelID, text)
		return err
	case "slack":
		var msgText = slack.MsgOptionText(text, false)

		if m.Source.Slack.Message != nil {
			_, _, err := m.Source.Slack.API.PostMessage(m.Source.Slack.Message.Channel, msgText)
			return err
		} else {
			_, _, err := m.Source.Slack.API.PostMessage(m.Source.Slack.SlashCommand.ChannelID, msgText)
			return err
		}
	}

	return nil
}

// SendText TODO: this is in the wrong place and extraArgs is a mess
func (m *Message) SendText(text string, notify bool, extraArgs ...interface{}) error {
	if len(text) == 0 {
		return nil
	}
	switch m.Interface {
	case "telegram":
		msg := tgbotapi.NewMessage(m.Source.Telegram.Update.Message.Chat.ID, text)
		msg.ParseMode = tgbotapi.ModeMarkdown

		if !notify {
			msg.DisableNotification = true
			msg.DisableWebPagePreview = true
		}

		if len(extraArgs) > 0 {
			msg.ReplyToMessageID = extraArgs[0].(int)
		}
		_, err := m.Source.Telegram.API.Send(msg)
		if err != nil {
			msg.ParseMode = ""
			_, err = m.Source.Telegram.API.Send(msg)
		}
		return err
	case "discord":
		_, err := m.Source.Discord.API.ChannelMessageSend(m.Source.Discord.Message.ChannelID, text)
		return err
	case "slack":
		var msgText = slack.MsgOptionText(text, false)

		if len(extraArgs) > 0 && extraArgs[0] == "notify_action" {
			msgText = slack.MsgOptionText(fmt.Sprintf("*%s*", strings.ReplaceAll(text, "*", "")), false)
		}

		if m.Source.Slack.Message != nil {
			_, _, err := m.Source.Slack.API.PostMessage(m.Source.Slack.Message.Channel, msgText)
			return err
		} else {
			_, _, err := m.Source.Slack.API.PostMessage(m.Source.Slack.SlashCommand.ChannelID, msgText)
			return err
		}
	}

	return fmt.Errorf("unsupported interface %s", m.Interface)
}

func (m *Message) NotifyAction(act string) {
	_ = m.SendText(fmt.Sprintf("*%s*", act), false, "notify_action")
}
