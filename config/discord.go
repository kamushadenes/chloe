package config

import "os"

type DiscordConfig struct {
	Token       string
	ImageCount  int
	OnlyMention bool
}

var Discord = &DiscordConfig{
	Token:       os.Getenv("CHLOE_DISCORD_TOKEN"),
	ImageCount:  envOrDefaultInt("CHLOE_DISCORD_IMAGE_COUNT", 4),
	OnlyMention: envOrDefaultBool("CHLOE_DISCORD_ONLY_MENTION", true),
}
