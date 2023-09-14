package memory

import (
	"context"
	"fmt"

	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/tokenizer"
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
		"content != '' AND interface != 'internal'", false).
		Find(&messages).Error; err != nil {
		return nil, errors.Wrap(errors.ErrLoadMessages, err)
	}

	return messages, nil
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
	if path != "" {
		m.audioPaths = append(m.audioPaths, path)
	}
}

func (m *Message) AddImage(path string) {
	if path != "" {
		m.imagePaths = append(m.imagePaths, path)
	}
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
	if m.Content == "" && m.FunctionCallName == "" {
		return nil
	}

	err := db.WithContext(ctx).
		Save(m).Error

	if err != nil {
		return errors.Wrap(errors.ErrSaveMessage, err)
	}

	logger := logging.FromContext(ctx).Info()

	if m.User != nil {
		logger = logger.Uint("userID", m.User.ID).
			Str("username", m.User.Username).
			Str("firstName", m.User.FirstName).
			Str("lastName", m.User.LastName)
	}

	logger = logger.
		Str("interface", m.Interface).
		Str("content", m.Content)

	if m.Interface != "internal" {
		logger.Msg("message received")
	} else {
		logger.Msg("internal message")
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

	m.TokenCount = tokenizer.CountTokens("gpt-3.5-turbo", content)
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
