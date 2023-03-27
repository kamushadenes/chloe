package scrape

import (
	"github.com/gocolly/colly"
	url2 "net/url"
	"strings"
)

func parseContent(c *colly.Collector, scrapeResult *ScrapeResult, u *url2.URL) *NewsSource {
	domain := u.Hostname()

	if strings.HasSuffix(domain, "cnpj.biz") {
		c.OnHTML("div.col-left", func(e *colly.HTMLElement) {
			scrapeResult.SetContent(e.Text)
		})
	} else {
		c.OnHTML("body", func(e *colly.HTMLElement) {
			text, _ := cleanText(e.Text)
			scrapeResult.SetContent(text)
		})
	}

	return nil
}
