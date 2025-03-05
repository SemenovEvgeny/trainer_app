package repository

import (
	"context"
	"fmt"
	"log/slog"

	"treners_app/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	dsn  string
	conn *pgxpool.Pool
	log  *slog.Logger
}

func NewRepository(ctx context.Context, cfg config.Storage, log *slog.Logger) (repo *Repository, err error) {
	repo = &Repository{
		dsn: cfg.DSN,
	}

	var pool *pgxpool.Config

	pool, err = pgxpool.ParseConfig(cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PostgreSQL config: %w", err)
	}

	pool.MaxConns = cfg.MaxConnection

	repo.conn, err = pgxpool.NewWithConfig(ctx, pool)
	if err != nil {
		return nil, fmt.Errorf("failed to create PostgreSQL Connection pool: %w", err)
	}

	repo.log = log

	return repo, nil
}
