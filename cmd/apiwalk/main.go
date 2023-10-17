package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/aichelnokov/apiwalk/internal/config"
	"github.com/aichelnokov/apiwalk/internal/lib/logger/sl"
	"github.com/aichelnokov/apiwalk/internal/routes"
	"github.com/aichelnokov/apiwalk/internal/storage/mysql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// Init config
	cfg := config.MustLoad()

	// Init logger
	logger := setupLogger(cfg.Env)
	logger = logger.With(slog.String("env", cfg.Env))
	logger.Debug("logger debug mode enabled")

	// Init storage
	_, err := mysql.New(cfg.DBConfig)
	if err != nil {
		logger.Error("failed to initialize storage", sl.Err(err))
	}

	// Init router
	r := chi.NewRouter()
  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	logger.Info("initializing server", slog.String("address", cfg.HTTPServer.Host + ":" + cfg.HTTPServer.Port))
	
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Walk " + cfg.ApiConfig.Version))
	})
	routes.Walk(r)
	
	// Run
	http.ListenAndServe(":" + cfg.HTTPServer.Port, r)
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
		case envLocal:
			logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		case envDev:
			logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		case envProd:
			logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}