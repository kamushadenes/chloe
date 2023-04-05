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
	"io"
)

type NewsByCountryAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
}

func NewNewsByCountryAction() structs.Action {
	return &NewsByCountryAction{
		Name: "news_by_country",
	}
}

func (a *NewsByCountryAction) SetMessage(message *memory.Message) {}

func (a *NewsByCountryAction) SetWriters(writers []io.WriteCloser) {
	a.Writers = writers
}

func (a *NewsByCountryAction) GetWriters() []io.WriteCloser {
	return a.Writers
}

func (a *NewsByCountryAction) GetName() string {
	return a.Name
}

func (a *NewsByCountryAction) GetNotification() string {
	return fmt.Sprintf("📰 Searching country news: **%s**", a.Params)
}

func (a *NewsByCountryAction) SetParams(params string) {
	a.Params = params
}

func (a *NewsByCountryAction) GetParams() string {
	return a.Params
}

func (a *NewsByCountryAction) Execute(request *structs.ActionRequest) error {
	logger := logging.GetLogger().With().Str("action", a.GetName()).Str("query", a.Params).Logger()

	source := config.React.NewsSource
	if source == "newsapi" && config.React.NewsAPIToken == "" {
		logger.Warn().Msg("NewsAPI token not set, falling back to Google Search")
		source = "google"
	}

	switch source {
	case "google":
		na := google.NewGoogleAction()
		na.SetParams(fmt.Sprintf("%s news", a.Params))
		na.SetMessage(request.Message)
		na.SetWriters(a.Writers)
		request.Message.NotifyAction(na.GetNotification())
		return na.Execute(request)
	case "newsapi":
		res, err := NewsAPITopHeadlines(a.Params)
		if err != nil {
			logger.Error().Err(err).Msg("NewsAPI query failed")
			return errors.Wrap(errors.ErrActionFailed, err)
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
			if err := na.Execute(request); err != nil {
				continue
			}
		}
	}

	return errors.ErrProceed
}

func (a *NewsByCountryAction) RunPreActions(request *structs.ActionRequest) error {
	return nil
}

func (a *NewsByCountryAction) RunPostActions(request *structs.ActionRequest) error {
	return errors.ErrProceed
}
