package config

import "time"

type DiscordConfig struct {
	Token                      string
	ImageCount                 int
	OnlyMention                bool
	RandomStatusUpdateInterval time.Duration
	StreamMessages             bool
	StreamFlushInterval        time.Duration
	SendProcessingMessage      bool
	ProcessingMessage          string
	MaxMessageLength           int
}

var Discord = &DiscordConfig{
	Token:                      envOrDefault("CHLOE_DISCORD_TOKEN", ""),
	ImageCount:                 envOrDefaultIntInRange("CHLOE_DISCORD_IMAGE_COUNT", 4, 1, 10),
	OnlyMention:                envOrDefaultBool("CHLOE_DISCORD_ONLY_MENTION", true),
	RandomStatusUpdateInterval: envOrDefaultDuration("CHLOE_DISCORD_RANDOM_STATUS_UPDATE_INTERVAL", 20*time.Second),
	StreamMessages:             envOrDefaultBool("CHLOE_DISCORD_STREAM_MESSAGES", false),
	StreamFlushInterval:        envOrDefaultDuration("CHLOE_DISCORD_STREAM_FLUSH_INTERVAL", 500*time.Millisecond),
	SendProcessingMessage:      envOrDefaultBool("CHLOE_DISCORD_SEND_PROCESSING_MESSAGE", false),
	ProcessingMessage:          envOrDefault("CHLOE_DISCORD_PROCESSING_MESSAGE", "↻ Processing..."),
	MaxMessageLength:           envOrDefaultIntInRange("CHLOE_DISCORD_MAX_MESSAGE_LENGTH", 2000, 1, 2000),
}
