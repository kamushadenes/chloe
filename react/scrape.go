package react

import (
	"context"
	"fmt"
	"github.com/anaskhan96/soup"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"io"
)

type ScrapeAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
}

func NewScrapeAction() *ScrapeAction {
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

func (a *ScrapeAction) SetUser(user *memory.User)          {}
func (a *ScrapeAction) SetMessage(message *memory.Message) {}

func (a *ScrapeAction) Execute(ctx context.Context) error {
	resp, err := soup.Get(a.Params)
	if err != nil {
		return err
	}

	doc := soup.HTMLParse(resp)

	for _, w := range a.Writers {
		_, err := w.Write([]byte(doc.FullText()))
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *ScrapeAction) RunPreActions(request *structs.ActionRequest) error {
	return scrapePreActions(a, request)
}

func (a *ScrapeAction) RunPostActions(request *structs.ActionRequest) error {
	return ErrProceed
}
