package google

import (
	"fmt"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/actions/scrape"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
	"github.com/kamushadenes/chloe/utils"
	googlesearch "github.com/rocketlaunchr/google-search"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

func (a *GoogleAction) GetNotification() string {
	return fmt.Sprintf("🔍 Searching Google: **%s**", a.MustGetParam("query"))
}

func (a *GoogleAction) Execute(request *action_structs.ActionRequest) ([]*response_object_structs.ResponseObject, error) {
	logger := logging.FromContext(request.Context).With().Str("action", a.GetName()).Logger()

	var results []googleResult
	var objs []*response_object_structs.ResponseObject

	fallback := true

	if config.React.GoogleCustomSearchID != "" && config.React.GoogleCustomSearchAPIKey != "" {
		logger.Info().Msg("using openai custom search api")

		svc, err := customsearch.NewService(request.Context, option.WithAPIKey(config.React.GoogleCustomSearchAPIKey))
		if err != nil {
			logger.Warn().Err(err).Msg("failed to create custom search service, falling back to openai search scraping")
		} else {
			s := svc.Cse.List()
			s.Q(a.MustGetParam("query"))
			s.Num(int64(config.React.GoogleMaxResults))
			s.Cx(config.React.GoogleCustomSearchID)
			search, err := s.Do()
			if err != nil {
				logger.Warn().Err(err).Msg("failed to perform api search, falling back to openai search scraping")
			} else {
				for k := range search.Items {
					results = append(results, googleResult{
						URL:         search.Items[k].Link,
						Title:       search.Items[k].Title,
						Description: search.Items[k].Snippet,
					})
				}
				if len(results) == 0 {
					logger.Warn().Msg("no results from api, falling back to openai search scraping")
				} else {
					fallback = false
				}
			}
		}
	}

	if fallback {
		res, err := googlesearch.Search(request.GetContext(), a.MustGetParam("query"), googlesearch.SearchOptions{Limit: config.React.GoogleMaxResults})
		if err != nil {
			return nil, errors.Wrap(errors.ErrActionFailed, err)
		}

		for k := range res {
			results = append(results, googleResult{
				URL:         res[k].URL,
				Title:       res[k].Title,
				Description: res[k].Description,
			})
		}
	}

	if !utils.Testing() {
		for k := range results {
			r := results[k]
			na := scrape.NewScrapeAction()
			na.SetParam("url", r.URL)
			if err := na.RunPreActions(request); err != nil {
				continue
			}
			request.Message.NotifyAction(na.GetNotification())

			aobjs, err := na.Execute(request)
			if err != nil && !errors.Is(err, errors.ErrProceed) {
				continue
			}

			objs = append(objs, aobjs...)
		}
	}

	return objs, errors.ErrProceed
}
func (a *GoogleAction) RunPostActions(request *action_structs.ActionRequest) error {
	return errors.ErrProceed
}
