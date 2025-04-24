package probe

import "github.com/gofiber/fiber/v2"

func Liveness(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"ready": true,
	})
}
