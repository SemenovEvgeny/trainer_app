package trainer

import (
	"treners_app/internal/domain"
	"treners_app/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func GetByName(repo *repository.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Query("name")
		if name == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "query param 'name' is required",
			})
		}

		list, err := repo.GetTrainerByName(c.Context(), name)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to search trainers",
			})
		}

		if list == nil {
			list = domain.TrainerList{}
		}
		return c.JSON(list)
	}
}
