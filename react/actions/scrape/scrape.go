package scrape

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs"
	utils2 "github.com/kamushadenes/chloe/utils"
)

type ScrapeAction struct {
	Name   string
	Params map[string]string
}

func (a *ScrapeAction) GetNotification() string {
	return fmt.Sprintf("🌐 Scraping web page: **%s**", a.Params["url"])
}

func (a *ScrapeAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Text)

	truncateTokenCount := structs.GetAvailableTokenCount(request)

	res, err := scrape(a.Params["url"])
	if err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	if _, err := obj.Write([]byte(utils2.Truncate(res.GetStorableContent(), truncateTokenCount))); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*structs.ResponseObject{obj}, errors.ErrProceed
}

func (a *ScrapeAction) RunPreActions(request *structs.ActionRequest) error {
	nurl, err := resolveSpecialUrl(a.Params["url"])
	if err != nil {
		return errors.Wrap(errors.ErrActionFailed, err)
	}

	a.SetParam("url", nurl)

	return nil
}

func (a *ScrapeAction) RunPostActions(request *structs.ActionRequest) error {
	return errors.ErrProceed
}
