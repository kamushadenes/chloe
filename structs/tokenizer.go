package structs

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/utils"
)

func GetAvailableTokenCount(request *action_structs.ActionRequest) int {
	return utils.SubtractIntWithMinimum(
		config.OpenAI.GetMinReplyTokens(),
		config.OpenAI.GetModel(config.Completion).GetContextSize(),
		request.CountTokens(),
		config.OpenAI.GetMinReplyTokens(),
	)
}
