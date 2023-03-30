package memory

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// SendText TODO: this is in the wrong place
func (m *Message) SendText(text string, notify bool, replyTo ...interface{}) error {
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

		if len(replyTo) > 0 {
			msg.ReplyToMessageID = replyTo[0].(int)
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
	}

	return fmt.Errorf("unsupported interface %s", m.Interface)
}

func (m *Message) NotifyAction(act string) {
	_ = m.SendText(fmt.Sprintf("*%s*", act), false)
}
