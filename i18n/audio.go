package i18n

import "math/rand"

func GetTranscriptionText() string {
	var messages = []string{
		"Transcribing...",
	}

	return messages[rand.Intn(len(messages))]
}

func GetTextToSpeechText() string {
	var messages = []string{
		"Generating audio...",
	}

	return messages[rand.Intn(len(messages))]
}
