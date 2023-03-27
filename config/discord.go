package config

import "time"

type DiscordConfig struct {
	Token                      string
	ImageCount                 int
	OnlyMention                bool
	RandomStatusUpdateInterval time.Duration
	StreamMessages             bool
	StreamFlushInterval        time.Duration
}

var Discord = &DiscordConfig{
	Token:                      envOrDefault("CHLOE_DISCORD_TOKEN", ""),
	ImageCount:                 envOrDefaultIntInRange("CHLOE_DISCORD_IMAGE_COUNT", 4, 1, 10),
	OnlyMention:                envOrDefaultBool("CHLOE_DISCORD_ONLY_MENTION", true),
	RandomStatusUpdateInterval: envOrDefaultDuration("CHLOE_DISCORD_RANDOM_STATUS_UPDATE_INTERVAL", 20*time.Second),
	StreamMessages:             envOrDefaultBool("CHLOE_DISCORD_STREAM_MESSAGES", false),
	StreamFlushInterval:        envOrDefaultDuration("CHLOE_DISCORD_STREAM_FLUSH_INTERVAL", 500*time.Millisecond),
}
