package scrape

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/kamushadenes/chloe/logging"
	url2 "net/url"
	"strings"
)

func cleanText(text string) (string, error) {
	reader := strings.NewReader(text)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return text, err
	}

	doc.Find("script").Each(func(i int, el *goquery.Selection) {
		el.Remove()
	})
	doc.Find("link").Each(func(i int, el *goquery.Selection) {
		el.Remove()
	})
	doc.Find("style").Each(func(i int, el *goquery.Selection) {
		el.Remove()
	})
	doc.Find("iframe").Each(func(i int, el *goquery.Selection) {
		el.Remove()
	})

	return doc.Text(), nil
}

func scrape(url string) (*ScrapeResult, error) {
	logger := logging.GetLogger().With().Str("action", "scrape").Str("url", url).Logger()

	scrapeResult := NewScrapeResult()

	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

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
		html, err := e.DOM.Html()
		if err == nil {
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

	u, err := url2.Parse(url)
	if err != nil {
		return nil, err
	}

	parseContent(c, scrapeResult, u)

	parseNews(c, scrapeResult, u)

	err = c.Visit(url)

	return scrapeResult, err
}
