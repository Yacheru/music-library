package postgres

import (
	"context"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"music-library/init/config"
	"music-library/init/logger"
	"music-library/pkg/constants"
)

func NewConnection(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "pgx", cfg.PostgresDSN)
	if err != nil {
		logger.Error(err.Error(), constants.PostgresCategory)
		return nil, err
	}

	if err := goose.Up(db.DB, "./migrations"); err != nil {
		logger.Error(err.Error(), constants.PostgresCategory)
		return nil, err
	}

	return db, nil
}
