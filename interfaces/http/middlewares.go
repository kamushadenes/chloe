package http

import (
	"context"
	"errors"
	"github.com/go-chi/render"
	"github.com/kamushadenes/chloe/memory"
	"github.com/rs/zerolog"
	"net/http"
	"strings"
)

type userCtxKey struct{}

func LoggingMiddleware(ctx context.Context) func(next http.Handler) http.Handler {
	logger := zerolog.Ctx(ctx)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := logger.WithContext(r.Context())
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// UserAPIKeyAuthMiddleware is a middleware that validates the user API key present in a Bearer token in the Authorization header
func UserAPIKeyAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")

		// Verify that the Authorization header is present and starts with "Bearer "
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			render.Render(w, r, ErrUnauthorized(errors.New("missing or invalid Authorization header")))
			return
		}

		// Extract the token from the Authorization header
		token := strings.TrimPrefix(authHeader, "Bearer ")

		apiKey, err := memory.GetAPIKey(token)
		if err != nil {
			render.Render(w, r, ErrUnauthorized(errors.New("invalid API key")))
			return
		}

		user := apiKey.User

		// Add the validated token to the request context
		ctx := context.WithValue(r.Context(), userCtxKey{}, user)

		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
