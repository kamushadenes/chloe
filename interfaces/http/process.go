package http

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/messages"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/users"
	"io"
	"net/http"
)

func aiContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqId := middleware.GetReqID(ctx)

		if reqId != "" {
			msg := messages.NewMessage(reqId)
			msg.Source.HTTP = &messages.HTTPMessageSource{
				Request: r,
			}
			msg.User, _ = users.GetUser(ctx, "1")
			msg.User.ExternalID = &users.ExternalID{
				ID:        fmt.Sprintf("%d", 1),
				Interface: "http",
			}

			ctx = context.WithValue(ctx, "msg", msg)

			channels.IncomingMessagesCh <- msg
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func aiComplete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	msg := ctx.Value("msg").(*messages.Message)

	var params = struct {
		Content string                 `json:"content"`
		Args    map[string]interface{} `json:"args"`
	}{}
	if err := parseFromRequest(r, &params); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if params.Content == "" {
		_ = render.Render(w, r, ErrInvalidRequest(fmt.Errorf("content is required")))
		return
	}

	request := &structs.CompletionRequest{}
	request.Context = ctx

	request.User = msg.User

	if request.Mode == "" {
		mode, _ := memory.GetUserMode(ctx, request.User.ID)
		request.Mode = mode
	}

	request.User.ExternalID = &users.ExternalID{
		ID:        fmt.Sprintf("%d", 1),
		Interface: "http",
	}

	request.Content = params.Content
	request.Args = params.Args

	request.Writer = NewHTTPResponseWriteCloser(w)

	channels.CompletionRequestsCh <- request

	for {
		select {
		case <-ctx.Done():
			return
		case <-request.Writer.(*HTTPResponseWriteCloser).CloseCh:
			return
		}
	}
}

func aiGenerate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	msg := ctx.Value("msg").(*messages.Message)

	var params = struct {
		Prompt string `json:"prompt"`
		Size   string `json:"size"`
	}{}
	if err := parseFromRequest(r, &params); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if params.Prompt == "" {
		_ = render.Render(w, r, ErrInvalidRequest(fmt.Errorf("prompt is required")))
		return
	}

	request := &structs.GenerationRequest{}
	request.Context = ctx

	request.User = msg.User

	request.User.ExternalID = &users.ExternalID{
		ID:        fmt.Sprintf("%d", 1),
		Interface: "http",
	}

	request.Prompt = params.Prompt
	request.Size = params.Size

	request.Writers = []io.WriteCloser{NewHTTPResponseWriteCloser(w)}

	channels.GenerationRequestsCh <- request

	for {
		select {
		case <-ctx.Done():
			return
		case <-request.Writers[0].(*HTTPResponseWriteCloser).CloseCh:
			return
		}
	}
}

func aiTTS(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	msg := ctx.Value("msg").(*messages.Message)

	var params = struct {
		Content string                 `json:"content"`
		Args    map[string]interface{} `json:"args"`
	}{}
	if err := parseFromRequest(r, &params); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if params.Content == "" {
		_ = render.Render(w, r, ErrInvalidRequest(fmt.Errorf("content is required")))
		return
	}

	request := &structs.TTSRequest{}
	request.Context = ctx

	request.User = msg.User

	request.Content = params.Content

	request.Writers = []io.WriteCloser{NewHTTPResponseWriteCloser(w)}

	channels.TTSRequestsCh <- request

	for {
		select {
		case <-ctx.Done():
			return
		case <-request.Writers[0].(*HTTPResponseWriteCloser).CloseCh:
			return
		}
	}
}
