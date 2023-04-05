package i18n

import (
	"github.com/kamushadenes/chloe/utils"
)

func GetTranscriptionText() string {
	return utils.PickRandomString(
		"Transcribing...",
	)
}

func GetTextToSpeechText() string {
	return utils.PickRandomString(
		"Generating audio...",
	)
}
