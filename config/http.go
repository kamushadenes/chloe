package config

type HTTPConfig struct {
	Host string
	Port string
}

var HTTP = &HTTPConfig{
	Host: "0.0.0.0",
	Port: "8080",
}
