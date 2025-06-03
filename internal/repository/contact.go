package repository

import (
	"context"
	"fmt"

	"treners_app/internal/domain"
)

func (r *Repository) CreateContact(ctx context.Context, contact *domain.Contact) error {
	query := `
		INSERT INTO contact (trainer_id, type_id, contact)
		VALUES ($1, $2, $3)
		RETURNING id`

	err := r.conn.QueryRow(ctx, query,
		contact.TrainerID,
		contact.TypeID,
		contact.Contact,
	).Scan(&contact.ID)

	if err != nil {
		return fmt.Errorf("failed to create contact: %w", err)
	}

	return nil
}
