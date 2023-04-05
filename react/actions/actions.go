package actions

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/react/actions/google"
	"github.com/kamushadenes/chloe/react/actions/image"
	"github.com/kamushadenes/chloe/react/actions/latex"
	"github.com/kamushadenes/chloe/react/actions/math"
	"github.com/kamushadenes/chloe/react/actions/mock"
	"github.com/kamushadenes/chloe/react/actions/news"
	"github.com/kamushadenes/chloe/react/actions/scrape"
	"github.com/kamushadenes/chloe/react/actions/transcribe"
	"github.com/kamushadenes/chloe/react/actions/tts"
	"github.com/kamushadenes/chloe/react/actions/wikipedia"
	"github.com/kamushadenes/chloe/react/actions/youtube_summarizer"
	"github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/structs"
	utils2 "github.com/kamushadenes/chloe/utils"
	"strings"
)

var actions = map[string]func() structs.Action{
	"mock": mock.NewMockAction,

	"google": google.NewGoogleAction,
	"search": google.NewGoogleAction,

	"news":            news.NewNewsAction,
	"news_by_country": news.NewNewsByCountryAction,

	"calculate": math.NewCalculateAction,
	"math":      math.NewCalculateAction,

	"scrape": scrape.NewScrapeAction,
	"web":    scrape.NewScrapeAction,

	"generate":  image.NewImageAction,
	"variation": image.NewVariationAction,

	"tts": tts.NewTTSAction,

	"transcribe": transcribe.NewTranscribeAction,

	"wikipedia": wikipedia.NewWikipediaAction,

	"summarize_youtube": youtube_summarizer.NewYoutubeSummarizerAction,

	"latex": latex.NewLatexAction,
}

func HandleAction(request *structs.ActionRequest) (err error) {
	logger := logging.GetLogger().With().Str("action", request.Action).Str("params", request.Params).Logger()

	/*
		defer func() {
			if r := recover(); r != nil {
				err = utils2.HandlePanic(r)
				logger.Error().Err(err).Msg("panic handling action")
			}
		}()
	*/

	request.Action = strings.ToLower(request.Action)

	actI, ok := actions[request.Action]
	if !ok {
		return errors.Wrap(errors.ErrInvalidAction, fmt.Errorf("action %s not found", request.Action))
	}
	act := actI()

	act.SetParams(request.Params)
	act.SetMessage(request.Message)

	if err = act.RunPreActions(request); err != nil {
		if errors.Is(err, errors.ErrNotImplemented) {
			if err = defaultPreActions(act, request); err != nil {
				return
			}
		} else {
			return
		}
	}

	if !utils2.Testing() {
		request.Message.NotifyAction(act.GetNotification())
		if err = utils.StoreActionDetectionResult(request, act.GetNotification()); err != nil {
			logger.Error().Err(err).Msg("error storing action detection result")
			return
		}
	}

	logger.Info().Msg("executing action")
	objs, err := act.Execute(request)
	if err != nil {
		if !errors.Is(err, errors.ErrProceed) {
			logger.Error().Err(err).Msg("error executing action")
			return err
		}
	}

	logger.Info().Msg("writing action results")

	for k := range objs {
		if !errors.Is(err, errors.ErrProceed) {
			// Handle HTTP writers
			for kk := range objs[k].Header() {
				for kkk := range objs[k].Header()[kk] {
					request.Writer.Header().Add(kk, objs[k].Header()[kk][kkk])
				}
			}

			if err = request.Writer.WriteObject(objs[k]); err != nil {
				return err
			}

			request.Writer.WriteHeader(objs[k].HTTPStatusCode)
		}

		if err := utils.StoreActionDetectionResult(request, objs[k].GetStorableContent()); err != nil {
			return errors.Wrap(errors.ErrActionFailed, err)
		}
	}

	if err = request.Writer.Close(); err != nil {
		return err
	}

	if err = act.RunPostActions(request); err != nil {
		if errors.Is(err, errors.ErrNotImplemented) {
			err = defaultPostActions(act, request)
		}
	}

	return
}
