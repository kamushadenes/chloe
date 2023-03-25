package memory

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"net/http"
)

type TelegramMessageSource struct {
	API    *tgbotapi.BotAPI
	Update tgbotapi.Update
}

type HTTPMessageSource struct {
	Request *http.Request
}

type DiscordMessageSource struct {
	API         *discordgo.Session
	Message     *discordgo.Message
	Interaction bool
}

type MessageSource struct {
	Telegram *TelegramMessageSource `json:"telegram,omitempty"`
	HTTP     *HTTPMessageSource     `json:"http,omitempty"`
	Discord  *DiscordMessageSource  `json:"discord,omitempty"`
}

type MessageModeration struct {
	CategoryHate            bool `json:"categoryHate"`
	CategoryHateThreatening bool `json:"categoryHateThreatening"`
	CategorySelfHarm        bool `json:"categorySelfHarm"`
	CategorySexual          bool `json:"categorySexual"`
	CategorySexualMinors    bool `json:"categorySexualMinors"`
	CategoryViolence        bool `json:"categoryViolence"`
	CategoryViolenceGraphic bool `json:"categoryViolenceGraphic"`

	CategoryScoreHate            float32 `json:"categoryScoreHate"`
	CategoryScoreHateThreatening float32 `json:"categoryScoreHateThreatening"`
	CategoryScoreSelfHarm        float32 `json:"categoryScoreSelfHarm"`
	CategoryScoreSexual          float32 `json:"categoryScoreSexual"`
	CategoryScoreSexualMinors    float32 `json:"categoryScoreSexualMinors"`
	CategoryScoreViolence        float32 `json:"categoryScoreViolence"`
	CategoryScoreViolenceGraphic float32 `json:"categoryScoreViolenceGraphic"`

	Flagged bool `json:"flagged"`
}

type Message struct {
	gorm.Model
	ExternalID     string             `json:"externalId"`
	Interface      string             `json:"interface"`
	User           *User              `json:"user,omitempty"`
	UserID         uint               `json:"userId,omitempty"`
	Source         *MessageSource     `json:"source" gorm:"-"`
	Role           string             `json:"role"`
	Content        string             `json:"content"`
	Summary        string             `json:"summary"`
	ChainOfThought string             `json:"chainOfThought"`
	Moderated      bool               `json:"moderated"`
	Moderation     *MessageModeration `json:"moderation" gorm:"embedded;embeddedPrefix:moderation_"`
	audioPaths     []string
	imagePaths     []string
}

func NewMessage(externalID string, interf string) *Message {
	return &Message{
		ExternalID: externalID,
		Interface:  interf,
		Source:     &MessageSource{},
	}
}

func (m *Message) Copy() *Message {
	msg := NewMessage(m.ExternalID, m.Interface)
	msg.User = m.User
	msg.Role = m.Role
	msg.Content = m.Content
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

func (m *Message) SetSummary(ctx context.Context, summary string) error {
	return db.WithContext(ctx).Model(m).Update("summary", summary).Error
}

func (m *Message) Save(ctx context.Context) error {
	return db.WithContext(ctx).Save(m).Error
}

func (m *Message) GetContent(chainOfThought bool) string {
	if chainOfThought {
		if m.ChainOfThought != "" {
			return m.ChainOfThought
		}
	}

	if len(m.Summary) < len(m.Content) && m.Summary != "" {
		return m.Summary
	}

	return m.Content
}

func (m *Message) SendText(text string, replyTo ...interface{}) error {
	if len(text) == 0 {
		return nil
	}
	switch m.Interface {
	case "telegram":
		msg := tgbotapi.NewMessage(m.Source.Telegram.Update.Message.Chat.ID, text)
		msg.ParseMode = tgbotapi.ModeMarkdown
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
