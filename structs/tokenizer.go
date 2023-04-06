package structs

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/utils"
)

func GetAvailableTokenCount(request *ActionRequest) int {
	return utils.SubtractIntWithMinimum(
		config.OpenAI.GetMinReplyTokens(),
		config.OpenAI.GetModel(config.ChainOfThought).GetContextSize(),
		request.CountTokens(),
		config.OpenAI.GetMinReplyTokens(),
	)
}
