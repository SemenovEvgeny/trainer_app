package auth

import (
	"treners_app/internal/repository"
	"treners_app/internal/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(repo *repository.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req UserLoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		// Валидация email
		if req.Email == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Email is required",
			})
		}

		// Валидация пароля
		if len(req.Password) < 6 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Password must be at least 6 characters",
			})
		}

		// Получение пользователя из базы данных
		user, err := repo.GetUserByEmail(c.Context(), req.Email)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid credentials",
			})
		}

		// Проверка пароля
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid credentials",
			})
		}
		// Генерация JWT токена
		token, err := utils.GenerateToken(user.ID, user.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to generate token: " + err.Error(),
			})
		}

		return c.JSON(fiber.Map{"token": token})
	}
}
