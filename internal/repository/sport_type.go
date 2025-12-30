package repository

import (
	"context"
	"fmt"

	"treners_app/internal/domain"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetAllSportTypes(ctx context.Context) (domain.SportTypeList, error) {
	query := `
		SELECT id, name
		FROM sport_type
		ORDER BY name`

	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query sport types: %w", err)
	}
	defer rows.Close()

	var sportTypes domain.SportTypeList
	for rows.Next() {
		var st domain.SportType
		if err := rows.Scan(&st.ID, &st.Name); err != nil {
			return nil, fmt.Errorf("failed to scan sport type: %w", err)
		}
		sportTypes = append(sportTypes, st)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", rows.Err())
	}

	return sportTypes, nil
}

func (r *Repository) GetSportTypesByTrainerID(ctx context.Context, trainerID int64) (domain.SportTypeList, error) {
	query := `
		SELECT st.id, st.name
		FROM sport_type st
		INNER JOIN trainer_sport ts ON st.id = ts.sport_id
		WHERE ts.trainer_id = $1
		ORDER BY st.name`

	rows, err := r.conn.Query(ctx, query, trainerID)
	if err != nil {
		return nil, fmt.Errorf("failed to query sport types by trainer: %w", err)
	}
	defer rows.Close()

	var sportTypes domain.SportTypeList
	for rows.Next() {
		var st domain.SportType
		if err := rows.Scan(&st.ID, &st.Name); err != nil {
			return nil, fmt.Errorf("failed to scan sport type: %w", err)
		}
		sportTypes = append(sportTypes, st)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", rows.Err())
	}

	return sportTypes, nil
}

func (r *Repository) GetSportTypesByTrainerIDTx(ctx context.Context, tx pgx.Tx, trainerID int64) (domain.SportTypeList, error) {
	query := `
		SELECT st.id, st.name
		FROM sport_type st
		INNER JOIN trainer_sport ts ON st.id = ts.sport_id
		WHERE ts.trainer_id = $1
		ORDER BY st.name`

	rows, err := tx.Query(ctx, query, trainerID)
	if err != nil {
		return nil, fmt.Errorf("failed to query sport types by trainer: %w", err)
	}
	defer rows.Close()

	var sportTypes domain.SportTypeList
	for rows.Next() {
		var st domain.SportType
		if err := rows.Scan(&st.ID, &st.Name); err != nil {
			return nil, fmt.Errorf("failed to scan sport type: %w", err)
		}
		sportTypes = append(sportTypes, st)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", rows.Err())
	}

	return sportTypes, nil
}

func (r *Repository) AddSportTypesToTrainer(ctx context.Context, tx pgx.Tx, trainerID int64, sportIDs []int64) error {
	query := `
		INSERT INTO trainer_sport (trainer_id, sport_id)
		VALUES ($1, $2)
		ON CONFLICT (trainer_id, sport_id) DO NOTHING`

	for _, sportID := range sportIDs {
		_, err := tx.Exec(ctx, query, trainerID, sportID)
		if err != nil {
			return fmt.Errorf("failed to add sport type to trainer: %w", err)
		}
	}

	return nil
}

func (r *Repository) RemoveSportTypesFromTrainer(ctx context.Context, tx pgx.Tx, trainerID int64, sportIDs []int64) error {
	query := `
		DELETE FROM trainer_sport
		WHERE trainer_id = $1 AND sport_id = $2`

	for _, sportID := range sportIDs {
		_, err := tx.Exec(ctx, query, trainerID, sportID)
		if err != nil {
			return fmt.Errorf("failed to remove sport type from trainer: %w", err)
		}
	}

	return nil
}

func (r *Repository) ReplaceTrainerSportTypes(ctx context.Context, tx pgx.Tx, trainerID int64, sportIDs []int64) error {
	// Удаляем все существующие связи
	deleteQuery := `DELETE FROM trainer_sport WHERE trainer_id = $1`
	_, err := tx.Exec(ctx, deleteQuery, trainerID)
	if err != nil {
		return fmt.Errorf("failed to delete trainer sport types: %w", err)
	}

	// Добавляем новые связи
	return r.AddSportTypesToTrainer(ctx, tx, trainerID, sportIDs)
}
