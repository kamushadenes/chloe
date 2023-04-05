package scrape

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	structs2 "github.com/kamushadenes/chloe/react/actions/structs"
	"github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/structs"
	utils2 "github.com/kamushadenes/chloe/utils"
	"io"
)

type ScrapeAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
}

func NewScrapeAction() structs2.Action {
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
	truncateTokenCount := utils.GetAvailableTokenCount(request)

	res, err := scrape(a.Params)
	if err != nil {
		return errors.Wrap(errors.ErrActionFailed, err)
	}

	if err := utils.StoreActionDetectionResult(request, utils2.Truncate(res.GetStorableContent(), truncateTokenCount)); err != nil {
		return errors.Wrap(errors.ErrActionFailed, err)
	}

	return errors.ErrProceed
}

func (a *ScrapeAction) RunPreActions(request *structs.ActionRequest) error {
	nurl, err := resolveSpecialUrl(a.Params)
	if err != nil {
		return errors.Wrap(errors.ErrActionFailed, err)
	}

	a.SetParams(nurl)

	return nil
}

func (a *ScrapeAction) RunPostActions(request *structs.ActionRequest) error {
	return errors.ErrProceed
}
