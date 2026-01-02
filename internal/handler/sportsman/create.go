package sportsman

import (
	"treners_app/internal/domain"
	"treners_app/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type Contact struct {
	TypeID  int64  `json:"type_id" validate:"required"`
	Contact string `json:"contact" validate:"required"`
}

type CreateSportsmanRequest struct {
	LastName   string    `json:"last_name" validate:"required"`
	FirstName  string    `json:"first_name" validate:"required"`
	MiddleName string    `json:"middle_name"`
	IsActive   bool      `json:"is_active,omitempty"`
	Contacts   []Contact `json:"contacts" validate:"required,min=1,dive"`
}

type CreateSportsmanResponse struct {
	Sportsman *domain.Sportsman `json:"sportsman"`
	Contacts  []domain.Contact  `json:"contacts"`
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

		// Создание контактов
		if len(req.Contacts) > 0 {
			contacts := make([]domain.Contact, 0, len(req.Contacts))
			for _, contact := range req.Contacts {
				newContact := domain.Contact{
					SportsmanID: sportsman.ID,
					TypeID:      contact.TypeID,
					Contact:     contact.Contact,
				}
				if err = repo.CreateContact(c.Context(), tx, &newContact); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error": "Failed to create contact",
					})
				}
				contacts = append(contacts, newContact)
			}
			response.Contacts = contacts
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
