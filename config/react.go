package config

type ReactConfig struct {
	ImproveImagePrompts      bool
	GoogleMaxResults         int
	GoogleCustomSearchAPIKey string
	GoogleCustomSearchID     string
	WikipediaMaxResults      int
	NewsSource               string
	NewsAPIToken             string
	NewsAPIMaxResults        int
	NewsAPISortStrategy      string
	UseAria2                 bool
	ReportThoughts           bool
	FileWorkspace            string
}

var React = &ReactConfig{
	ImproveImagePrompts:      envOrDefaultBool("CHLOE_REACT_IMPROVE_IMAGE_PROMPTS", false),
	GoogleMaxResults:         envOrDefaultIntInRange("CHLOE_REACT_GOOGLE_MAX_RESULTS", 4, 1, 10),
	GoogleCustomSearchAPIKey: envOrDefault("CHLOE_REACT_GOOGLE_CUSTOM_SEARCH_API_KEY", ""),
	GoogleCustomSearchID:     envOrDefault("CHLOE_REACT_GOOGLE_CUSTOM_SEARCH_ID", ""),
	WikipediaMaxResults:      envOrDefaultInt("CHLOE_REACT_WIKIPEDIA_MAX_RESULTS", 3),
	NewsAPIMaxResults:        envOrDefaultInt("CHLOE_REACT_NEWSAPI_MAX_RESULTS", 5),
	NewsSource: envOrDefaultWithOptions("CHLOE_REACT_NEWS_SOURCE", "openai",
		[]string{"openai", "newsapi"}),
	NewsAPIToken: envOrDefault("CHLOE_REACT_NEWSAPI_TOKEN", ""),
	NewsAPISortStrategy: envOrDefaultWithOptions("CHLOE_REACT_NEWSAPI_SORT_STRATEGY", "relevancy",
		[]string{"relevancy", "popularity", "publishedAt"}),
	UseAria2:       envOrDefaultBool("CHLOE_REACT_USE_ARIA2", true),
	ReportThoughts: envOrDefaultBool("CHLOE_REACT_REPORT_THOUGHTS", false),
	FileWorkspace:  envOrDefault("CHLOE_REACT_FILE_WORKSPACE", "workspace/"),
}
