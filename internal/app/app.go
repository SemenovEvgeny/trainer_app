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

	cfg, log := initBase()

	err := startApp(ctx, cfg, log)
	if err != nil {
		log.Error("Error starting app: ", logger.Err(err))

		os.Exit(1)
	}

}

func initBase() (cfg *config.Config, log *slog.Logger) {
	cfg = config.NewConfig()

	log = logger.NewLogger(cfg.Env)
	log.Info("Logger initialized successfully")

	log.Info("Start on server",
		slog.String("env", cfg.Env),
	)

	return cfg, log
}

func startApp(ctx context.Context, cfg *config.Config, log *slog.Logger) (err error) {

	var repo *repository.Repository

	repo, err = repository.NewRepository(ctx, cfg.Storage)
	if err != nil {
		log.Error("failed to initialize repository", logger.Err(err))

		return err
	}

	log.Info("Repository initialized successfully")

	var srv *http.Service

	srv, err = http.NewService(repo, log)
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
