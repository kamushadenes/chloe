package wikipedia

import (
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/structs"
	utils2 "github.com/kamushadenes/chloe/utils"
	"github.com/trietmn/go-wiki"
)

type WikipediaAction struct {
	Name   string
	Params string
}

func NewWikipediaAction() structs.Action {
	return &WikipediaAction{
		Name: "wikipedia",
	}
}

func (a *WikipediaAction) SetMessage(message *memory.Message) {}

func (a *WikipediaAction) GetName() string {
	return a.Name
}

func (a *WikipediaAction) GetNotification() string {
	return fmt.Sprintf("🔍 Searching Wikipedia: **%s**", a.Params)
}

func (a *WikipediaAction) SetParams(params string) {
	a.Params = params
}

func (a *WikipediaAction) GetParams() string {
	return a.Params
}

func (a *WikipediaAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Text)

	var truncateTokenCount int
	if utils2.Testing() {
		truncateTokenCount = 1000
	} else {
		truncateTokenCount = utils.GetAvailableTokenCount(request)
	}

	res, _, err := gowiki.Search(a.Params, config.React.WikipediaMaxResults, false)
	if err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	for _, r := range res {
		page, err := gowiki.GetPage(r, -1, false, true)
		if err != nil {
			if utils2.Testing() {
				return nil, errors.Wrap(errors.ErrActionFailed, err)
			}
			continue
		}
		content, err := page.GetContent()
		if err != nil {
			if utils2.Testing() {
				return nil, errors.Wrap(errors.ErrActionFailed, err)
			}
			continue
		}

		if _, err := obj.Write([]byte(
			fmt.Sprintf("URL: %s\nTitle: %s\nContent: %s",
				page.URL, page.Title, utils2.Truncate(content, truncateTokenCount)))); err != nil {
			return nil, errors.Wrap(errors.ErrActionFailed, err)
		}

		// TODO: store in the main loop outside
		/*
			if !utils2.Testing() {
				msg := fmt.Sprintf("URL: %s\nTitle: %s\nContent: %s", page.URL, page.Title, content)
				if err := utils.StoreActionDetectionResult(request, utils2.Truncate(msg, truncateTokenCount)); err != nil {
					return nil, errors.Wrap(errors.ErrActionFailed, err)
				}
			}
		*/
	}

	return []*structs.ResponseObject{obj}, errors.ErrProceed
}

func (a *WikipediaAction) RunPreActions(request *structs.ActionRequest) error {
	return nil
}

func (a *WikipediaAction) RunPostActions(request *structs.ActionRequest) error {
	return errors.ErrProceed
}
