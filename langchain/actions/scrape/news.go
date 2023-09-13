package scrape

import (
	"encoding/json"
	url2 "net/url"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type News struct {
	Title      string
	URL        string
	Summary    string
	Content    string
	Author     string
	Paragraphs []string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	s          *ScrapeResult
}

func NewNews(s *ScrapeResult) *News {
	return &News{s: s}
}

func (n *News) SetTitle(title string) {
	n.Title = strings.TrimSpace(title)
}

func (n *News) SetURL(url string) {
	n.URL = strings.TrimSpace(url)
}

func (n *News) SetSummary(summary string) {
	n.Summary = strings.TrimSpace(summary)
}

func (n *News) SetContent(content string) {
	n.Content = strings.TrimSpace(content)
}

func (n *News) AddParagraph(paragraph string) {
	if paragraph != "" {
		n.Paragraphs = append(n.Paragraphs, strings.TrimSpace(paragraph))
	}
}

func (n *News) SetCreatedAt(createdAt *time.Time) {
	n.CreatedAt = createdAt
}

func (n *News) SetUpdatedAt(updatedAt *time.Time) {
	n.UpdatedAt = updatedAt
}

func (n *News) SetAuthor(author string) {
	n.Author = strings.TrimSpace(author)
}

func (n *News) GetResponse() string {
	resm := make(map[string]string)

	if len(n.Title) > 0 {
		resm["Title"] = n.Title
	} else {
		resm["Title"] = n.s.Title
	}
	resm["URL"] = n.s.URL
	if len(n.Author) > 0 {
		resm["Author"] = n.Author
	}
	if len(n.Summary) > 0 {
		resm["Summary"] = n.Summary
	}
	if len(n.Paragraphs) > 0 {
		resm["Content"] = strings.Join(n.Paragraphs, "\n")
	} else if len(n.Content) > 0 {
		resm["Content"] = n.Content
	} else {
		resm["Content"] = n.s.Content
	}

	resb, _ := json.Marshal(resm)

	return string(resb)
}

func parseNews(c *colly.Collector, scrapeResult *ScrapeResult, u *url2.URL) {
	if source := GetNewsSource(u.Hostname()); source != nil {
		scrapeResult.News = NewNews(scrapeResult)

		c.OnHTML(source.TitleSelector, func(e *colly.HTMLElement) {
			scrapeResult.News.SetTitle(e.Text)
		})

		c.OnHTML(source.SummarySelector, func(e *colly.HTMLElement) {
			scrapeResult.News.SetSummary(e.Text)
		})

		c.OnHTML(source.ContentSelector, func(e *colly.HTMLElement) {
			scrapeResult.News.SetContent(e.Text)
		})

		c.OnHTML(source.ParagraphSelector, func(e *colly.HTMLElement) {
			scrapeResult.News.AddParagraph(e.Text)
		})

		c.OnHTML(source.CreatedAtSelector, func(e *colly.HTMLElement) {
			var createdAt time.Time
			var err error
			if strings.HasPrefix(source.UpdatedAtSelector, "time") {
				createdAt, err = time.Parse(source.TimeFormat, e.Attr("datetime"))
			} else {
				createdAt, err = time.Parse(source.TimeFormat, e.Text)
			}

			if err == nil {
				scrapeResult.News.SetUpdatedAt(&createdAt)
			}
		})

		c.OnHTML(source.UpdatedAtSelector, func(e *colly.HTMLElement) {
			var updatedAt time.Time
			var err error
			if strings.HasPrefix(source.UpdatedAtSelector, "time") {
				updatedAt, err = time.Parse(source.TimeFormat, e.Attr("datetime"))
			} else {
				updatedAt, err = time.Parse(source.TimeFormat, e.Text)
			}

			if err == nil {
				scrapeResult.News.SetUpdatedAt(&updatedAt)
			}
		})
	}
}
