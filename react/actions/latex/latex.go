package latex

import (
	"fmt"
	"github.com/go-latex/latex/drawtex/drawimg"
	"github.com/go-latex/latex/mtex"
	"github.com/kamushadenes/chloe/memory"
	structs2 "github.com/kamushadenes/chloe/react/actions/structs"
	"github.com/kamushadenes/chloe/structs"
	"io"
)

type LatexAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
}

func NewLatexAction() structs2.Action {
	return &LatexAction{
		Name: "image",
	}
}

func (a *LatexAction) SetWriters(writers []io.WriteCloser) {
	a.Writers = writers
}

func (a *LatexAction) GetWriters() []io.WriteCloser {
	return a.Writers
}
func (a *LatexAction) GetName() string {
	return a.Name
}

func (a *LatexAction) GetNotification() string {
	return fmt.Sprintf("⚗️ Rendering latex: **%s**", a.Params)
}

func (a *LatexAction) SetParams(params string) {
	a.Params = params
}

func (a *LatexAction) GetParams() string {
	return a.Params
}

func (a *LatexAction) SetMessage(message *memory.Message) {}

func (a *LatexAction) Execute(request *structs.ActionRequest) error {
	for _, w := range a.Writers {
		dst := drawimg.NewRenderer(w)
		if err := mtex.Render(dst, request.Params, 64, 600, nil); err != nil {
			return err
		}
	}

	return nil
}

func (a *LatexAction) RunPreActions(request *structs.ActionRequest) error {
	return latexPreActions(a, request)
}

func (a *LatexAction) RunPostActions(request *structs.ActionRequest) error {
	return nil
}
