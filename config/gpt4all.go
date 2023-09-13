package config

type GPT4AllConfig struct {
	Endpoint     string `json:"endpoint"`
	DefaultModel string `json:"default_model"`
}

var GPT4All = &GPT4AllConfig{
	Endpoint: envOrDefault("CHLOE_GPT4ALL_ENDPOINT", "http://localhost:4891/v1"),
	DefaultModel: envOrDefaultWithOptions("CHLOE_GPT4ALL_MODEL", "gpt-3.5-turbo",
		[]string{
			"mpt-7b-chat",
			"gpt-3.5-turbo",
			"gpt-4",
			"gpt4all-j-v1.3-groovy",
		}),
}
