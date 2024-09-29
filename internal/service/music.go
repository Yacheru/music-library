package service

import "music-library/internal/repository"

type Music struct {
	postgres repository.MusicPostgres
}

func NewMusicService(postgres repository.MusicPostgres) *Music {
	return &Music{postgres: postgres}
}
