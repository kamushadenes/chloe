package config

type ReactConfig struct {
	ImproveImagePrompts bool
	GoogleMaxResults    int
	WikipediaMaxResults int
}

var React = &ReactConfig{
	ImproveImagePrompts: envOrDefaultBool("CHLOE_REACT_IMPROVE_IMAGE_PROMPTS", false),
	GoogleMaxResults:    envOrDefaultInt("CHLOE_REACT_GOOGLE_MAX_RESULTS", 3),
	WikipediaMaxResults: envOrDefaultInt("CHLOE_REACT_WIKIPEDIA_MAX_RESULTS", 2),
}
