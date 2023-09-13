package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/i18n"
	"github.com/kamushadenes/chloe/langchain/actions"
	"github.com/kamushadenes/chloe/langchain/chat_models"
	"github.com/kamushadenes/chloe/langchain/chat_models/messages"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
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

	if err := msg.Save(ctx); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	writer := NewHTTPResponseWriteCloser(w)

	chat := chat_models.NewChatWithDefaultModel(config.Chat.Provider, msg.User)

	_, err := chat.ChatStreamWithContext(ctx, writer, messages.UserMessage(params.Content))
	if err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
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
	if err := msg.Save(ctx); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	req := action_structs.NewActionRequest()
	req.Context = ctx
	req.Message = msg
	req.Action = "generate"
	req.Params["prompt"] = params.Prompt
	req.Writer = NewHTTPResponseWriteCloser(w)
	req.Count = config.Telegram.ImageCount
	req.SkipClose = true

	if err := actions.HandleAction(req); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
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
	if err := msg.Save(ctx); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	req := action_structs.NewActionRequest()
	req.Message = msg
	req.Context = ctx
	req.Action = "tts"
	req.Params["text"] = params.Content
	req.Writer = NewHTTPResponseWriteCloser(w)
	req.SkipClose = true

	if err := actions.HandleAction(req); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
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
	if err := msg.Save(ctx); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	req := action_structs.NewActionRequest()
	req.Context = ctx
	req.Message = msg
	req.Action = params.Action
	req.Params["text"] = params.Params
	req.Writer = &HTTPWriter{Writer: w}

	/*go func() {
		if err := structs.RunAction(req); err != nil {
			_ = render.Render(w, r, ErrInvalidRequest(err))
			return
		}
	}()*/

	for {
		select {
		case <-ctx.Done():
			return
		case <-req.Writer.(*HTTPWriter).CloseCh:
			return
		}
	}
}
