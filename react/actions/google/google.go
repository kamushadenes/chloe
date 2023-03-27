package google

import (
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/react/actions/scrape"
	structs2 "github.com/kamushadenes/chloe/react/actions/structs"
	"github.com/kamushadenes/chloe/react/errors"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rocketlaunchr/google-search"
	"io"
)

type GoogleAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
}

func NewGoogleAction() structs2.Action {
	return &GoogleAction{
		Name: "google",
	}
}

func (a *GoogleAction) SetMessage(message *memory.Message) {}

func (a *GoogleAction) SetWriters(writers []io.WriteCloser) {
	a.Writers = writers
}

func (a *GoogleAction) GetWriters() []io.WriteCloser {
	return a.Writers
}

func (a *GoogleAction) GetName() string {
	return a.Name
}

func (a *GoogleAction) GetNotification() string {
	return fmt.Sprintf("🔍 Searching Google: **%s**", a.Params)
}

func (a *GoogleAction) SetParams(params string) {
	a.Params = params
}

func (a *GoogleAction) GetParams() string {
	return a.Params
}

func (a *GoogleAction) Execute(request *structs.ActionRequest) error {
	res, err := googlesearch.Search(request.GetContext(), a.Params, googlesearch.SearchOptions{Limit: config.React.GoogleMaxResults})
	if err != nil {
		return err
	}

	for _, r := range res {
		na := scrape.NewScrapeAction()
		na.SetParams(r.URL)
		na.SetMessage(request.Message)
		if err := na.RunPreActions(request); err != nil {
			continue
		}
		request.Message.NotifyAction(na.GetNotification())
		if err := na.Execute(request); err != nil {
			continue
		}
	}

	return errors.ErrProceed
}

func (a *GoogleAction) RunPreActions(request *structs.ActionRequest) error {
	return nil
}

func (a *GoogleAction) RunPostActions(request *structs.ActionRequest) error {
	return errors.ErrProceed
}