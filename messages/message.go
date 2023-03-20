package messages

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/users"
	"net/http"
)

type TelegramMessageSource struct {
	API    *tgbotapi.BotAPI
	Update tgbotapi.Update
}

type HTTPMessageSource struct {
	Request *http.Request
}

type MessageSource struct {
	Telegram *TelegramMessageSource `json:"telegram,omitempty"`
	HTTP     *HTTPMessageSource     `json:"http,omitempty"`
}

type Message struct {
	ExternalID string         `json:"externalId"`
	User       *users.User    `json:"user,omitempty"`
	Source     *MessageSource `json:"source"`
	audioPaths []string
	imagePaths []string
}

func NewMessage(externalId string) *Message {
	return &Message{
		ExternalID: externalId,
		Source:     &MessageSource{},
	}
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

	if m.Source.HTTP != nil {
		//return m.Source.HTTP.Request.GetBody()
	}

	if m.Source.Telegram != nil {
		txts = append(txts, m.Source.Telegram.Update.Message.Text)
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
