package config

import "time"

type TelegramConfig struct {
	Token                 string
	ImageCount            int
	StreamMessages        bool
	StreamFlushInterval   time.Duration
	SendProcessingMessage bool
	ProcessingMessage     string
}

var Telegram = &TelegramConfig{
	Token:                 envOrDefault("CHLOE_TELEGRAM_TOKEN", ""),
	ImageCount:            envOrDefaultIntInRange("CHLOE_TELEGRAM_IMAGE_COUNT", 4, 1, 10),
	StreamMessages:        envOrDefaultBool("CHLOE_TELEGRAM_STREAM_MESSAGES", true),
	StreamFlushInterval:   envOrDefaultDuration("CHLOE_TELEGRAM_STREAM_FLUSH_INTERVAL", 500*time.Millisecond),
	SendProcessingMessage: envOrDefaultBool("CHLOE_TELEGRAM_SEND_PROCESSING_MESSAGE", true),
	ProcessingMessage:     envOrDefault("CHLOE_TELEGRAM_PROCESSING_MESSAGE", "↻ Processing..."),
}
