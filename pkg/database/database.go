package database

import (
	"context"

	"github.com/LigeronAhill/go-fiber/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDBPool(cfg *config.DatabaseConfig) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), cfg.URL)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
