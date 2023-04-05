package memory

import (
	"context"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"gorm.io/gorm"
	"net/http"
)

type CLIMessageSource struct {
	PauseSpinnerCh  chan bool
	ResumeSpinnerCh chan bool
}

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

type SlackMessageSource struct {
	API          *slack.Client
	Message      *slackevents.MessageEvent
	SlashCommand slack.SlashCommand
}

type MessageSource struct {
	Telegram *TelegramMessageSource `json:"telegram,omitempty"`
	HTTP     *HTTPMessageSource     `json:"http,omitempty"`
	Discord  *DiscordMessageSource  `json:"discord,omitempty"`
	Slack    *SlackMessageSource    `json:"slack,omitempty"`
	CLI      *CLIMessageSource      `json:"cli,omitempty"`
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
	Context    context.Context    `json:"-" gorm:"-"`
	ExternalID string             `json:"externalId"`
	Interface  string             `json:"interface"`
	User       *User              `json:"user,omitempty"`
	UserID     uint               `json:"userId,omitempty"`
	Source     *MessageSource     `json:"source" gorm:"-"`
	Role       string             `json:"role"`
	Content    string             `json:"content"`
	Summary    string             `json:"summary"`
	Moderated  bool               `json:"moderated"`
	TokenCount int                `json:"tokenCount"`
	Moderation *MessageModeration `json:"moderation" gorm:"embedded;embeddedPrefix:moderation_"`
	ErrorCh    chan error         `json:"-" gorm:"-"`
	audioPaths []string
	imagePaths []string
}
