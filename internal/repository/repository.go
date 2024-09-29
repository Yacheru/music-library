package repository

import (
	"github.com/jmoiron/sqlx"
	"music-library/internal/repository/postgres"
)

type MusicPostgres interface {
}

type Repository struct {
	MusicPostgres
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		MusicPostgres: postgres.NewMusicPostgres(db),
	}
}
