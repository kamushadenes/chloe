package actions

//go:generate go run action_generator.go

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/actions/file"
	"github.com/kamushadenes/chloe/langchain/actions/google"
	"github.com/kamushadenes/chloe/langchain/actions/image"
	"github.com/kamushadenes/chloe/langchain/actions/latex"
	"github.com/kamushadenes/chloe/langchain/actions/math"
	"github.com/kamushadenes/chloe/langchain/actions/mock"
	"github.com/kamushadenes/chloe/langchain/actions/news"
	"github.com/kamushadenes/chloe/langchain/actions/scrape"
	"github.com/kamushadenes/chloe/langchain/actions/transcribe"
	"github.com/kamushadenes/chloe/langchain/actions/tts"
	"github.com/kamushadenes/chloe/langchain/actions/wikipedia"
	"github.com/kamushadenes/chloe/langchain/actions/youtube"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
	"github.com/kamushadenes/chloe/utils"
)

var actions = map[string]func() action_structs.Action{
	"mock": mock.NewMockAction,

	"search": google.NewGoogleAction,

	"news":            news.NewNewsAction,
	"news_by_country": news.NewNewsByCountryAction,

	"math": math.NewMathAction,

	"scrape": scrape.NewScrapeAction,

	"generate":  image.NewImageAction,
	"variation": image.NewVariationAction,

	"tts": tts.NewTTSAction,

	"transcribe": transcribe.NewTranscribeAction,

	"wikipedia": wikipedia.NewWikipediaAction,

	"summarize_youtube":  youtube.NewYoutubeSummarizeAction,
	"transcribe_youtube": youtube.NewYoutubeSummarizeAction,

	"latex": latex.NewLatexAction,

	"append_file": file.NewAppendFileAction,
	"write_file":  file.NewWriteFileAction,
	"delete_file": file.NewDeleteFileAction,
	"read_file":   file.NewReadFileAction,
}

func HandleAction(request *action_structs.ActionRequest) (err error) {
	paramsJson, _ := json.Marshal(request.Params)

	logger := logging.GetLogger().With().Str("action", request.Action).RawJSON("params", paramsJson).Logger()

	/*
		defer func() {
			if r := recover(); r != nil {
				err = utils.HandlePanic(r)
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

	for k := range request.Params {
		act.SetParam(k, request.Params[k])
	}
	act.SetMessage(request.Message)

	if err = act.CheckRequiredParams(); err != nil {
		return
	}

	if err = act.RunPreActions(request); err != nil {
		if errors.Is(err, errors.ErrNotImplemented) {
			if err = defaultPreActions(act, request); err != nil {
				return
			}
		} else {
			return
		}
	}

	if !utils.Testing() {
		request.Message.NotifyAction(act.GetNotification())
		if err = StoreActionDetectionResult(request, "assistant", act.GetNotification(), ""); err != nil {
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

		if !utils.Testing() {
			var summary string
			switch objs[k].Type {
			case response_object_structs.Image, response_object_structs.Audio:
				summary = objs[k].GetStorableContent()
			}
			if err := StoreActionDetectionResult(request, "user", objs[k].GetStorableContent(), summary); err != nil {
				return errors.Wrap(errors.ErrActionFailed, err)
			}
		}
	}

	if !request.SkipClose {
		if err = request.Writer.Close(); err != nil {
			return err
		}
	}

	if err = act.RunPostActions(request); err != nil {
		if errors.Is(err, errors.ErrNotImplemented) {
			err = defaultPostActions(act, request)
		}
	}

	return
}
