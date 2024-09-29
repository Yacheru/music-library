package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"music-library/init/config"
	"music-library/pkg/spotify"
	"net/http"
	"sync"
	"time"

	"music-library/init/logger"
	"music-library/internal/entities"
	"music-library/pkg/constants"
)

type HTTPClient struct {
	client *http.Client
	wg     *sync.WaitGroup

	spotifyId     string
	spotifySecret string
}

func NewHTTPClient(cfg *config.Config) *HTTPClient {
	return &HTTPClient{
		client: new(http.Client),
		wg:     new(sync.WaitGroup),

		spotifyId:     cfg.SpotifyId,
		spotifySecret: cfg.SpotifySecret,
	}
}

func (c *HTTPClient) GetSongMetadata(ctx context.Context, artist, title string) (*entities.Metadata, error) {
	lyrics, err := c.getSongLyrics(ctx, artist, title)
	if err != nil {
		logger.Error(err.Error(), constants.ClientCategory)

		return nil, err
	}

	date, link, err := c.getSongLinkAndReleaseDate(ctx, artist, title)
	if err != nil {
		logger.Error(err.Error(), constants.ClientCategory)

		return nil, err
	}

	return &entities.Metadata{
		ReleaseDate: date,
		Lyrics:      lyrics,
		Link:        link,
	}, nil

}

func (c *HTTPClient) getSongLyrics(ctx context.Context, artist, title string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	var lyrics = new(entities.Lyrics)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://api.lyrics.ovh/v1/%s/%s", artist, title), nil)
	if err != nil {
		return "", err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return "", constants.TimeOutError
		}

		return "", err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(b, lyrics)
	if err != nil {
		return "", err
	}

	return lyrics.Lyrics, nil
}

func (c *HTTPClient) getSongLinkAndReleaseDate(ctx context.Context, artist, title string) (*time.Time, string, error) {
	s := spotify.NewClient(c.spotifyId, c.spotifySecret)

	res, err := s.SearchForTrack(ctx, artist, title)
	if err != nil {
		logger.Error(err.Error(), constants.ClientCategory)

		return nil, "", err
	}

	if res.Tracks == nil {
		return nil, "", constants.SongNotFoundError
	}

	date := s.ReleaseDateTime(res.Tracks.Tracks[0].Album.ReleaseDatePrecision, res.Tracks.Tracks[0].Album.ReleaseDate)

	return date, res.Tracks.Tracks[0].ExternalURLs["spotify"], nil
}
