package actions

import (
	"errors"
	"fmt"
	errors3 "github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/react/actions/google"
	"github.com/kamushadenes/chloe/react/actions/image"
	"github.com/kamushadenes/chloe/react/actions/latex"
	"github.com/kamushadenes/chloe/react/actions/math"
	"github.com/kamushadenes/chloe/react/actions/mock"
	"github.com/kamushadenes/chloe/react/actions/news"
	"github.com/kamushadenes/chloe/react/actions/scrape"
	structs2 "github.com/kamushadenes/chloe/react/actions/structs"
	"github.com/kamushadenes/chloe/react/actions/transcribe"
	"github.com/kamushadenes/chloe/react/actions/tts"
	"github.com/kamushadenes/chloe/react/actions/wikipedia"
	"github.com/kamushadenes/chloe/react/actions/youtube_summarizer"
	errors2 "github.com/kamushadenes/chloe/react/errors"
	"github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/structs"
	utils2 "github.com/kamushadenes/chloe/utils"
	"strings"
)

var actions = map[string]func() structs2.Action{
	"mock": mock.NewMockAction,

	"google":            google.NewGoogleAction,
	"news":              news.NewNewsAction,
	"news_by_country":   news.NewNewsByCountryAction,
	"calculate":         math.NewCalculateAction,
	"math":              math.NewCalculateAction,
	"scrape":            scrape.NewScrapeAction,
	"web":               scrape.NewScrapeAction,
	"generate":          image.NewImageAction,
	"tts":               tts.NewTTSAction,
	"transcribe":        transcribe.NewTranscribeAction,
	"variation":         image.NewVariationAction,
	"wikipedia":         wikipedia.NewWikipediaAction,
	"summarize_youtube": youtube_summarizer.NewYoutubeSummarizerAction,
	"latex":             latex.NewLatexAction,
}

func HandleAction(request *structs.ActionRequest) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = utils2.HandlePanic(r)
		}
	}()

	logger := logging.GetLogger().With().Str("action", request.Action).Str("params", request.Params).Logger()

	request.Action = strings.ToLower(request.Action)

	actI, ok := actions[request.Action]
	if !ok {
		return errors3.Wrap(errors3.ErrInvalidAction, fmt.Errorf("action %s not found", request.Action))
	}
	act := actI()

	act.SetParams(request.Params)
	act.SetMessage(request.Message)

	if err = act.RunPreActions(request); err != nil {
		if errors.Is(err, errors2.ErrNotImplemented) {
			if err = defaultPreActions(act, request); err != nil {
				return
			}
		} else {
			return
		}
	}

	if !utils2.Testing() {
		request.Message.NotifyAction(act.GetNotification())
		if err = utils.StoreChainOfThoughtResult(request, act.GetNotification()); err != nil {
			return
		}
	}

	logger.Info().Msg("executing action")
	err = act.Execute(request)
	if err != nil {
		if !errors.Is(err, errors2.ErrProceed) {
			logger.Error().Err(err).Msg("error executing action")
		}
		return
	}

	ws := act.GetWriters()
	for k := range ws {
		if err = ws[k].Close(); err != nil {
			return
		}
	}

	if err = act.RunPostActions(request); err != nil {
		if errors.Is(err, errors2.ErrNotImplemented) {
			err = defaultPostActions(act, request)
		}
	}

	return
}
