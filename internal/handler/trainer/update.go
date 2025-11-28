package trainer

import (
	"strconv"

	"treners_app/internal/domain"
	"treners_app/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type UpdateTrainerRequest struct {
	ID           int       `json:"id"`
	LastName     string    `json:"last_name" validate:"required"`
	FirstName    string    `json:"first_name" validate:"required"`
	MiddleName   string    `json:"middle_name"`
	Description  string    `json:"description"`
	Achievements []string  `json:"achievements,omitempty"`
	Titles       []string  `json:"titles,omitempty"`
	Contacts     []Contact `json:"contacts" validate:"required,min=1,dive"`
}

// обновляем тренера с учетом его ид - неизменяемый параметр
// меняем у тренера фио
// меняем у тренера титулы и прочее
// меняем у тренера контакты (на будущее)
func Update(repo *repository.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req UpdateTrainerRequest
		ID := c.Query("id")
		if ID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "query param 'id' is required",
			})
		}

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
		}

		// Начало транзакции
		tx, err := repo.BeginTx(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to start transaction",
			})
		}
		defer tx.Rollback(c.Context())

		idTrainer, err := strconv.Atoi(ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed trainer ID",
			})
		}

		// Изменение тренера
		if err = repo.UpdateTrainer(c.Context(), tx, trainer, idTrainer); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update trainer",
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
				if err = repo.CreateAchievement(c.Context(), tx, &achievement); err != nil {
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
				if err = repo.CreateTitle(c.Context(), tx, &title); err != nil {
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

		return c.Status(fiber.StatusAccepted).JSON(response)
	}
}
