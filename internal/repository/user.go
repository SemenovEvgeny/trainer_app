package repository

import (
	"context"
	"fmt"

	"treners_app/internal/domain"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetRoleByValue(ctx context.Context, roleValue string) (*domain.Role, error) {
	query := `SELECT id, value FROM role WHERE value = $1`

	var role domain.Role
	err := r.conn.QueryRow(ctx, query, roleValue).Scan(&role.ID, &role.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to get role by value: %w", err)
	}

	return &role, nil
}

func (r *Repository) CreateUser(ctx context.Context, tx pgx.Tx, user *domain.User) error {
	query := `
		INSERT INTO users (email, password_hash, role_id)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at`

	err := tx.QueryRow(ctx, query,
		user.Email,
		user.PasswordHash,
		user.RoleID,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT id, email, password_hash, role_id, created_at, updated_at
		FROM users
		WHERE email = $1`

	var user domain.User
	err := r.conn.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return &user, nil
}
