package postgres

import "github.com/jmoiron/sqlx"

type MusicPostgres struct {
	db *sqlx.DB
}

func NewMusicPostgres(db *sqlx.DB) *MusicPostgres {
	return &MusicPostgres{db: db}
}
