package react

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rocketlaunchr/google-search"
	"io"
)

type GoogleAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
}

func NewGoogleAction() *GoogleAction {
	return &GoogleAction{
		Name: "google",
	}
}

func (a *GoogleAction) SetUser(user *memory.User)          {}
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

func (a *GoogleAction) Execute(ctx context.Context) error {
	res, err := googlesearch.Search(ctx, a.Params, googlesearch.SearchOptions{Limit: 5})
	if err != nil {
		return err
	}

	for _, r := range res {
		for _, w := range a.Writers {
			msg := fmt.Sprintf("URL: %s\nTitle: %s\nDescription: %s\n\n\n", r.URL, r.Title, r.Description)
			_, err := w.Write([]byte(msg))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *GoogleAction) RunPreActions(request *structs.ActionRequest) error {
	return scrapePreActions(a, request)
}

func (a *GoogleAction) RunPostActions(request *structs.ActionRequest) error {
	return ErrProceed
}
