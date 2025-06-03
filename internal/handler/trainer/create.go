package trainer

import (
	"treners_app/internal/domain"
	"treners_app/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type CreateTrainerRequest struct {
	LastName    string `json:"last_name" validate:"required"`
	FirstName   string `json:"first_name" validate:"required"`
	MiddleName  string `json:"middle_name"`
	Description string `json:"description"`
	IsActive    *bool  `json:"is_active,omitempty"`
}

func Create(repo *repository.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateTrainerRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		if req.LastName == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Last name is required",
			})
		}

		if req.FirstName == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "First name is required",
			})
		}

		if req.IsActive == nil {
			defaultActive := true
			req.IsActive = &defaultActive
		}

		trainer := &domain.Trainer{
			LastName:    req.LastName,
			FirstName:   req.FirstName,
			MiddleName:  req.MiddleName,
			Description: req.Description,
			IsActive:    *req.IsActive,
		}

		if err := repo.CreateTrainer(c.Context(), trainer); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create trainer",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(trainer)
	}
}
