package wikipedia

import (
	"fmt"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
	"github.com/kamushadenes/chloe/utils"
	gowiki "github.com/trietmn/go-wiki"
)

func (a *WikipediaAction) GetNotification() string {
	return fmt.Sprintf("🔍 Searching Wikipedia: **%s**", a.MustGetParam("query"))
}

func (a *WikipediaAction) Execute(request *action_structs.ActionRequest) ([]*response_object_structs.ResponseObject, error) {
	obj := response_object_structs.NewResponseObject(response_object_structs.Text)

	var truncateTokenCount int
	if utils.Testing() {
		truncateTokenCount = 1000
	} else {
		truncateTokenCount = structs.GetAvailableTokenCount(request)
	}

	res, _, err := gowiki.Search(a.MustGetParam("query"), config.React.WikipediaMaxResults, false)
	if err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	for _, r := range res {
		page, err := gowiki.GetPage(r, -1, false, true)
		if err != nil {
			if utils.Testing() {
				return nil, errors.Wrap(errors.ErrActionFailed, err)
			}
			continue
		}
		content, err := page.GetContent()
		if err != nil {
			if utils.Testing() {
				return nil, errors.Wrap(errors.ErrActionFailed, err)
			}
			continue
		}

		if _, err := obj.Write([]byte(
			fmt.Sprintf("URL: %s\nTitle: %s\nContent: %s",
				page.URL, page.Title, utils.Truncate(content, truncateTokenCount)))); err != nil {
			return nil, errors.Wrap(errors.ErrActionFailed, err)
		}
	}

	return []*response_object_structs.ResponseObject{obj}, errors.ErrProceed
}

func (a *WikipediaAction) RunPostActions(request *action_structs.ActionRequest) error {
	return errors.ErrProceed
}
