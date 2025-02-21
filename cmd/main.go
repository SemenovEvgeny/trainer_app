package main

import (
	"context"
	"fmt"
	"log"
	"treners_app/internal/config"
	"treners_app/internal/handler"
	"treners_app/internal/logger"
	"treners_app/internal/repository"

	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	//TODO: init config
	cfg := config.NewConfig()

	log.Println("Configuration loaded successfully")

	//TODO: init logger
	log, err := logger.NewLogger(cfg.Env)
	if err != nil {
		fmt.Errorf("failed to initialize logger: %w", err)
	}
	defer log.Sync()

	log.Info("Logger initialized successfully")

	log.Info("Start on server",
		zap.String("env", cfg.Env),
	)

	//TODO: init db
	var repo *repository.Repository

	repo, err = repository.NewRepository(ctx, cfg.Storage, log)
	if err != nil {
		log.Error(fmt.Sprintf("failed to initialize repository: %w", err))
	}

	log.Info("Repository initialized successfully")

	//TODO: init middleware

	var srv *handler.Service

	var secret string //временная заглушка

	srv, err = handler.NewService(repo, log, secret)
	if err != nil {
		log.Error(fmt.Sprintf("failed to initialize service: %w", err))
	}

	log.Info("Service initialized successfully")

	//TODO: star server
	log.Info("Starting server at", zap.Int("port", cfg.HTTPServer.Port))

	err = srv.ListenAndServe(cfg.HTTPServer.Port)
	if err != nil {
		log.Error(fmt.Sprintf("error in Listen abd Serve: %w", err))
	}

	//app.Get("/probe/readiness", func(c *fiber.Ctx) error {
	//	return c.JSON(fiber.Map{
	//		"ready": true,
	//	})
	//})
	//
	//app.Listen(":3000")
}
