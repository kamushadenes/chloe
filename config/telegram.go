package config

import "os"

type TelegramConfig struct {
	Token string
}

var Telegram = &TelegramConfig{
	Token: os.Getenv("TELEGRAM_TOKEN"),
}
