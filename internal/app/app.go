package app

import (
	"context"
	"log/slog"
	"os"

	"treners_app/internal/config"
	"treners_app/internal/http"
	"treners_app/internal/logger"
	"treners_app/internal/repository"
)

func NewApp() {

	ctx := context.Background()

	cfg := config.GetConfig()

	log := logger.GetLogger()
	log.Info("Logger initialized successfully")

	log.Info("Start on server",
		slog.String("env", cfg.Env),
	)

	err := startApp(ctx)
	if err != nil {
		log.Error("Error starting app: ", logger.Err(err))

		os.Exit(1)
	}
}

func startApp(ctx context.Context) (err error) {
	cfg := config.GetConfig()
	log := logger.GetLogger()

	var repo *repository.Repository

	repo, err = repository.NewRepository(ctx)
	if err != nil {
		log.Error("failed to initialize repository", logger.Err(err))

		return err
	}

	log.Info("Repository initialized successfully")

	var srv *http.Service

	srv, err = http.NewService(repo)
	if err != nil {
		log.Error("failed to initialize service", logger.Err(err))

		return err
	}

	log.Info("Service initialized successfully")

	log.Info("Starting server at", slog.Int("port", cfg.HTTPServer.Port))

	err = srv.ListenAndServe(cfg.HTTPServer.Port)
	if err != nil {
		log.Error("error in Listen abd Serve", logger.Err(err))

		return err
	}

	return nil
}
