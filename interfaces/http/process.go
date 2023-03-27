package http

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/i18n"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/structs"
	"io"
	"net/http"
)

var msgCtxKey = struct{}{}

func aiContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqID := middleware.GetReqID(ctx)

		if reqID != "" {
			msg := memory.NewMessage(reqID, "http")
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

			ctx = context.WithValue(ctx, msgCtxKey, msg)

		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func aiComplete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	msg := ctx.Value(msgCtxKey).(*memory.Message)

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

	msg.Content = params.Content
	channels.IncomingMessagesCh <- msg
	if err := <-msg.ErrorCh; err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	request := structs.NewCompletionRequest()
	request.Message = msg
	request.ID = msg.ExternalID
	request.Context = ctx

	if request.Mode == "" {
		request.Mode = request.GetMessage().User.Mode
	}

	request.Args = params.Args

	request.Writer = utils.NewHTTPResponseWriteCloser(w)

	channels.CompletionRequestsCh <- request

	for {
		select {
		case <-ctx.Done():
			return
		case <-request.Writer.(*utils.HTTPResponseWriteCloser).CloseCh:
			return
		}
	}
}

func aiGenerate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	msg := ctx.Value(msgCtxKey).(*memory.Message)

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

	msg.Content = params.Prompt
	channels.IncomingMessagesCh <- msg
	if err := <-msg.ErrorCh; err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	req := structs.NewActionRequest()
	req.ID = msg.ExternalID
	req.Context = ctx
	req.Action = "image"
	req.Params = params.Prompt
	req.Message = msg
	req.Writers = []io.WriteCloser{utils.NewHTTPResponseWriteCloser(w)}

	channels.ActionRequestsCh <- req

	for {
		select {
		case <-ctx.Done():
			return
		case <-req.Writers[0].(*utils.HTTPResponseWriteCloser).CloseCh:
			return
		}
	}
}

func aiTTS(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	msg := ctx.Value(msgCtxKey).(*memory.Message)

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

	msg.Content = params.Content
	channels.IncomingMessagesCh <- msg
	if err := <-msg.ErrorCh; err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	req := structs.NewActionRequest()
	req.ID = msg.ExternalID
	req.Context = ctx
	req.Action = "tts"
	req.Params = params.Content
	req.Message = msg
	req.Writers = []io.WriteCloser{utils.NewHTTPResponseWriteCloser(w)}

	channels.ActionRequestsCh <- req

	for {
		select {
		case <-ctx.Done():
			return
		case <-req.Writers[0].(*utils.HTTPResponseWriteCloser).CloseCh:
			return
		}
	}
}

func aiForget(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	msg := ctx.Value(msgCtxKey).(*memory.Message)

	if err := msg.User.DeleteMessages(ctx); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	if _, err := w.Write([]byte(i18n.GetForgetText())); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
}
