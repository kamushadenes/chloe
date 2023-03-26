package react

import (
	"errors"
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs"
	"strings"
)

var actions = map[string]func() Action{
	"google":            NewGoogleAction,
	"calculate":         NewCalculateAction,
	"math":              NewCalculateAction,
	"scrape":            NewScrapeAction,
	"web":               NewScrapeAction,
	"image":             NewImageAction,
	"dalle":             NewImageAction,
	"dall-e":            NewImageAction,
	"audio":             NewAudioAction,
	"tts":               NewAudioAction,
	"speak":             NewAudioAction,
	"transcribe":        NewTranscribeAction,
	"transcription":     NewTranscribeAction,
	"variation":         NewVariationAction,
	"wikipedia":         NewWikipediaAction,
	"summarize_youtube": NewYoutubeSummarizerAction,
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

	actI, ok := actions[request.Action]
	if !ok {
		return fmt.Errorf("unknown request.Action: %s", request.Action)
	}
	act := actI()

	act.SetParams(request.Params)
	act.SetMessage(request.Message)

	if err := act.RunPreActions(request); err != nil {
		return err
	}

	request.Message.NotifyAction(act.GetNotification())
	if err := storeChainOfThoughtResult(request, act.GetNotification()); err != nil {
		return err
	}

	logger.Info().Msg("executing action")
	err := act.Execute(request)
	if err != nil {
		if !errors.Is(err, ErrProceed) {
			logger.Error().Err(err).Msg("error executing action")
		}
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
