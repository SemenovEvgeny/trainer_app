package sportsman

import (
	"treners_app/internal/domain"
	"treners_app/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type CreateSportsmanRequest struct {
	LastName   string `json:"last_name" validate:"required"`
	FirstName  string `json:"first_name" validate:"required"`
	MiddleName string `json:"middle_name"`
	IsActive   bool   `json:"is_active,omitempty"`
}

type CreateSportsmanResponse struct {
	Sportsman *domain.Sportsman `json:"sportsman"`
}

func Create(repo *repository.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateSportsmanRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Валидация обязательных полей
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

		sportsman := &domain.Sportsman{
			LastName:   req.LastName,
			FirstName:  req.FirstName,
			MiddleName: req.MiddleName,
			IsActive:   req.IsActive,
		}

		// Начало транзакции
		tx, err := repo.BeginTx(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to start transaction",
			})
		}
		defer tx.Rollback(c.Context())

		// Создание спортсмена (клиента)
		if err = repo.CreateSportsman(c.Context(), tx, sportsman); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create sportsman",
			})
		}

		response := &CreateSportsmanResponse{
			Sportsman: sportsman,
		}

		// Подтверждение транзакции
		if err = tx.Commit(c.Context()); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to commit transaction",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(response)
	}
}
