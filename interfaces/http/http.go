package http

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"sync"
	"time"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/utils"
)

func setMiddlewares(ctx context.Context, r *chi.Mux) {
	logger := logging.FromContext(ctx)

	for _, mid := range []func(http.Handler) http.Handler{
		LoggingMiddleware(ctx),
		// httplog.RequestLogger(*logger),
		middleware.Heartbeat("/ping"),
		middleware.RequestID,
		middleware.Recoverer,
		middleware.NoCache,
		middleware.Timeout(60 * time.Second),
		middleware.AllowContentType("application/json"),
		UserAPIKeyAuthMiddleware,
	} {
		logger.Debug().Str("middleware", utils.GetFunctionName(mid)).Msg("adding middleware")
		r.Use(mid)
	}
}

func setRoutes(ctx context.Context, r *chi.Mux) {
	logger := logging.FromContext(ctx)

	logger.Debug().Msg("adding routes")

	// Create a subdirectory filesystem that only contains the contents of the "web" folder
	webContentsFS, err := fs.Sub(WebFS, "web")
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create subdirectory filesystem for 'web'")
		return
	}

	// Custom file server that serves files from the embedded "web" folder contents
	fileServer := http.FileServer(http.FS(webContentsFS))

	r.Handle("/web/*", http.StripPrefix("/web", fileServer))

	r.Route("/api", func(r chi.Router) {
		r.Use(aiContext)
		r.Post("/complete", complete)
		r.Post("/generate", generate)
		r.Post("/tts", tts)
		r.Post("/forget", forget)
		r.Post("/action", action)
	})
}

func getRouter(ctx context.Context) *chi.Mux {
	logger := logging.FromContext(ctx)

	logger.Debug().Msg("getting router")

	r := chi.NewRouter()

	setMiddlewares(ctx, r)
	setRoutes(ctx, r)

	return r
}

func listen(ctx context.Context, server *http.Server, wg *sync.WaitGroup) {
	logger := logging.FromContext(ctx)

	logger.Debug().Msg("listening")

	wg.Add(1)
	defer wg.Done()

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Panic().Err(err).Msg("error in http listen")
	}
}

func Start(ctx context.Context) {
	logger := logging.GetLogger().With().Str("interface", "http").Logger()
	ctx = logger.WithContext(ctx)

	logger.Info().Str("host", config.HTTP.Host).Int("port", config.HTTP.Port).Msg("starting http interface")

	router := getRouter(ctx)

	server := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", config.HTTP.Host, config.HTTP.Port),
		Handler:           router,
		ReadHeaderTimeout: 60 * time.Second,
	}

	wg := sync.WaitGroup{}

	go listen(ctx, server, &wg)

	for {
		select {
		case <-ctx.Done():
			logger.Warn().Err(ctx.Err()).Msg("closing http interface")
			if err := server.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
				logger.Error().Err(err).Msg("error in http interface")
			}
			wg.Wait()
			return
		}
	}
}
