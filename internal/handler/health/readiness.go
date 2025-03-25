package health

import "github.com/gofiber/fiber/v2"

func ProbeReadiness(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"ready": true,
	})
}
