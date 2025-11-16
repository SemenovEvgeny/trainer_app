package repository

import (
	"context"
	"fmt"

	"treners_app/internal/domain"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) CreateTitle(ctx context.Context, tx pgx.Tx, title *domain.Title) error {
	query := `
		INSERT INTO title (trainer_id, value)
		VALUES ($1, $2)
		ON CONFLICT (trainer_id, value)
		DO UPDATE SET value = EXCLUDED.value
		RETURNING id`

	err := tx.QueryRow(ctx, query,
		title.TrainerID,
		title.Value,
	).Scan(&title.ID)

	if err != nil {
		return fmt.Errorf("failed to create title: %w", err)
	}

	return nil
}
