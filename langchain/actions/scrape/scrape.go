package scrape

import (
	"fmt"

	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

func (a *ScrapeAction) GetNotification() string {
	return fmt.Sprintf("🌐 Scraping web page: **%s**", a.MustGetParam("url"))
}

func (a *ScrapeAction) Execute(request *action_structs.ActionRequest) ([]*response_object_structs.ResponseObject, error) {
	obj := response_object_structs.NewResponseObject(response_object_structs.Text)

	// truncateTokenCount := structs.GetAvailableTokenCount(request)

	res, err := scrape(a.MustGetParam("url"))
	if err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	/*
		if _, err := obj.Write([]byte(utils2.Truncate(res.GetResponse(), truncateTokenCount))); err != nil {
			return nil, errors.Wrap(errors.ErrActionFailed, err)
		}
	*/
	if _, err := obj.Write([]byte(res.GetResponse())); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*response_object_structs.ResponseObject{obj}, errors.ErrProceed
}

func (a *ScrapeAction) RunPreActions(request *action_structs.ActionRequest) error {
	nurl, err := resolveSpecialUrl(a.MustGetParam("url"))
	if err != nil {
		return errors.Wrap(errors.ErrActionFailed, err)
	}

	a.SetParam("url", nurl)

	return nil
}

func (a *ScrapeAction) RunPostActions(request *action_structs.ActionRequest) error {
	return errors.ErrProceed
}
