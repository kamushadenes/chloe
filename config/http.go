package config

type HTTPConfig struct {
	Host string
	Port int
}

var HTTP = &HTTPConfig{
	Host: envOrDefault("CHLOE_HTTP_HOST", "0.0.0.0"),
	Port: envOrDefaultIntInRange("CHLOE_HTTP_PORT", 8080, 1, 65535),
}
