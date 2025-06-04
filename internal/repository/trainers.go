package repository

import (
	"context"
	"fmt"

	"treners_app/internal/domain"
)

func (r *Repository) CreateTrainer(ctx context.Context, trainer *domain.Trainer) error {
	query := `
		INSERT INTO trainer (last_name, first_name, middle_name, description, is_active)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	err := r.conn.QueryRow(ctx, query,
		trainer.LastName,
		trainer.FirstName,
		trainer.MiddleName,
		trainer.Description,
		trainer.IsActive,
	).Scan(&trainer.ID)

	if err != nil {
		return fmt.Errorf("failed to create trainer: %w", err)
	}

	return nil
}

func (r *Repository) GetTrainer(ctx context.Context, trainer *domain.Trainer) error {
	query := `
		INSERT INTO trainer (last_name, first_name, middle_name, description, is_active)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	err := r.conn.QueryRow(ctx, query,
		trainer.LastName,
		trainer.FirstName,
		trainer.MiddleName,
		trainer.Description,
		trainer.IsActive,
	).Scan(&trainer.ID)

	if err != nil {
		return fmt.Errorf("failed to create trainer: %w", err)
	}

	return nil
}

func (r *Repository) GetAllTrainers(ctx context.Context, trainer *domain.Trainer) error {
	query := `
		INSERT INTO trainer (last_name, first_name, middle_name, description, is_active)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	err := r.conn.QueryRow(ctx, query,
		trainer.LastName,
		trainer.FirstName,
		trainer.MiddleName,
		trainer.Description,
		trainer.IsActive,
	).Scan(&trainer.ID)

	if err != nil {
		return fmt.Errorf("failed to create trainer: %w", err)
	}

	return nil
}
