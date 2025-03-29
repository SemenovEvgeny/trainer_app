package probe

import "github.com/gofiber/fiber/v2"

func Readiness(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"ready": true,
	})
}

// TODO: изучить все эндпоинты (startup загуглить)
