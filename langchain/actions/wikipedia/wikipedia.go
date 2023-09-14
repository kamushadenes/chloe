package wikipedia

import (
	"encoding/json"
	"fmt"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
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

	var truncateTokenCount = 1000

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

		var resm = make(map[string]string)
		resm["URL"] = page.URL
		resm["Title"] = page.Title
		resm["Content"] = utils.Truncate(content, truncateTokenCount)

		resb, err := json.Marshal(resm)
		if err != nil {
			if utils.Testing() {
				return nil, errors.Wrap(errors.ErrActionFailed, err)
			}
		}

		if _, err := obj.Write(resb); err != nil {
			return nil, errors.Wrap(errors.ErrActionFailed, err)
		}
	}

	return []*response_object_structs.ResponseObject{obj}, errors.ErrProceed
}

func (a *WikipediaAction) RunPostActions(request *action_structs.ActionRequest) error {
	return errors.ErrProceed
}
