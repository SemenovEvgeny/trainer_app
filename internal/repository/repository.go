package repository

import (
	"context"
	"fmt"

	"treners_app/internal/config"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	dsn  string
	conn *pgx.Conn
}

func NewRepository(ctx context.Context) (repo *Repository, err error) {
	cfg := config.GetConfig()

	repo = &Repository{
		dsn: cfg.DSN,
	}

	var conn *pgx.ConnConfig

	conn, err = pgx.ParseConfig(cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PostgreSQL config: %w", err)
	}

	repo.conn, err = pgx.ConnectConfig(ctx, conn)
	if err != nil {
		return nil, fmt.Errorf("failed to create PostgreSQL Connection pool: %w", err)
	}

	return repo, nil
}
