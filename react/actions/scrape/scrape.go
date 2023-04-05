package scrape

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/structs"
	utils2 "github.com/kamushadenes/chloe/utils"
)

type ScrapeAction struct {
	Name   string
	Params string
}

func NewScrapeAction() structs.Action {
	return &ScrapeAction{
		Name: "scrape",
	}
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

func (a *ScrapeAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Text)

	truncateTokenCount := utils.GetAvailableTokenCount(request)

	res, err := scrape(a.Params)
	if err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	if _, err := obj.Write([]byte(utils2.Truncate(res.GetStorableContent(), truncateTokenCount))); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	// TODO: store in the main loop outside
	/*
		if err := utils.StoreActionDetectionResult(request, utils2.Truncate(res.GetStorableContent(), truncateTokenCount)); err != nil {
			return errors.Wrap(errors.ErrActionFailed, err)
		}
	*/

	return []*structs.ResponseObject{obj}, errors.ErrProceed
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
