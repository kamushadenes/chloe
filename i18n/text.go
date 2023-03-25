package i18n

import "math/rand"

func GetForgetText() string {
	var messages = []string{
		"I have forgotten you.",
		"Forgot.",
		"Who? Where?",
	}

	return messages[rand.Intn(len(messages))]
}
