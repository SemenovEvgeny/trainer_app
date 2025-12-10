package trainer

import (
	"treners_app/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type ActivateTrainerRequest struct {
	ID       int  `json:"id"`
	IsActive bool `json:"is_active,omitempty"`
}

// при активации тренера происходит перевод isActive в true

func Activate(repo *repository.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID := c.Query("id")
		if ID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "query param 'id' is required",
			})
		}

		t, err := repo.ActivateTrainer(c.Context(), ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to search trainers",
			})
		}

		return c.JSON(t)
	}
}
