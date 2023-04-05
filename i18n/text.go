package i18n

import (
	"github.com/kamushadenes/chloe/utils"
)

func GetForgetText() string {
	return utils.PickRandomString(
		"I have forgotten you.",
		"Forgot.",
		"Who? Where?",
	)
}
