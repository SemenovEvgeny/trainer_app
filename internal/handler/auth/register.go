package auth

import (
	"treners_app/internal/domain"
	"treners_app/internal/repository"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Contact struct {
	TypeID  int64  `json:"type_id" validate:"required"`
	Contact string `json:"contact" validate:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required,oneof=admin trainer sportsman"`

	// Данные для тренера (если role = trainer)
	TrainerData *TrainerData `json:"trainer_data,omitempty"`

	// Данные для спортсмена (если role = sportsman)
	SportsmanData *SportsmanData `json:"sportsman_data,omitempty"`
}

type TrainerData struct {
	LastName     string    `json:"last_name" validate:"required"`
	FirstName    string    `json:"first_name" validate:"required"`
	MiddleName   string    `json:"middle_name"`
	Description  string    `json:"description"`
	IsActive     bool      `json:"is_active"`
	Achievements []string  `json:"achievements,omitempty"`
	SportIDs     []int64   `json:"sport_ids,omitempty"` // ID видов спорта
	Contacts     []Contact `json:"contacts" validate:"required,min=1,dive"`
}

type SportsmanData struct {
	LastName   string    `json:"last_name" validate:"required"`
	FirstName  string    `json:"first_name" validate:"required"`
	MiddleName string    `json:"middle_name"`
	IsActive   bool      `json:"is_active,omitempty"`
	Contacts   []Contact `json:"contacts" validate:"required,min=1,dive"`
}

type RegisterResponse struct {
	User       *domain.User       `json:"user"`
	Trainer    *domain.Trainer    `json:"trainer,omitempty"`
	Sportsman  *domain.Sportsman  `json:"sportsman,omitempty"`
	SportTypes []domain.SportType `json:"sport_types,omitempty"`
	Contacts   []domain.Contact   `json:"contacts,omitempty"`
	Message    string             `json:"message"`
}

func Register(repo *repository.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req RegisterRequest
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

		// Валидация роли
		if req.Role != "admin" && req.Role != "trainer" && req.Role != "sportsman" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Role must be one of: admin, trainer, sportsman",
			})
		}

		// Проверка наличия данных для тренера/спортсмена
		if req.Role == "trainer" && req.TrainerData == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Trainer data is required for trainer role",
			})
		}

		if req.Role == "sportsman" && req.SportsmanData == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Sportsman data is required for sportsman role",
			})
		}

		// Проверка существования пользователя
		existingUser, _ := repo.GetUserByEmail(c.Context(), req.Email)
		if existingUser != nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "User with this email already exists",
			})
		}

		// Получение роли
		role, err := repo.GetRoleByValue(c.Context(), req.Role)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get role",
			})
		}

		// Хеширование пароля
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to hash password",
			})
		}

		// Начало транзакции
		tx, err := repo.BeginTx(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to start transaction",
			})
		}
		defer tx.Rollback(c.Context())

		// Создание пользователя
		user := &domain.User{
			Email:        req.Email,
			PasswordHash: string(passwordHash),
			RoleID:       role.ID,
			Role:         role,
		}

		if err = repo.CreateUser(c.Context(), tx, user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create user",
			})
		}

		response := &RegisterResponse{
			User:    user,
			Message: "User registered successfully",
		}

		// Создание тренера, если роль trainer
		if req.Role == "trainer" && req.TrainerData != nil {
			trainer := &domain.Trainer{
				LastName:    req.TrainerData.LastName,
				FirstName:   req.TrainerData.FirstName,
				MiddleName:  req.TrainerData.MiddleName,
				Description: req.TrainerData.Description,
				IsActive:    req.TrainerData.IsActive,
			}

			// Обновляем trainer с user_id после создания
			if err = repo.CreateTrainer(c.Context(), tx, trainer); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to create trainer",
				})
			}

			// Связываем тренера с пользователем
			if err = repo.UpdateTrainerUserID(c.Context(), tx, trainer.ID, user.ID); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to link trainer to user",
				})
			}

			// // Создание достижений
			// if len(req.TrainerData.Achievements) > 0 {
			// 	for _, value := range req.TrainerData.Achievements {
			// 		achievement := domain.Achievement{
			// 			TrainerID: trainer.ID,
			// 			Value:     value,
			// 		}
			// 		if err := repo.CreateAchievement(c.Context(), tx, &achievement); err != nil {
			// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			// 				"error": "Failed to create achievement",
			// 			})
			// 		}
			// 	}
			// }
			// Добавление видов спорта
			if len(req.TrainerData.SportIDs) > 0 {
				if err := repo.AddSportTypesToTrainer(c.Context(), tx, trainer.ID, req.TrainerData.SportIDs); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error": "Failed to add sport types to trainer",
					})
				}
				// Получаем добавленные виды спорта для ответа
				sportTypes, err := repo.GetSportTypesByTrainerIDTx(c.Context(), tx, trainer.ID)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error": "Failed to get sport types",
					})
				}
				response.SportTypes = sportTypes
			}

			// Создание контактов
			if len(req.TrainerData.Contacts) > 0 {
				contacts := make([]domain.Contact, 0, len(req.TrainerData.Contacts))
				for _, contact := range req.TrainerData.Contacts {
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

			response.Trainer = trainer
		}

		// Создание спортсмена, если роль sportsman
		if req.Role == "sportsman" && req.SportsmanData != nil {
			sportsman := &domain.Sportsman{
				LastName:   req.SportsmanData.LastName,
				FirstName:  req.SportsmanData.FirstName,
				MiddleName: req.SportsmanData.MiddleName,
				IsActive:   req.SportsmanData.IsActive,
			}

			if err = repo.CreateSportsman(c.Context(), tx, sportsman); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to create sportsman",
				})
			}

			// Связываем спортсмена с пользователем
			if err = repo.UpdateSportsmanUserID(c.Context(), tx, sportsman.ID, user.ID); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to link sportsman to user",
				})
			}

			// Создание контактов
			if len(req.SportsmanData.Contacts) > 0 {
				contacts := make([]domain.Contact, 0, len(req.SportsmanData.Contacts))
				for _, contact := range req.SportsmanData.Contacts {
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

			response.Sportsman = sportsman
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
