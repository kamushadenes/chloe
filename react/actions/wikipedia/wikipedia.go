package wikipedia

import (
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	structs2 "github.com/kamushadenes/chloe/react/actions/structs"
	"github.com/kamushadenes/chloe/react/errors"
	"github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/structs"
	"github.com/trietmn/go-wiki"
	"io"
)

type WikipediaAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
}

func NewWikipediaAction() structs2.Action {
	return &WikipediaAction{
		Name: "wikipedia",
	}
}

func (a *WikipediaAction) SetMessage(message *memory.Message) {}

func (a *WikipediaAction) SetWriters(writers []io.WriteCloser) {
	a.Writers = writers
}

func (a *WikipediaAction) GetWriters() []io.WriteCloser {
	return a.Writers
}

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

func (a *WikipediaAction) Execute(request *structs.ActionRequest) error {
	truncateTokenCount := utils.GetAvailableTokenCount(request)

	res, _, err := gowiki.Search(a.Params, config.React.WikipediaMaxResults, false)
	if err != nil {
		return err
	}

	for _, r := range res {
		page, err := gowiki.GetPage(r, -1, false, true)
		if err != nil {
			continue
		}
		content, err := page.GetContent()
		if err != nil {
			continue
		}
		msg := fmt.Sprintf("URL: %s\nTitle: %s\nContent: %s", page.URL, page.Title, content)
		if err := utils.StoreChainOfThoughtResult(request, utils.Truncate(msg, truncateTokenCount)); err != nil {
			return err
		}
	}

	return errors.ErrProceed
}

func (a *WikipediaAction) RunPreActions(request *structs.ActionRequest) error {
	return nil
}

func (a *WikipediaAction) RunPostActions(request *structs.ActionRequest) error {
	return errors.ErrProceed
}
