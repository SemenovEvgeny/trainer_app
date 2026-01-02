package repository

import (
	"context"
	"fmt"

	"treners_app/internal/domain"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) CreateContact(ctx context.Context, tx pgx.Tx, contact *domain.Contact) error {
	query := `
		INSERT INTO contact (trainer_id, sportsman_id, type_id, contact)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (COALESCE(trainer_id, sportsman_id), type_id, contact) 
		DO UPDATE SET
			contact = EXCLUDED.contact,
			updated_at = CURRENT_TIMESTAMP
		RETURNING id`

	var trainerID interface{} = nil
	var sportsmanID interface{} = nil

	if contact.TrainerID != 0 {
		trainerID = contact.TrainerID
	}
	if contact.SportsmanID != 0 {
		sportsmanID = contact.SportsmanID
	}

	err := tx.QueryRow(ctx, query,
		trainerID,
		sportsmanID,
		contact.TypeID,
		contact.Contact,
	).Scan(&contact.ID)

	if err != nil {
		return fmt.Errorf("failed to create contact: %w", err)
	}

	return nil
}

func (r *Repository) UpdateContact(ctx context.Context, tx pgx.Tx, contact *domain.Contact, ID string) error {
	var query string
	var err error

	if contact.TrainerID != 0 {
		query = `
			UPDATE contact 
			SET type_id = $1, contact = $2
			WHERE trainer_id = $3 AND id = $4
			RETURNING id, trainer_id, type_id, contact`

		err = tx.QueryRow(ctx, query,
			contact.TypeID,
			contact.Contact,
			contact.TrainerID,
			ID,
		).Scan(&contact.ID, &contact.TrainerID, &contact.TypeID, &contact.Contact)
	} else if contact.SportsmanID != 0 {
		query = `
			UPDATE contact 
			SET type_id = $1, contact = $2
			WHERE sportsman_id = $3 AND id = $4
			RETURNING id, sportsman_id, type_id, contact`

		err = tx.QueryRow(ctx, query,
			contact.TypeID,
			contact.Contact,
			contact.SportsmanID,
			ID,
		).Scan(&contact.ID, &contact.SportsmanID, &contact.TypeID, &contact.Contact)
	} else {
		return fmt.Errorf("either trainer_id or sportsman_id must be provided")
	}

	if err != nil {
		return fmt.Errorf("failed to update contact: %w", err)
	}

	return nil
}
