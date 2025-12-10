package repository

import (
	"context"
	"fmt"

	"treners_app/internal/domain"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) CreateSportsman(ctx context.Context, tx pgx.Tx, sportsman *domain.Sportsman) error {
	query := `
		INSERT INTO client (last_name, first_name, middle_name, is_active)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (last_name, first_name, middle_name)
		DO UPDATE SET 
		    last_name = EXCLUDED.last_name
		    first_name = EXCLUDED.first_name
		    middle_name = EXCLUDED.middle_name
		RETURNING id;`

	err := tx.QueryRow(ctx, query,
		sportsman.LastName,
		sportsman.FirstName,
		sportsman.MiddleName,
		sportsman.IsActive,
	).Scan(&sportsman.ID)

	if err != nil {
		return fmt.Errorf("failed to create sportsman (tx): %w", err)
	}

	return nil
}

func (r *Repository) GetSportsmanByName(ctx context.Context, q string) (domain.SportsmanList, error) {
	query := `
		SELECT id, last_name, first_name, middle_name, is_active
		FROM client
		WHERE first_name ILIKE '%' || $1 || '%' OR last_name ILIKE '%' || $1 || '%'
		ORDER BY id`

	rows, err := r.conn.Query(ctx, query, q)
	if err != nil {
		return nil, fmt.Errorf("failed to search sportsman: %w", err)
	}
	defer rows.Close()

	var sportsman domain.SportsmanList
	for rows.Next() {
		var c domain.Sportsman
		if err := rows.Scan(
			&c.ID,
			&c.LastName,
			&c.FirstName,
			&c.MiddleName,
			&c.IsActive,
		); err != nil {
			return nil, fmt.Errorf("failed to scan sportsman: %w", err)
		}
		sportsman = append(sportsman, c)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", rows.Err())
	}

	return sportsman, nil
}

func (r *Repository) GetAllClient(ctx context.Context) (domain.SportsmanList, error) {
	query := `
		SELECT id, last_name, first_name, middle_name, description, is_active
		FROM client
		ORDER BY id`

	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query sportsman: %w", err)
	}
	defer rows.Close()

	var sportsman domain.SportsmanList
	for rows.Next() {
		var s domain.Sportsman
		if err := rows.Scan(
			&s.ID,
			&s.LastName,
			&s.FirstName,
			&s.MiddleName,
			&s.IsActive,
		); err != nil {
			return nil, fmt.Errorf("failed to scan sportsman: %w", err)
		}
		sportsman = append(sportsman, s)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", rows.Err())
	}

	return sportsman, nil
}
