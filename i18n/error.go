package i18n

import "math/rand"

func GetErrorText(err error) string {
	var messages = []string{
		"I apologize, but I'm unable to complete your request at this time. The following error occurred: " + err.Error(),
		"I'm sorry, but I'm unable to fulfill your request at this moment. The following error occurred: " + err.Error(),
		"I'm afraid I'm unable to complete your request at this time. The following error occurred: " + err.Error(),
	}

	return messages[rand.Intn(len(messages))]
}
