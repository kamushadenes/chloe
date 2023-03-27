package config

type ReactConfig struct {
	ImproveImagePrompts bool
	GoogleMaxResults    int
	WikipediaMaxResults int
	NewsSource          string
	NewsAPIToken        string
}

var React = &ReactConfig{
	ImproveImagePrompts: envOrDefaultBool("CHLOE_REACT_IMPROVE_IMAGE_PROMPTS", false),
	GoogleMaxResults:    envOrDefaultInt("CHLOE_REACT_GOOGLE_MAX_RESULTS", 4),
	WikipediaMaxResults: envOrDefaultInt("CHLOE_REACT_WIKIPEDIA_MAX_RESULTS", 3),
	NewsSource:          envOrDefaultWithOptions("CHLOE_REACT_NEWS_SOURCE", "google", []string{"google", "newsapi"}),
	NewsAPIToken:        envOrDefault("CHLOE_REACT_NEWSAPI_TOKEN", ""),
}
