package config

import (
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"fmt"
	"github.com/kamushadenes/chloe/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

func envOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func envOrDefaultWithOptions(key, defaultValue string, options []string) string {
	value := envOrDefault(key, defaultValue)
	if !utils.StringInSlice(value, options) {
		panic(fmt.Sprintf("invalid value for %s: %s\nvalid values are %s", key, value, strings.Join(options, ", ")))
	}

	return value
}

func envOrDefaultGCPTTSEncoding(key string, defaultValue texttospeechpb.AudioEncoding) texttospeechpb.AudioEncoding {
	if value := os.Getenv(key); value != "" {
		switch value {
		case "LINEAR16":
			return texttospeechpb.AudioEncoding_LINEAR16
		case "MP3":
			return texttospeechpb.AudioEncoding_MP3
		case "OGG_OPUS":
			return texttospeechpb.AudioEncoding_OGG_OPUS
		case "MULAW":
			return texttospeechpb.AudioEncoding_MULAW
		case "ALAW":
			return texttospeechpb.AudioEncoding_ALAW
		default:
			panic("invalid encoding: " + value)
		}
	}

	return defaultValue
}

func envOrDefaultDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		dur, err := time.ParseDuration(value)
		if err != nil {
			panic(err)
		}
		return dur
	}

	return defaultValue
}

func envOrDefaultBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		b, err := strconv.ParseBool(value)
		if err != nil {
			panic(fmt.Errorf("invalid value for %s: %s", key, value))
		}
		return b
	}

	return defaultValue
}

func envOrDefaultInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		i, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		return i
	}

	return defaultValue
}

func envOrDefaultIntInRange(key string, defaultValue, min, max int) int {
	value := envOrDefaultInt(key, defaultValue)
	if value < min || value > max {
		panic(fmt.Sprintf("invalid value for %s: %d\nvalid values are between %d and %d", key, value, min, max))
	}

	return value
}

func envOrDefaultFloat64(key string, defaultValue float64) float64 {
	if value := os.Getenv(key); value != "" {
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			panic(err)
		}
		return f
	}

	return defaultValue
}

func envOrDefaultFloat64InRange(key string, defaultValue, min, max float64) float64 {
	value := envOrDefaultFloat64(key, defaultValue)
	if value < min || value > max {
		panic(fmt.Sprintf("invalid value for %s: %f\nvalid values are between %f and %f", key, value, min, max))
	}

	return value
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("environment variable %s is not set", key))
	}

	return value
}
