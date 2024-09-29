package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"music-library/init/logger"
	"music-library/internal/entities"
	"music-library/internal/repository"
	"music-library/internal/server/http/client"
	"music-library/pkg/constants"
	"strings"
)

type MusicService struct {
	postgres repository.MusicPostgres
	http     *client.HTTPClient
}

func NewMusicService(postgres repository.MusicPostgres, http *client.HTTPClient) *MusicService {
	return &MusicService{postgres: postgres, http: http}
}

func (m *MusicService) EditSong(ctx *gin.Context, title, link string, song *entities.Song) (*entities.Song, error) {
	song, err := m.postgres.EditSong(ctx, title, link, song)
	if err != nil {
		logger.Error(err.Error(), constants.ServiceCategory)
		return nil, err
	}

	return song, nil
}

func (m *MusicService) DeleteSong(ctx *gin.Context, title, link string) error {
	if err := m.postgres.DeleteSong(ctx, title, link); err != nil {
		logger.Error(err.Error(), constants.ServiceCategory)

		return err
	}

	return nil
}

func (m *MusicService) GetVerse(ctx *gin.Context, title, link string, limit, offset int) ([]string, error) {
	verse, err := m.postgres.GetVerse(ctx, title, link)
	if err != nil {
		logger.Error(err.Error(), constants.ServiceCategory)

		return nil, err
	}

	verses := strings.Split(verse, "\n\n\n\n")

	if offset > len(verses) {
		return []string{}, errors.New("offset out of range")
	}

	end := offset + limit
	if end > len(verses) {
		end = len(verses)
	}

	return verses[offset:end], nil
}

func (m *MusicService) GetAllSongs(ctx *gin.Context, limit, offset int, filter *entities.Filter) ([]*entities.Song, error) {
	songs, err := m.postgres.GetAllSongs(ctx, limit, offset, filter)
	if err != nil {
		logger.Error(err.Error(), constants.ServiceCategory)

		return nil, err
	}

	return songs, nil
}

func (m *MusicService) StorageNewSong(ctx *gin.Context, song *entities.Song) error {
	metadata, err := m.http.GetSongMetadata(ctx.Request.Context(), song.Group, song.Song)
	if err != nil {
		return err
	}

	song.Lyrics = metadata.Lyrics
	song.ReleaseDate = metadata.ReleaseDate
	song.Link = metadata.Link

	if err := m.postgres.StorageNewSong(ctx, song); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return constants.SongAlreadyExistsError
		}

		logger.Error(err.Error(), constants.ServiceCategory)
		return err
	}

	return nil
}
