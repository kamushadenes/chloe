package memory

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/tokenizer"
	"github.com/sashabaranov/go-openai"
)

func LoadNonSummarizedMessages(ctx context.Context) ([]*Message, error) {
	var messages []*Message

	if err := db.WithContext(ctx).Where("(summary IS NULL OR summary == '') AND " +
		"(content IS NOT NULL AND content != '')").Find(&messages).Error; err != nil {
		return nil, errors.Wrap(errors.ErrLoadMessages, err)
	}

	return messages, nil
}

func LoadNonModeratedMessages(ctx context.Context) ([]*Message, error) {
	var messages []*Message

	if err := db.WithContext(ctx).Where("moderated = ? AND "+
		"content IS NOT NULL AND "+
		"content != ''", false).
		Find(&messages).Error; err != nil {
		return nil, errors.Wrap(errors.ErrLoadMessages, err)
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
	err := db.WithContext(ctx).
		Model(m).
		Update("summary", summary).Error

	if err != nil {
		return errors.Wrap(errors.ErrUpdateMessage, err)
	}

	return nil
}

func (m *Message) Save(ctx context.Context) error {
	if m.Content == "" {
		return nil
	}
	err := db.WithContext(ctx).
		Save(m).Error

	if err != nil {
		return errors.Wrap(errors.ErrSaveMessage, err)
	}

	return nil
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

func BulkChangeMessageOwner(ctx context.Context, oldUser *User, newUser *User) error {
	err := db.WithContext(ctx).
		Model(&Message{}).
		Where("user_id = ?", oldUser.ID).
		Update("user_id", newUser.ID).Error

	if err != nil {
		return errors.Wrap(errors.ErrUpdateMessage, err)
	}

	return nil
}
