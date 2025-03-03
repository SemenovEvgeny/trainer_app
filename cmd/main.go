package main

import (
	"context"
	"log/slog"
	"treners_app/internal/config"
	"treners_app/internal/handler"
	"treners_app/internal/logger"
	"treners_app/internal/repository"
)

func main() {
	ctx := context.Background()

	cfg := config.NewConfig()

	log := logger.NewLogger(cfg.Env)

	log.Info("Logger initialized successfully")

	log.Info("Start on server",
		slog.String("env", cfg.Env),
	)

	var repo *repository.Repository

	repo, err := repository.NewRepository(ctx, cfg.Storage, log)
	if err != nil {
		log.Error("failed to initialize repository", logger.Err(err))
	}

	log.Info("Repository initialized successfully")

	var srv *handler.Service

	var secret string

	srv, err = handler.NewService(repo, log, secret)
	if err != nil {
		log.Error("failed to initialize service", logger.Err(err))
	}

	log.Info("Service initialized successfully")

	log.Info("Starting server at", slog.Int("port", cfg.HTTPServer.Port))

	err = srv.ListenAndServe(cfg.HTTPServer.Port)
	if err != nil {
		log.Error("error in Listen abd Serve", logger.Err(err))
	}
}
