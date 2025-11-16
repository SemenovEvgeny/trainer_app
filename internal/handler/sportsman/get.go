package sportsman

import (
	"treners_app/internal/domain"
	"treners_app/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func GetByName(repo *repository.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		q := c.Query("q")
		if q == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "query param 'q' is required",
			})
		}

		list, err := repo.GetSportsmanByName(c.Context(), q)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to search sportsman",
			})
		}

		if list == nil {
			list = domain.SportsmanList{}
		}
		return c.JSON(list)
	}
}
