package http

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"io"
	"net/http"
)

func aiContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqId := middleware.GetReqID(ctx)

		if reqId != "" {
			msg := memory.NewMessage(reqId, "http")
			msg.Role = "user"
			msg.Source.HTTP = &memory.HTTPMessageSource{
				Request: r,
			}

			user, err := memory.GetUserByExternalID(ctx, "http", "http")
			if err != nil {
				user, err = memory.NewUser(ctx, "User", "HTTP", "http")
				if err != nil {
					_ = render.Render(w, r, ErrInvalidRequest(err))
					return
				}
				err = user.AddExternalID(ctx, "http", "http")
				if err != nil {
					_ = render.Render(w, r, ErrInvalidRequest(err))
					return
				}
			}

			msg.User = user

			ctx = context.WithValue(ctx, "msg", msg)

			channels.IncomingMessagesCh <- msg
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func aiComplete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	msg := ctx.Value("msg").(*memory.Message)

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

	request := structs.NewCompletionRequest()
	request.ID = msg.ExternalID
	request.Context = ctx

	request.User = msg.User

	if request.Mode == "" {
		request.Mode = request.User.Mode
	}

	if err := msg.User.AddExternalID(ctx, "1", "http"); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	request.Message = msg
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
	msg := ctx.Value("msg").(*memory.Message)

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

	request := structs.NewGenerationRequest()
	request.ID = msg.ExternalID
	request.Context = ctx

	request.User = msg.User

	if err := msg.User.AddExternalID(ctx, "1", "http"); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
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
	msg := ctx.Value("msg").(*memory.Message)

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

	request := structs.NewTTSRequest()
	request.ID = msg.ExternalID
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
