package service

import "music-library/internal/repository"

type MusicService interface {
}

type Service struct {
	MusicService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		MusicService: NewMusicService(repo.MusicPostgres),
	}
}
