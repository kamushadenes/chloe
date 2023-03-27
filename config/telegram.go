package config

import "time"

type TelegramConfig struct {
	Token                 string
	ImageCount            int
	StreamMessages        bool
	StreamFlushInterval   time.Duration
	SendProcessingMessage bool
	ProcessingMessage     string
	MaxMessageLength      int
}

var Telegram = &TelegramConfig{
	Token:                 envOrDefault("CHLOE_TELEGRAM_TOKEN", ""),
	ImageCount:            envOrDefaultIntInRange("CHLOE_TELEGRAM_IMAGE_COUNT", 4, 1, 10),
	StreamMessages:        envOrDefaultBool("CHLOE_TELEGRAM_STREAM_MESSAGES", false),
	StreamFlushInterval:   envOrDefaultDuration("CHLOE_TELEGRAM_STREAM_FLUSH_INTERVAL", 500*time.Millisecond),
	SendProcessingMessage: envOrDefaultBool("CHLOE_TELEGRAM_SEND_PROCESSING_MESSAGE", false),
	ProcessingMessage:     envOrDefault("CHLOE_TELEGRAM_PROCESSING_MESSAGE", "↻ Processing..."),
	MaxMessageLength:      envOrDefaultIntInRange("CHLOE_TELEGRAM_MAX_MESSAGE_LENGTH", 4096, 1, 4096),
}
