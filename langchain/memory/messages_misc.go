package memory

import (
	"fmt"
	markdown "github.com/MichaelMure/go-term-markdown"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/colors"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/flags"
	"github.com/kamushadenes/chloe/i18n"
	"github.com/slack-go/slack"
	"golang.org/x/crypto/ssh/terminal"
	"os"
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
			if err != nil {
				return errors.Wrap(errors.ErrSendMessage, err)
			}
		}
		return nil
	case "discord":
		_, err := m.Source.Discord.API.ChannelMessageSend(m.Source.Discord.Message.ChannelID, text)
		if err != nil {
			return errors.Wrap(errors.ErrSendMessage, err)
		}

		return nil
	case "slack":
		var msgText = slack.MsgOptionText(text, false)

		if m.Source.Slack.Message != nil {
			_, _, err := m.Source.Slack.API.PostMessage(m.Source.Slack.Message.Channel, msgText)
			if err != nil {
				return errors.Wrap(errors.ErrSendMessage, err)
			}
		} else {
			_, _, err := m.Source.Slack.API.PostMessage(m.Source.Slack.SlashCommand.ChannelID, msgText)
			if err != nil {
				return errors.Wrap(errors.ErrSendMessage, err)
			}
		}

		return nil
	case "cli":
		if flags.InteractiveCLI {
			width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
			if err != nil {
				width = 80
			}

			result := markdown.Render(text, width, 0)

			m.Source.CLI.PauseSpinnerCh <- true

			fmt.Printf("%s %s\n", colors.BoldRed("Chloe:"), strings.Trim(strings.TrimSpace(string(result)), "*"))

			m.Source.CLI.ResumeSpinnerCh <- false
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

		if len(extraArgs) > 0 && extraArgs[0] != "notify_action" {
			msg.ReplyToMessageID = extraArgs[0].(int)
		}
		_, err := m.Source.Telegram.API.Send(msg)
		if err != nil {
			msg.ParseMode = ""
			_, err = m.Source.Telegram.API.Send(msg)
			if err != nil {
				return errors.Wrap(errors.ErrSendMessage, err)
			}
		}
		return nil
	case "discord":
		_, err := m.Source.Discord.API.ChannelMessageSend(m.Source.Discord.Message.ChannelID, text)
		if err != nil {
			return errors.Wrap(errors.ErrSendMessage, err)
		}

		return nil
	case "slack":
		var msgText = slack.MsgOptionText(text, false)

		if len(extraArgs) > 0 && extraArgs[0] == "notify_action" {
			msgText = slack.MsgOptionText(fmt.Sprintf("*%s*", strings.ReplaceAll(text, "*", "")), false)
		}

		if m.Source.Slack.Message != nil {
			_, _, err := m.Source.Slack.API.PostMessage(m.Source.Slack.Message.Channel, msgText)
			if err != nil {
				return errors.Wrap(errors.ErrSendMessage, err)
			}

			return nil
		} else {
			_, _, err := m.Source.Slack.API.PostMessage(m.Source.Slack.SlashCommand.ChannelID, msgText)
			if err != nil {
				return errors.Wrap(errors.ErrSendMessage, err)
			}

			return nil
		}
	case "cli":
		if flags.InteractiveCLI {
			width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
			if err != nil {
				width = 80
			}

			result := markdown.Render(text, width, 0)

			m.Source.CLI.PauseSpinnerCh <- true

			fmt.Printf("%s %s\n", colors.Yellow("Chloe:"), strings.Trim(strings.TrimSpace(string(result)), "*"))

			m.Source.CLI.ResumeSpinnerCh <- false
		}
	}

	return fmt.Errorf("unsupported interface %s", m.Interface)
}

func (m *Message) NotifyAction(act string) {
	_ = m.SendText(fmt.Sprintf("*%s*", act), false, "notify_action")
}
