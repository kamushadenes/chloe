package config

type DiscordConfig struct {
	Token       string
	ImageCount  int
	OnlyMention bool
}

var Discord = &DiscordConfig{
	Token:       envOrDefault("CHLOE_DISCORD_TOKEN", ""),
	ImageCount:  envOrDefaultIntInRange("CHLOE_DISCORD_IMAGE_COUNT", 4, 1, 10),
	OnlyMention: envOrDefaultBool("CHLOE_DISCORD_ONLY_MENTION", true),
}
