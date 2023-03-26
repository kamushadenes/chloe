package react

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"io"
	"net/http"
)

type ScrapeAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
}

func NewScrapeAction() Action {
	return &ScrapeAction{
		Name: "scrape",
	}
}

func (a *ScrapeAction) SetWriters(writers []io.WriteCloser) {
	a.Writers = writers
}

func (a *ScrapeAction) GetWriters() []io.WriteCloser {
	return a.Writers
}

func (a *ScrapeAction) GetName() string {
	return a.Name
}

func (a *ScrapeAction) GetNotification() string {
	return fmt.Sprintf("🌐 Scraping web page: **%s**", a.Params)
}

func (a *ScrapeAction) SetParams(params string) {
	a.Params = params
}

func (a *ScrapeAction) GetParams() string {
	return a.Params
}

func (a *ScrapeAction) SetMessage(message *memory.Message) {}

func (a *ScrapeAction) Execute(request *structs.ActionRequest) error {
	truncateTokenCount := getTokenCount(request)

	client := &http.Client{}

	req, err := http.NewRequest("GET", a.Params, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
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

	if err := storeChainOfThoughtResult(request, Truncate(fmt.Sprintf("Content: %s", doc.Text()), truncateTokenCount)); err != nil {
		return err
	}

	return ErrProceed
}

func (a *ScrapeAction) RunPreActions(request *structs.ActionRequest) error {
	return nil
}

func (a *ScrapeAction) RunPostActions(request *structs.ActionRequest) error {
	return ErrProceed
}
