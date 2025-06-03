package repository

import (
	"context"
	"fmt"

	"treners_app/internal/domain"
)

func (r *Repository) CreateTitle(ctx context.Context, title *domain.Title) error {
	query := `
		INSERT INTO title (trainer_id, value)
		VALUES ($1, $2)
		RETURNING id`

	err := r.conn.QueryRow(ctx, query,
		title.TrainerID,
		title.Value,
	).Scan(&title.ID)

	if err != nil {
		return fmt.Errorf("failed to create title: %w", err)
	}

	return nil
}
