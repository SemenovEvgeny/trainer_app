package middleware

import (
	"strings"

	"treners_app/internal/utils"

	"github.com/gofiber/fiber/v2"
)

const (
	UserIDKey    = "userID"
	UserEmailKey = "userEmail"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header required",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Authorization header format. Expected: Bearer <token>",
			})
		}

		tokenString := parts[1]
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		c.Locals(UserIDKey, claims.UserID)
		c.Locals(UserEmailKey, claims.Email)

		return c.Next()
	}
}

func RequireAuth() fiber.Handler {
	return AuthMiddleware()
}

func GetUserID(c *fiber.Ctx) (int64, bool) {
	userID, ok := c.Locals(UserIDKey).(int64)
	return userID, ok
}

func GetUserEmail(c *fiber.Ctx) (string, bool) {
	email, ok := c.Locals(UserEmailKey).(string)
	return email, ok
}

func RequireRole(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := GetUserID(c)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not authenticated",
			})
		}
		c.Locals(UserIDKey, userID)

		return c.Next()
	}
}
