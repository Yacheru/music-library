package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"music-library/internal/entities"
	"music-library/internal/repository/postgres"
)

type MusicPostgres interface {
	StorageNewSong(ctx *gin.Context, song *entities.Song) (*entities.Song, error)
	GetAllSongs(ctx *gin.Context, limit, offset int, filter *entities.Filter) ([]*entities.Song, error)
	GetVerse(ctx *gin.Context, title, link string) (string, error)
	DeleteSong(ctx *gin.Context, title, link string) error
	EditSong(ctx *gin.Context, title, link string, song *entities.Song) (*entities.Song, error)
}

type Repository struct {
	MusicPostgres
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		MusicPostgres: postgres.NewMusicPostgres(db),
	}
}
