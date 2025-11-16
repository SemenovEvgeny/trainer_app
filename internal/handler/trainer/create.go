package trainer

import (
	"treners_app/internal/domain"
	"treners_app/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type Contact struct {
	TypeID  int64  `json:"type_id" validate:"required"`
	Contact string `json:"contact" validate:"required"`
}

type CreateTrainerRequest struct {
	LastName     string    `json:"last_name" validate:"required"`
	FirstName    string    `json:"first_name" validate:"required"`
	MiddleName   string    `json:"middle_name"`
	Description  string    `json:"description"`
	IsActive     bool      `json:"is_active,omitempty"`
	Achievements []string  `json:"achievements,omitempty"`
	Titles       []string  `json:"titles,omitempty"`
	Contacts     []Contact `json:"contacts" validate:"required,min=1,dive"`
}

type CreateTrainerResponse struct {
	Trainer      *domain.Trainer      `json:"trainer"`
	Achievements []domain.Achievement `json:"achievements,omitempty"`
	Titles       []domain.Title       `json:"titles,omitempty"`
	Contacts     []domain.Contact     `json:"contacts,required"`
}

func Create(repo *repository.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateTrainerRequest
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

		if len(req.Contacts) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Contacts is required",
			})
		}

		for _, con := range req.Contacts {
			if con.TypeID == 0 || con.Contact == "" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "All contact fields are required",
				})
			}
		}

		trainer := &domain.Trainer{
			LastName:    req.LastName,
			FirstName:   req.FirstName,
			MiddleName:  req.MiddleName,
			Description: req.Description,
			IsActive:    req.IsActive,
		}

		// Начало транзакции
		tx, err := repo.BeginTx(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to start transaction",
			})
		}
		defer tx.Rollback(c.Context())

		// Создание тренера
		if err := repo.CreateTrainer(c.Context(), tx, trainer); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create trainer",
			})
		}

		response := &CreateTrainerResponse{
			Trainer: trainer,
		}

		// Создание достижений
		if len(req.Achievements) > 0 {
			achievements := make(domain.AchievementList, 0, len(req.Achievements))
			for _, value := range req.Achievements {
				achievement := domain.Achievement{
					TrainerID: trainer.ID,
					Value:     value,
				}
				if err := repo.CreateAchievement(c.Context(), tx, &achievement); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error": "Failed to create achievement",
					})
				}
				achievements = append(achievements, achievement)
			}
			response.Achievements = achievements
		}

		// Создание званий
		if len(req.Titles) > 0 {
			titles := make([]domain.Title, 0, len(req.Titles))
			for _, value := range req.Titles {
				title := domain.Title{
					TrainerID: trainer.ID,
					Value:     value,
				}
				if err := repo.CreateTitle(c.Context(), tx, &title); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error": "Failed to create title",
					})
				}
				titles = append(titles, title)
			}
			response.Titles = titles
		}

		// Создание контактов
		if len(req.Contacts) > 0 {
			contacts := make([]domain.Contact, 0, len(req.Contacts))
			for _, contact := range req.Contacts {
				newContact := domain.Contact{
					TrainerID: trainer.ID,
					TypeID:    contact.TypeID,
					Contact:   contact.Contact,
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
