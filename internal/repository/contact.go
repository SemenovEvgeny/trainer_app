package repository

import (
	"context"
	"fmt"

	"treners_app/internal/domain"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) CreateContact(ctx context.Context, tx pgx.Tx, contact *domain.Contact) error {
	query := `
		INSERT INTO contact (trainer_id, type_id, contact)
		VALUES ($1, $2, $3)
		RETURNING id`

	err := tx.QueryRow(ctx, query,
		contact.TrainerID,
		contact.TypeID,
		contact.Contact,
	).Scan(&contact.ID)

	if err != nil {
		return fmt.Errorf("failed to create contact: %w", err)
	}

	return nil
}

func (r *Repository) UpdateContact(ctx context.Context, tx pgx.Tx, contact *domain.Contact, ID string) error {
	query := `
		UPDATE contact 
		SET (type_id, contact)
		VALUES ($1, $2)
		WHERE trainer_id = $3 AND id = $4
		RETURNING trainer_id, type_id, contact`

	err := tx.QueryRow(ctx, query,
		contact.TrainerID,
		contact.TypeID,
		contact.Contact,
	).Scan(&contact.ID)

	if err != nil {
		return fmt.Errorf("failed to create contact: %w", err)
	}

	return nil
}
