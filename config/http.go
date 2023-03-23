package config

type HTTPConfig struct {
	Host string
	Port string
}

var HTTP = &HTTPConfig{
	Host: envOrDefault("CHLOE_HTTP_HOST", "0.0.0.0"),
	Port: envOrDefault("CHLOE_HTTP_PORT", "8080"),
}
