package postgres

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/golang-migrate/migrate/v4"
	"music-library/init/config"
	"music-library/init/logger"
	"music-library/pkg/constants"
)

func NewConnection(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", cfg.PostgresDSN)
	if err != nil {
		logger.Error(err.Error(), constants.PostgresCategory)

		return nil, err
	}

	if err := db.Ping(); err != nil {
		logger.Error(err.Error(), constants.PostgresCategory)
		return nil, err
	}

	if err := migrate.New()
}
