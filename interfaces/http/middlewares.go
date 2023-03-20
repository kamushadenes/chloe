package http

import (
	"context"
	"github.com/rs/zerolog"
	"net/http"
)

func LoggingMiddleware(ctx context.Context) func(next http.Handler) http.Handler {
	logger := zerolog.Ctx(ctx)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := logger.WithContext(r.Context())
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
