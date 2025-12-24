package repository

import (
	"context"
	"fmt"

	"treners_app/internal/domain"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) CreateAchievement(ctx context.Context, tx pgx.Tx, achievement *domain.Achievement) error {
	query := `
		INSERT INTO achievement (trainer_id, value)
		VALUES ($1, $2)
		RETURNING id`

	err := tx.QueryRow(ctx, query,
		achievement.TrainerID,
		achievement.Value,
	).Scan(&achievement.ID)

	if err != nil {
		return fmt.Errorf("failed to create achievement: %w", err)
	}

	return nil
}
