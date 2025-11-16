package http

import (
	"log/slog"
	"strconv"

	"treners_app/internal/handler/probe"
	"treners_app/internal/handler/sportsman"
	"treners_app/internal/handler/trainer"
	"treners_app/internal/logger"
	"treners_app/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type Service struct {
	repo *repository.Repository
	log  *slog.Logger
}

func NewService(repo *repository.Repository) (*Service, error) {
	log := logger.GetLogger()

	return &Service{
		repo: repo,
		log:  log,
	}, nil
}

func (s *Service) setupRoutes() *fiber.App {
	app := fiber.New()

	app.Get("/probe/readiness", probe.Readiness)
	app.Get("/probe/liveness", probe.Liveness)

	trainerGroup := app.Group("/api/trainers")
	trainerGroup.Post("/create", trainer.Create(s.repo))
	trainerGroup.Get("/getByName", trainer.GetByName(s.repo))
	// trainerGroup.Post("/update", trainer.Update(s.repo))
	// trainerGroup.Post("/delete", trainer.Delete(s.repo))

	sportsmanGroup := app.Group("/api/sportsman")
	sportsmanGroup.Post("/create", sportsman.Create(s.repo))
	sportsmanGroup.Get("/getByName", sportsman.GetByName(s.repo))

	return app
}

func (s *Service) ListenAndServe(port int) error {
	app := s.setupRoutes()

	err := app.Listen(":" + strconv.Itoa(port))
	if err != nil {
		s.log.Error("Server failed to start", logger.Err(err))
	}

	return err
}
