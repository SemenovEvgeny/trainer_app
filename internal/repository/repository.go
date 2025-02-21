package repository

import (
	"context"
	"fmt"
	"treners_app/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repository struct {
	dsn  string
	Conn *pgxpool.Pool
	log  *zap.Logger
}

func NewRepository(ctx context.Context, cfg config.Storage, log *zap.Logger) (repo *Repository, err error) {
	repo = &Repository{
		dsn: cfg.DSN,
	}

	var pool *pgxpool.Config

	pool, err = pgxpool.ParseConfig(cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PostgreSQL config: %w", err)
	}

	pool.MaxConns = int32(cfg.DBMaxConnection)

	repo.Conn, err = pgxpool.NewWithConfig(ctx, pool)
	if err != nil {
		return nil, fmt.Errorf("failed to create PostgreSQL Connection pool: %w", err)
	}

	repo.log = log

	return repo, nil
}
