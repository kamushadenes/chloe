package config

import "time"

type SlackConfig struct {
	Token                 string
	AppLevelToken         string
	ImageCount            int
	OnlyMention           bool
	StreamMessages        bool
	StreamFlushInterval   time.Duration
	SendProcessingMessage bool
	ProcessingMessage     string
	MaxMessageLength      int
}

var Slack = &SlackConfig{
	Token:                 envOrDefault("CHLOE_SLACK_TOKEN", ""),
	AppLevelToken:         envOrDefault("CHLOE_SLACK_APP_LEVEL_TOKEN", ""),
	ImageCount:            envOrDefaultIntInRange("CHLOE_SLACK_IMAGE_COUNT", 4, 1, 10),
	OnlyMention:           envOrDefaultBool("CHLOE_SLACK_ONLY_MENTION", true),
	StreamMessages:        envOrDefaultBool("CHLOE_SLACK_STREAM_MESSAGES", true),
	StreamFlushInterval:   envOrDefaultDuration("CHLOE_SLACK_STREAM_FLUSH_INTERVAL", 500*time.Millisecond),
	SendProcessingMessage: envOrDefaultBool("CHLOE_SLACK_SEND_PROCESSING_MESSAGE", false),
	ProcessingMessage:     envOrDefault("CHLOE_SLACK_PROCESSING_MESSAGE", "↻ Processing..."),
	MaxMessageLength:      envOrDefaultIntInRange("CHLOE_SLACK_MAX_MESSAGE_LENGTH", 40000, 1, 40000),
}
