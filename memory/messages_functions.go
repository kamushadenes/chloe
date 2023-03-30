package memory

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/tokenizer"
	"github.com/sashabaranov/go-openai"
)

func LoadNonSummarizedMessages(ctx context.Context) ([]*Message, error) {
	var messages []*Message

	if err := db.WithContext(ctx).Where("(summary IS NULL OR summary == '') AND " +
		"(content IS NOT NULL AND content != '')").Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func LoadNonModeratedMessages(ctx context.Context) ([]*Message, error) {
	var messages []*Message

	if err := db.WithContext(ctx).Where("moderated = ? AND "+
		"content IS NOT NULL AND "+
		"content != ''", false).
		Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func MessagesFromOpenAIChatCompletionResponse(user *User, interf string, resp *openai.ChatCompletionResponse) []*Message {
	var messages []*Message

	for k := range resp.Choices {
		msg := NewMessage(uuid.Must(uuid.NewV4()).String(), interf)

		msg.SetContent(resp.Choices[k].Message.Content)
		msg.Role = resp.Choices[k].Message.Role
		msg.User = user

		messages = append(messages, msg)
	}

	return messages
}

func NewMessage(externalID string, interf string) *Message {
	return &Message{
		ExternalID: externalID,
		Interface:  interf,
		Source:     &MessageSource{},
		ErrorCh:    make(chan error),
	}
}

func (m *Message) Copy() *Message {
	msg := NewMessage(m.ExternalID, m.Interface)
	msg.User = m.User
	msg.Role = m.Role
	msg.Content = m.Content
	msg.TokenCount = m.TokenCount
	msg.Summary = m.Summary
	msg.Moderated = m.Moderated
	msg.Moderation = m.Moderation
	msg.Source = m.Source

	return msg
}

func (m *Message) GetExternalMessageID() string {
	if m.Source.Telegram != nil {
		return fmt.Sprintf("%d", m.Source.Telegram.Update.Message.MessageID)
	}

	if m.Source.HTTP != nil {
		return m.Source.HTTP.Request.Header.Get("X-Request-ID")
	}

	return ""
}

func (m *Message) GetTexts() []string {
	var txts []string

	if m.Source.Telegram != nil {
		txts = append(txts, m.Source.Telegram.Update.Message.Text)
	}

	if m.Source.Discord != nil {
		txts = append(txts, m.Source.Discord.Message.Content)
	}

	return txts
}

func (m *Message) AddAudio(path string) {
	m.audioPaths = append(m.audioPaths, path)
}

func (m *Message) AddImage(path string) {
	m.imagePaths = append(m.imagePaths, path)
}

func (m *Message) GetImages() []string {
	return m.imagePaths
}

func (m *Message) GetAudios() []string {
	return m.audioPaths
}

func (m *Message) SetSummary(ctx context.Context, summary string) error {
	return db.WithContext(ctx).
		Model(m).
		Update("summary", summary).Error
}

func (m *Message) Save(ctx context.Context) error {
	if m.Content == "" {
		return nil
	}
	return db.WithContext(ctx).
		Save(m).Error
}

func (m *Message) GetContent() string {
	if len(m.Summary) < len(m.Content) && m.Summary != "" {
		return m.Summary
	}

	return m.Content
}

func (m *Message) SetContent(content string) {
	m.Content = content

	m.TokenCount = tokenizer.CountTokens(config.OpenAI.DefaultModel.Completion.String(), content)
}

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
