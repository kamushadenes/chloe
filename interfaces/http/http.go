package http

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/utils"
	"github.com/rs/zerolog"
	"net/http"
	"sync"
	"time"
)

func setMiddlewares(ctx context.Context, r *chi.Mux) {
	logger := zerolog.Ctx(ctx)

	for _, mid := range []func(http.Handler) http.Handler{
		LoggingMiddleware(ctx),
		//httplog.RequestLogger(*logger),
		middleware.Heartbeat("/ping"),
		middleware.RequestID,
		middleware.Recoverer,
		middleware.NoCache,
		middleware.Timeout(60 * time.Second),
		middleware.AllowContentType("application/json"),
	} {
		logger.Debug().Str("middleware", utils.GetFunctionName(mid)).Msg("adding middleware")
		r.Use(mid)
	}
}

func setRoutes(ctx context.Context, r *chi.Mux) {
	logger := zerolog.Ctx(ctx)
	logger.Debug().Msg("adding routes")

	r.Route("/api", func(r chi.Router) {
		r.Use(aiContext)
		r.Post("/complete", aiComplete)
		r.Post("/generate", aiGenerate)
		r.Post("/tts", aiTTS)
	})
}

func getRouter(ctx context.Context) *chi.Mux {
	logger := zerolog.Ctx(ctx)
	logger.Debug().Msg("getting router")

	r := chi.NewRouter()

	setMiddlewares(ctx, r)
	setRoutes(ctx, r)

	return r
}

func listen(ctx context.Context, server *http.Server, wg *sync.WaitGroup) {
	logger := zerolog.Ctx(ctx)

	logger.Debug().Msg("listening")

	wg.Add(1)
	defer wg.Done()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Panic().Err(err).Msg("error in http listen")
	}
}

func Start(ctx context.Context) {
	logger := zerolog.Ctx(ctx).With().Str("interface", "http").Logger()
	ctx = logger.WithContext(ctx)

	logger.Info().Msg("starting http interface")

	router := getRouter(ctx)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.HTTP.Host, config.HTTP.Port),
		Handler: router,
	}

	wg := sync.WaitGroup{}

	go listen(ctx, server, &wg)

	<-ctx.Done()

	if err := server.Shutdown(ctx); err != nil {
		logger.Panic().Err(err).Msg("error in http interface")
	}

	wg.Wait()
}
