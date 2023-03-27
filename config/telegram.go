package config

type TelegramConfig struct {
	Token      string
	ImageCount int
}

var Telegram = &TelegramConfig{
	Token:      envOrDefault("CHLOE_TELEGRAM_TOKEN", ""),
	ImageCount: envOrDefaultIntInRange("CHLOE_TELEGRAM_IMAGE_COUNT", 4, 1, 10),
}
