package actions

import (
	"errors"
	"fmt"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/react/actions/google"
	"github.com/kamushadenes/chloe/react/actions/image"
	"github.com/kamushadenes/chloe/react/actions/math"
	"github.com/kamushadenes/chloe/react/actions/scrape"
	structs2 "github.com/kamushadenes/chloe/react/actions/structs"
	"github.com/kamushadenes/chloe/react/actions/transcribe"
	"github.com/kamushadenes/chloe/react/actions/tts"
	"github.com/kamushadenes/chloe/react/actions/wikipedia"
	"github.com/kamushadenes/chloe/react/actions/youtube_summarizer"
	errors2 "github.com/kamushadenes/chloe/react/errors"
	"github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/structs"
	"strings"
)

var actions = map[string]func() structs2.Action{
	"google":            google.NewGoogleAction,
	"news":              google.NewGoogleAction,
	"calculate":         math.NewCalculateAction,
	"math":              math.NewCalculateAction,
	"scrape":            scrape.NewScrapeAction,
	"web":               scrape.NewScrapeAction,
	"image":             image.NewImageAction,
	"dalle":             image.NewImageAction,
	"dall-e":            image.NewImageAction,
	"audio":             tts.NewAudioAction,
	"tts":               tts.NewAudioAction,
	"speak":             tts.NewAudioAction,
	"transcribe":        transcribe.NewTranscribeAction,
	"transcription":     transcribe.NewTranscribeAction,
	"variation":         image.NewVariationAction,
	"wikipedia":         wikipedia.NewWikipediaAction,
	"summarize_youtube": youtube_summarizer.NewYoutubeSummarizerAction,
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
		if errors.Is(err, errors2.ErrNotImplemented) {
			if err := defaultPreActions(act, request); err != nil {
				return err
			}
		}
		return err
	}

	request.Message.NotifyAction(act.GetNotification())
	if err := utils.StoreChainOfThoughtResult(request, act.GetNotification()); err != nil {
		return err
	}

	logger.Info().Msg("executing action")
	err := act.Execute(request)
	if err != nil {
		if !errors.Is(err, errors2.ErrProceed) {
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
		if errors.Is(err, errors2.ErrNotImplemented) {
			return defaultPostActions(act, request)
		}
		return err
	}

	return nil
}
