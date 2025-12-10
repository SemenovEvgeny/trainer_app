package trainer

import (
	"context"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"treners_app/internal/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) ActivateTrainer(ctx context.Context, ID string) (domain.Trainer, error) {
	args := m.Called(ctx, ID)
	return args.Get(0).(domain.Trainer), args.Error(1)
}

func (m *MockRepository) BeginTx(ctx context.Context) (interface{}, error) {
	args := m.Called(ctx)
	return args.Get(0), args.Error(1)
}

func TestActivate(t *testing.T) {
	tests := []struct {
		name           string
		queryID        string
		mockTrainer    domain.Trainer
		mockError      error
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:    "successful activation",
			queryID: "1",
			mockTrainer: domain.Trainer{
				ID:          1,
				LastName:    "Иванов",
				FirstName:   "Иван",
				MiddleName:  "Иванович",
				Description: "Опытный тренер",
				IsActive:    true,
			},
			mockError:      nil,
			expectedStatus: fiber.StatusOK,
			expectedBody: map[string]interface{}{
				"id":          float64(1),
				"last_name":   "Иванов",
				"first_name":  "Иван",
				"middle_name": "Иванович",
				"description": "Опытный тренер",
				"is_active":   true,
			},
		},
		{
			name:           "missing id query param",
			queryID:        "",
			mockTrainer:    domain.Trainer{},
			mockError:      nil,
			expectedStatus: fiber.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "query param 'id' is required",
			},
		},
		{
			name:    "repository error",
			queryID: "1",
			mockTrainer: domain.Trainer{
				ID: 1,
			},
			mockError:      errors.New("database error"),
			expectedStatus: fiber.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"error": "failed to search trainers",
			},
		},
		{
			name:    "trainer not found",
			queryID: "999",
			mockTrainer: domain.Trainer{
				ID: 0,
			},
			mockError:      errors.New("trainer not found"),
			expectedStatus: fiber.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"error": "failed to search trainers",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			mockRepo := new(MockRepository)

			if tt.queryID != "" && tt.mockError == nil {
				mockRepo.On("ActivateTrainer", mock.Anything, tt.queryID).
					Return(tt.mockTrainer, nil)
			} else if tt.queryID != "" && tt.mockError != nil {
				mockRepo.On("ActivateTrainer", mock.Anything, tt.queryID).
					Return(tt.mockTrainer, tt.mockError)
			}

			app.Post("/activate", func(c *fiber.Ctx) error {
				ID := c.Query("id")
				if ID == "" {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "query param 'id' is required",
					})
				}

				trainer, err := mockRepo.ActivateTrainer(c.Context(), ID)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error": "failed to search trainers",
					})
				}

				return c.JSON(trainer)
			})

			req := httptest.NewRequest("POST", "/activate?id="+tt.queryID, nil)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var responseBody map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			assert.NoError(t, err)

			if tt.expectedBody != nil {
				for key, expectedValue := range tt.expectedBody {
					assert.Equal(t, expectedValue, responseBody[key], "mismatch for key: %s", key)
				}
			}

			if tt.queryID != "" {
				mockRepo.AssertExpectations(t)
			}
		})
	}
}
