package config

import (
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"fmt"
	"os"
	"strconv"
	"time"
)

func envOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
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

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("environment variable %s is not set", key))
	}

	return value
}
