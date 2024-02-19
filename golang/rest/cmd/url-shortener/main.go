package main

import (
	"log/slog"
	"os"

	"github.com/4aykovski/learning/golang/rest/internal/config"
	"github.com/4aykovski/learning/golang/rest/internal/database/Postgres"
	mwLogger "github.com/4aykovski/learning/golang/rest/internal/http-server/middleware/logger"
	"github.com/4aykovski/learning/golang/rest/internal/lib/logger/slogHelper"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: разобраться с prettyslog
	// init config: cleanenv
	cfg := config.MustLoad()

	// init logger: slog ? grafana ? kibana ? grep

	log := setupLogger(cfg.Env)
	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// init db: postgres

	pq, err := Postgres.New(cfg.Postgres)
	if err != nil {
		log.Error("failed to init postgres database", slogHelper.Err(err))
		os.Exit(1)
	}

	repo := Postgres.NewUserRepository(pq)

	// init router: chi, "chi render"

	router := chi.NewRouter()

	// middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {

	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
