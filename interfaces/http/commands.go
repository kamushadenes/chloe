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

			user := ctx.Value(userCtxKey{}).(*memory.User)

			msg.User = user

			ctx = context.WithValue(ctx, msgCtxKey, msg)

		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func complete(w http.ResponseWriter, r *http.Request) {
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

	msg.SetContent(params.Content)

	if err := channels.RegisterIncomingMessage(msg); err != nil {
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

	go func() {
		if err := channels.RunCompletion(request); err != nil {
			_ = render.Render(w, r, ErrInvalidRequest(err))
			return
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-request.Writer.(*utils.HTTPResponseWriteCloser).CloseCh:
			return
		}
	}
}

func generate(w http.ResponseWriter, r *http.Request) {
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

	msg.SetContent(params.Prompt)
	if err := channels.RegisterIncomingMessage(msg); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	req := structs.NewActionRequest()
	req.ID = msg.ExternalID
	req.Context = ctx
	req.Action = "generate"
	req.Params = params.Prompt
	req.Message = msg
	req.Writer = utils.NewHTTPResponseWriteCloser(w)

	go func() {
		if err := channels.RunAction(req); err != nil {
			_ = render.Render(w, r, ErrInvalidRequest(err))
			return
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-req.Writer.(*utils.HTTPResponseWriteCloser).CloseCh:
			return
		}
	}
}

func tts(w http.ResponseWriter, r *http.Request) {
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

	msg.SetContent(params.Content)
	if err := channels.RegisterIncomingMessage(msg); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	req := structs.NewActionRequest()
	req.ID = msg.ExternalID
	req.Context = ctx
	req.Action = "tts"
	req.Params = params.Content
	req.Message = msg
	req.Writer = utils.NewHTTPResponseWriteCloser(w)

	go func() {
		if err := channels.RunAction(req); err != nil {
			_ = render.Render(w, r, ErrInvalidRequest(err))
			return
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-req.Writer.(*utils.HTTPResponseWriteCloser).CloseCh:
			return
		}
	}
}

func forget(w http.ResponseWriter, r *http.Request) {
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

func action(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	msg := ctx.Value(msgCtxKey).(*memory.Message)

	var params = struct {
		Action string `json:"action"`
		Params string `json:"params"`
	}{}
	if err := parseFromRequest(r, &params); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if params.Action == "" {
		_ = render.Render(w, r, ErrInvalidRequest(fmt.Errorf("action is required")))
		return
	}
	if params.Params == "" {
		_ = render.Render(w, r, ErrInvalidRequest(fmt.Errorf("params is required")))
		return
	}

	msg.SetContent(fmt.Sprintf("%s %s", params.Action, params.Params))
	if err := channels.RegisterIncomingMessage(msg); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	req := structs.NewActionRequest()
	req.Context = ctx
	req.Message = msg
	req.Action = params.Action
	req.Params = params.Params
	req.Thought = fmt.Sprintf("User wants to run action %s", params.Action)
	req.Writer = &utils.HTTPResponseWriteCloser{Writer: w}

	go func() {
		if err := channels.RunAction(req); err != nil {
			_ = render.Render(w, r, ErrInvalidRequest(err))
			return
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-req.Writer.(*utils.HTTPResponseWriteCloser).CloseCh:
			return
		}
	}
}
