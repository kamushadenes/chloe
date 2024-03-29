package scrape

import (
	"encoding/json"
	"strings"
)

type Link struct {
	Text string
	URL  string
}

type Image struct {
	Alt   string
	URL   string
	Image []byte
}

type ScrapeResult struct {
	URL     string
	Title   string
	Content string
	HTML    string
	Links   []*Link
	Images  []*Image
	News    *News
}

func NewScrapeResult() *ScrapeResult {
	return &ScrapeResult{}
}

func (s *ScrapeResult) SetURL(url string) {
	s.URL = strings.TrimSpace(url)
}

func (s *ScrapeResult) SetTitle(title string) {
	s.Title = strings.TrimSpace(title)
}

func (s *ScrapeResult) SetContent(content string) {
	s.Content = strings.TrimSpace(content)
}

func (s *ScrapeResult) SetHTML(html string) {
	s.HTML = strings.TrimSpace(html)
}

func (s *ScrapeResult) AddLink(url string, text string) {
	s.Links = append(s.Links, &Link{URL: strings.TrimSpace(url), Text: strings.TrimSpace(text)})
}

func (s *ScrapeResult) AddImage(url string, alt string) {
	s.Images = append(s.Images, &Image{URL: strings.TrimSpace(url), Alt: strings.TrimSpace(alt)})
}

func (s *ScrapeResult) GetResponse() string {
	if s.News != nil {
		return s.News.GetResponse()
	}

	var resm = make(map[string]string)
	resm["URL"] = s.URL
	resm["Title"] = s.Title
	resm["Content"] = s.Content

	resb, _ := json.Marshal(resm)

	return string(resb)
}
