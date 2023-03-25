package react

import (
	"context"
	"fmt"
	"github.com/anaskhan96/soup"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rocketlaunchr/google-search"
	"github.com/rs/zerolog"
)

func Google(ctx context.Context, request *structs.ScrapeRequest) error {
	logger := zerolog.Ctx(ctx).With().Str("requestID", request.GetID()).Logger()

	logger.Info().Str("args", request.Content).Msg("searching google")

	res, err := googlesearch.Search(ctx, request.Content, googlesearch.SearchOptions{Limit: 5})
	if err != nil {
		return NotifyError(request, err)
	}

	for _, r := range res {
		if _, err := request.Writer.Write([]byte(fmt.Sprintf("URL: %s\n", r.URL))); err != nil {
			return NotifyError(request, err)
		}
		if _, err := request.Writer.Write([]byte(fmt.Sprintf("Title: %s\n", r.Title))); err != nil {
			return NotifyError(request, err)
		}
		if _, err := request.Writer.Write([]byte(fmt.Sprintf("Description: %s\n\n\n", r.Description))); err != nil {
			return NotifyError(request, err)
		}
	}

	if !request.SkipClose {
		err := request.Writer.Close()
		return NotifyError(request, err)
	}

	return NotifyError(request, nil)
}

func Scrape(ctx context.Context, request *structs.ScrapeRequest) error {
	logger := zerolog.Ctx(ctx).With().Str("requestID", request.GetID()).Logger()

	logger.Info().Str("url", request.Content).Msg("scraping page")

	resp, err := soup.Get(request.Content)
	if err != nil {
		return NotifyError(request, err)
	}

	doc := soup.HTMLParse(resp)

	if _, err := request.Writer.Write([]byte(doc.FullText())); err != nil {
		return NotifyError(request, err)
	}
	if !request.SkipClose {
		err := request.Writer.Close()
		return NotifyError(request, err)
	}

	return NotifyError(request, nil)
}
