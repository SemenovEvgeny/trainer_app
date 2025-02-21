package handler

import (
	"strconv"
	"treners_app/internal/repository"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Service struct {
	repo   *repository.Repository
	log    *zap.Logger
	secret string
}

func NewService(repo *repository.Repository, log *zap.Logger, secret string) (*Service, error) {
	return &Service{
		repo:   repo,
		log:    log,
		secret: secret,
	}, nil
}

func (s *Service) SetupRoutes() *fiber.App {
	app := fiber.New()

	app.Get("/probe/readiness", probeReadiness)

	return app
}

func (s *Service) ListenAndServe(port int) error {
	app := s.SetupRoutes()

	err := app.Listen(":" + strconv.Itoa(port))
	if err != nil {
		s.log.Error("Server failed to start", zap.Error(err))
	}

	return err
}

func probeReadiness(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"ready": true,
	})
}
