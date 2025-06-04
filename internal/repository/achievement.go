package repository

import (
	"context"
	"fmt"

	"treners_app/internal/domain"
)

func (r *Repository) CreateAchievement(ctx context.Context, achievement *domain.Achievement) error {
	query := `
		INSERT INTO achievement (trainer_id, value)
		VALUES ($1, $2)
		RETURNING id`

	err := r.conn.QueryRow(ctx, query,
		achievement.TrainerID,
		achievement.Value,
	).Scan(&achievement.ID)

	if err != nil {
		return fmt.Errorf("failed to create achievement: %w", err)
	}

	return nil
}
