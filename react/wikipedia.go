package react

import (
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"github.com/trietmn/go-wiki"
	"io"
)

type WikipediaAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
}

func NewWikipediaAction() *WikipediaAction {
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
	truncateTokenCount := getTokenCount(request)

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
		if err := storeChainOfThoughtResult(request, Truncate(msg, truncateTokenCount)); err != nil {
			return err
		}
	}

	return ErrProceed
}

func (a *WikipediaAction) RunPreActions(request *structs.ActionRequest) error {
	return nil
}

func (a *WikipediaAction) RunPostActions(request *structs.ActionRequest) error {
	return ErrProceed
}
