package react

import (
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs"
	"strings"
)

var actions = map[string]Action{
	"google":        NewGoogleAction(),
	"calculate":     NewCalculateAction(),
	"math":          NewCalculateAction(),
	"scrape":        NewScrapeAction(),
	"web":           NewScrapeAction(),
	"image":         NewImageAction(),
	"dalle":         NewImageAction(),
	"dall-e":        NewImageAction(),
	"audio":         NewAudioAction(),
	"tts":           NewAudioAction(),
	"speak":         NewAudioAction(),
	"transcribe":    NewTranscribeAction(),
	"transcription": NewTranscribeAction(),
	"variation":     NewVariationAction(),
	"wikipedia":     NewWikipediaAction(),
}

func getTokenCount(request *structs.ActionRequest) int {
	tokenCount := request.CountTokens()
	truncateTokenCount := config.OpenAI.MaxTokens[config.OpenAI.DefaultModel.ChainOfThought] - tokenCount
	if truncateTokenCount < config.OpenAI.MinReplyTokens[config.OpenAI.DefaultModel.ChainOfThought] {
		truncateTokenCount = config.OpenAI.MinReplyTokens[config.OpenAI.DefaultModel.ChainOfThought]
	}

	return truncateTokenCount
}

func HandleAction(request *structs.ActionRequest) error {
	logger := logging.GetLogger().With().Str("action", request.Action).Str("params", request.Params).Logger()

	request.Action = strings.ToLower(request.Action)

	act, ok := actions[request.Action]
	if !ok {
		return fmt.Errorf("unknown request.Action: %s", request.Action)
	}

	act.SetParams(request.Params)
	act.SetMessage(request.Message)

	request.Message.NotifyAction(act.GetNotification())
	if err := storeChainOfThoughtResult(request, act.GetNotification()); err != nil {
		return err
	}

	if err := act.RunPreActions(request); err != nil {
		return err
	}

	logger.Info().Msg("executing action")
	err := act.Execute(request)
	if err != nil {
		return err
	}

	ws := act.GetWriters()
	for k := range ws {
		if err := ws[k].Close(); err != nil {
			return err
		}
	}

	if err := act.RunPostActions(request); err != nil {
		return err
	}

	return nil
}
