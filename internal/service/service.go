package service

import (
	"github.com/gin-gonic/gin"
	"music-library/internal/entities"
	"music-library/internal/repository"
	"music-library/internal/server/http/client"
)

type Music interface {
	StorageNewSong(ctx *gin.Context, song *entities.NewSong) (*entities.Song, error)
	GetAllSongs(ctx *gin.Context, limit, offset int, filter *entities.Filter) ([]*entities.Song, error)
	GetVerse(ctx *gin.Context, title, link string, limit, offset int) ([]string, error)
	DeleteSong(ctx *gin.Context, title, link string) error
	EditSong(ctx *gin.Context, title, link string, song *entities.Song) (*entities.Song, error)
}

type Service struct {
	Music
}

func NewService(repo *repository.Repository, httpClient *client.HTTPClient) *Service {
	return &Service{
		Music: NewMusicService(repo.MusicPostgres, httpClient),
	}
}
