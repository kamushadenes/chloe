package news

import (
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/react/actions/google"
	"github.com/kamushadenes/chloe/react/actions/scrape"
	"github.com/kamushadenes/chloe/structs"
)

type NewsAction struct {
	Name   string
	Params map[string]string
}

func (a *NewsAction) GetNotification() string {
	return fmt.Sprintf("📰 Searching news: **%s**", a.Params["query"])
}

func (a *NewsAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	logger := logging.FromContext(request.Context).With().Str("action", a.GetName()).Str("query", a.Params["query"]).Logger()

	var objs []*structs.ResponseObject

	source := config.React.NewsSource
	if source == "newsapi" && config.React.NewsAPIToken == "" {
		logger.Warn().Msg("NewsAPI token not set, falling back to Google Search")
		source = "google"
	}

	switch source {
	case "google":
		na := google.NewGoogleAction()
		na.SetParam("query", a.Params["query"])
		request.Message.NotifyAction(na.GetNotification())
		return na.Execute(request)

	case "newsapi":
		res, err := NewsAPIQuery(a.Params["query"])
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
			na.SetParam("url", r.URL)
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
