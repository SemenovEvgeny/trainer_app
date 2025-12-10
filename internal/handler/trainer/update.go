package trainer

//
// import (
// 	"treners_app/internal/domain"
// 	"treners_app/internal/repository"
//
// 	"github.com/gofiber/fiber/v2"
// )
//
// type UpdateTrainerRequest struct {
// 	ID          int64  `json:"id"`
// 	LastName    string `json:"last_name" validate:"required"`
// 	FirstName   string `json:"first_name" validate:"required"`
// 	MiddleName  string `json:"middle_name"`
// 	Description string `json:"description"`
// }
//
// type UpdateTrainerResponse struct {
// 	Trainer *domain.Trainer `json:"trainer"`
// }
//
// // обновляем тренера с учетом его ид - неизменяемый параметр
// // меняем у тренера фио
// // меняем у тренера титулы и прочее
// // меняем у тренера контакты (на будущее)
// func UpdateTrainer(repo *repository.Repository) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		var req UpdateTrainerRequest
// 		ID := c.Query("id")
// 		if ID == "" {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": "query param 'id' is required",
// 			})
// 		}
//
// 		if err := c.BodyParser(&req); err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": "Invalid request body",
// 			})
// 		}
//
// 		// Валидация обязательных полей
// 		if req.LastName == "" {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": "Last name is required",
// 			})
// 		}
//
// 		if req.FirstName == "" {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": "First name is required",
// 			})
// 		}
//
// 		trainer := &domain.Trainer{
// 			LastName:    req.LastName,
// 			FirstName:   req.FirstName,
// 			MiddleName:  req.MiddleName,
// 			Description: req.Description,
// 		}
//
// 		err := repo.IsExistsTrainer(c.Context(), ID)
// 		if err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": "Trainer not found",
// 			})
// 		}
//
// 		err = repo.UpdateTrainer(c.Context(), trainer, ID)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": "failed to update trainer",
// 			})
// 		}
//
// 		response := &UpdateTrainerResponse{
// 			Trainer: trainer,
// 		}
//
// 		return c.Status(fiber.StatusAccepted).JSON(response)
// 	}
// }
