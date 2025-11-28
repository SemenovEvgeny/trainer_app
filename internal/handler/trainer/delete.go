package trainer

import (
	"treners_app/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type DeleteTrainerRequest struct {
	ID       int  `json:"id"`
	IsActive bool `json:"is_active,omitempty"`
}

// при удалении тренера происходит перевод isActive в false

func Delete(repo *repository.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID := c.Query("id")
		if ID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "query param 'id' is required",
			})
		}

		t, err := repo.DeleteTrainer(c.Context(), ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to search trainers",
			})
		}

		return c.JSON(t)
	}
}
