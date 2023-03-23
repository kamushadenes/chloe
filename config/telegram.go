package config

import "os"

type TelegramConfig struct {
	Token      string
	ImageCount int
}

var Telegram = &TelegramConfig{
	Token:      os.Getenv("CHLOE_TELEGRAM_TOKEN"),
	ImageCount: envOrDefaultInt("CHLOE_TELEGRAM_IMAGE_COUNT", 4),
}
