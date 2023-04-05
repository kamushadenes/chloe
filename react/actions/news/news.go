package news

import (
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/react/actions/google"
	"github.com/kamushadenes/chloe/react/actions/scrape"
	"github.com/kamushadenes/chloe/structs"
)

type NewsAction struct {
	Name   string
	Params string
}

func NewNewsAction() structs.Action {
	return &NewsAction{
		Name: "news",
	}
}

func (a *NewsAction) SetMessage(message *memory.Message) {}

func (a *NewsAction) GetName() string {
	return a.Name
}

func (a *NewsAction) GetNotification() string {
	return fmt.Sprintf("📰 Searching news: **%s**", a.Params)
}

func (a *NewsAction) SetParams(params string) {
	a.Params = params
}

func (a *NewsAction) GetParams() string {
	return a.Params
}

func (a *NewsAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	logger := logging.FromContext(request.Context).With().Str("action", a.GetName()).Str("query", a.Params).Logger()

	var objs []*structs.ResponseObject

	source := config.React.NewsSource
	if source == "newsapi" && config.React.NewsAPIToken == "" {
		logger.Warn().Msg("NewsAPI token not set, falling back to Google Search")
		source = "google"
	}

	switch source {
	case "google":
		na := google.NewGoogleAction()
		na.SetParams(a.Params)
		na.SetMessage(request.Message)
		request.Message.NotifyAction(na.GetNotification())
		return na.Execute(request)

	case "newsapi":
		res, err := NewsAPIQuery(a.Params)
		if err != nil {
			logger.Error().Err(err).Msg("NewsAPI query failed")
			return nil, errors.Wrap(errors.ErrActionFailed, err)
		}

		cnt := 0
		for _, r := range res.Articles {
			if cnt >= config.React.NewsAPIMaxResults {
				break
			}
			cnt++
			na := scrape.NewScrapeAction()
			na.SetParams(r.URL)
			na.SetMessage(request.Message)
			if err := na.RunPreActions(request); err != nil {
				continue
			}
			aobjs, err := na.Execute(request)
			if err != nil && err != errors.ErrProceed {
				continue
			}
			objs = append(objs, aobjs...)
		}
	}

	return objs, errors.ErrProceed
}

func (a *NewsAction) RunPreActions(request *structs.ActionRequest) error {
	return nil
}

func (a *NewsAction) RunPostActions(request *structs.ActionRequest) error {
	return errors.ErrProceed
}
