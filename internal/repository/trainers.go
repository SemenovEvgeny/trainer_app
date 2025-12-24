package repository

import (
	"context"
	"fmt"

	"treners_app/internal/domain"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) CreateTrainer(ctx context.Context, tx pgx.Tx, trainer *domain.Trainer) error {
	query := `
		INSERT INTO trainer (last_name, first_name, middle_name, description, is_active)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	err := tx.QueryRow(ctx, query,
		trainer.LastName,
		trainer.FirstName,
		trainer.MiddleName,
		trainer.Description,
		trainer.IsActive,
	).Scan(&trainer.ID)

	if err != nil {
		return fmt.Errorf("failed to create trainer (tx): %w", err)
	}

	return nil
}

func (r *Repository) GetTrainerByName(ctx context.Context, name string) (domain.TrainerList, error) {
	query := `
		SELECT id, last_name, first_name, middle_name, description, is_active
		FROM trainer
		WHERE first_name ILIKE '%' || $1 || '%' OR last_name ILIKE '%' || $1 || '%'
		ORDER BY id`

	rows, err := r.conn.Query(ctx, query, name)
	if err != nil {
		return nil, fmt.Errorf("failed to search trainers: %w", err)
	}
	defer rows.Close()

	var trainers domain.TrainerList
	for rows.Next() {
		var t domain.Trainer
		if err = rows.Scan(
			&t.ID,
			&t.LastName,
			&t.FirstName,
			&t.MiddleName,
			&t.Description,
			&t.IsActive,
		); err != nil {
			return nil, fmt.Errorf("failed to scan trainer: %w", err)
		}
		trainers = append(trainers, t)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", rows.Err())
	}

	return trainers, nil
}

func (r *Repository) GetAllTrainers(ctx context.Context) (domain.TrainerList, error) {
	query := `
		SELECT id, last_name, first_name, middle_name, description, is_active
		FROM trainer
		ORDER BY id`

	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query trainers: %w", err)
	}
	defer rows.Close()

	var trainers domain.TrainerList
	for rows.Next() {
		var t domain.Trainer
		if err = rows.Scan(
			&t.ID,
			&t.LastName,
			&t.FirstName,
			&t.MiddleName,
			&t.Description,
			&t.IsActive,
		); err != nil {
			return nil, fmt.Errorf("failed to scan trainer: %w", err)
		}
		trainers = append(trainers, t)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", rows.Err())
	}

	return trainers, nil
}
func (r *Repository) UpdateTrainer(ctx context.Context, t *domain.Trainer, ID string) error {
	query := `UPDATE trainer 
    SET last_name = $1, 
        first_name = $2, 
        middle_name = $3, 
        description = $4
    WHERE id = $5
    RETURNING id, last_name, first_name, middle_name, description, is_active`

	err := r.conn.QueryRow(ctx, query,
		t.LastName,
		t.FirstName,
		t.MiddleName,
		t.Description,
		ID,
	).Scan(
		&t.ID,
		&t.LastName,
		&t.FirstName,
		&t.MiddleName,
		&t.Description,
		&t.IsActive,
	)

	if err != nil {
		return fmt.Errorf("failed to update trainer with id %s: %w", ID, err)
	}

	return nil
}

func (r *Repository) DeleteTrainer(ctx context.Context, ID string) (domain.Trainer, error) {
	query := `UPDATE trainer 
	SET is_active = false
	WHERE id = $1
	RETURNING id, last_name, first_name, middle_name, description, is_active`

	var t domain.Trainer

	rows, err := r.conn.Query(ctx, query, ID)
	if err != nil {
		return t, fmt.Errorf("failed to search trainers: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&t.ID,
			&t.LastName,
			&t.FirstName,
			&t.MiddleName,
			&t.Description,
			&t.IsActive,
		); err != nil {
			return t, fmt.Errorf("failed to scan trainer: %w", err)
		}
	}
	if rows.Err() != nil {
		return t, fmt.Errorf("rows error: %w", rows.Err())
	}

	return t, nil
}

func (r *Repository) ActivateTrainer(ctx context.Context, ID string) (domain.Trainer, error) {
	query := `UPDATE trainer 
	SET is_active = true
	WHERE id = $1
	RETURNING id, last_name, first_name, middle_name, description, is_active`

	var t domain.Trainer

	rows, err := r.conn.Query(ctx, query, ID)
	if err != nil {
		return t, fmt.Errorf("failed to search trainers: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&t.ID,
			&t.LastName,
			&t.FirstName,
			&t.MiddleName,
			&t.Description,
			&t.IsActive,
		); err != nil {
			return t, fmt.Errorf("failed to scan trainer: %w", err)
		}
	}
	if rows.Err() != nil {
		return t, fmt.Errorf("rows error: %w", rows.Err())
	}

	return t, nil
}

func (r *Repository) IsExistsTrainer(ctx context.Context, ID string) error {
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM trainer WHERE id = $1)`
	err := r.conn.QueryRow(ctx, checkQuery, ID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check trainer existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("trainer with id %s not found", ID)
	}
	return nil
}

func (r *Repository) UpdateTrainerUserID(ctx context.Context, tx pgx.Tx, trainerID, userID int64) error {
	query := `UPDATE trainer SET user_id = $1 WHERE id = $2`
	_, err := tx.Exec(ctx, query, userID, trainerID)
	if err != nil {
		return fmt.Errorf("failed to update trainer user_id: %w", err)
	}
	return nil
}
