package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"music-library/pkg/spotify/entities"
	"net/http"
	"net/url"
)

type Client struct {
	http *http.Client

	id     string
	secret string

	accessToken string
}

func NewClient(id string, secret string) *Client {
	return &Client{
		http:        new(http.Client),
		accessToken: "",
		id:          id,
		secret:      secret,
	}
}

func (c *Client) SearchForTrack(ctx context.Context, artist, track string) (*entities.SearchResult, error) {
	v := url.Values{}
	v.Set("q", fmt.Sprintf("track:%s artist:%s", track, artist))
	v.Set("type", "track")
	v.Set("limit", "1")

	spotifyURL := "https://api.spotify.com/v1/" + "search?" + v.Encode()

	var searchResult = new(entities.SearchResult)

	req, err := http.NewRequest(http.MethodGet, spotifyURL, nil)
	if err != nil {
		return nil, err
	}

	if c.accessToken == "" {
		err := c.getAccessToken(ctx)
		if err != nil {
			return nil, err
		}
	}

	req.Header.Add("Authorization", "Bearer "+c.accessToken)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		err := c.getAccessToken(ctx)
		if err != nil {
			return nil, err
		}
		return c.SearchForTrack(ctx, artist, track)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, searchResult)
	if err != nil {
		return nil, err
	}

	return searchResult, nil
}

func (c *Client) getAccessToken(ctx context.Context) error {
	var token = new(entities.AccessTokenResponse)
	uri := fmt.Sprintf("https://accounts.spotify.com/api/token?grant_type=client_credentials&client_id=%s&client_secret=%s", c.id, c.secret)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, token)
	if err != nil {
		return err
	}

	c.accessToken = token.AccessToken

	return nil
}
