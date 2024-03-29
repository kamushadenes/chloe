package scrape

import (
	"encoding/base64"
	url2 "net/url"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/kamushadenes/chloe/logging"
	"mvdan.cc/xurls/v2"
)

func clean(c *colly.Collector) {
	// Try to clean up scripts and other stuff
	for _, s := range []string{"script", "link", "style", "iframe"} {
		c.OnHTML(s, func(e *colly.HTMLElement) {
			e.DOM.ReplaceWithHtml("")
		})
	}
}

func scrape(url string) (*ScrapeResult, error) {
	logger := logging.GetLogger().With().Str("action", "scrape").Str("url", url).Logger()

	u, err := url2.Parse(url)
	if err != nil {
		return nil, err
	}

	scrapeResult := NewScrapeResult()

	c := colly.NewCollector()

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	clean(c)

	c.OnRequest(func(r *colly.Request) {
		scrapeResult.SetURL(r.URL.String())

		logger.Info().Msg("scraping")
	})

	c.OnError(func(r *colly.Response, err error) {
		logger.Error().Err(err).Msg("error scraping")
	})

	c.OnScraped(func(r *colly.Response) {
		logger.Info().Msg("finished scraping")
	})

	c.OnHTML("html", func(e *colly.HTMLElement) {
		if html, err := e.DOM.Html(); err == nil {
			scrapeResult.SetHTML(html)
		}
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		scrapeResult.SetTitle(e.Text)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		scrapeResult.AddLink(e.Attr("href"), e.Text)
	})

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		scrapeResult.AddImage(e.Attr("src"), e.Attr("alt"))
	})

	parseContent(c, scrapeResult, u)

	parseNews(c, scrapeResult, u)

	err = c.Visit(url)

	return scrapeResult, err
}

func resolveSpecialUrl(url string) (string, error) {
	u, err := url2.Parse(url)
	if err != nil {
		return "", err
	}

	// Handle special cases
	if strings.HasSuffix(u.Hostname(), "news.google.com") && strings.HasPrefix(u.Path, "/rss/articles") {
		fields := strings.Split(u.Path, "/")
		encodedUrl := fields[len(fields)-1]
		decodedUrl, err := base64.StdEncoding.DecodeString(encodedUrl)
		if err != nil {
			return "", err
		}

		rxRelaxed := xurls.Relaxed()
		nurl := rxRelaxed.FindString(string(decodedUrl))
		return nurl, nil
	}

	return url, nil
}
