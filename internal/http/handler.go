package http

import (
	"log/slog"
	"strconv"

	"treners_app/internal/handler/health"
	"treners_app/internal/logger"
	"treners_app/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type Service struct {
	repo   *repository.Repository
	log    *slog.Logger
	secret string
}

func NewService(repo *repository.Repository, log *slog.Logger, secret string) (*Service, error) {
	return &Service{
		repo:   repo,
		log:    log,
		secret: secret,
	}, nil
}

func (s *Service) SetupRoutes() *fiber.App {
	app := fiber.New()

	app.Get("/probe/readiness", health.ProbeReadiness)

	return app
}

func (s *Service) ListenAndServe(port int) error {
	app := s.SetupRoutes()

	err := app.Listen(":" + strconv.Itoa(port))
	if err != nil {
		s.log.Error("Server failed to start", logger.Err(err))
	}

	return err
}
