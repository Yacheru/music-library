package entities

import "time"

type NewSong struct {
	Group string `json:"group" db:"group"`
	Song  string `json:"song" db:"title"`
}

type Song struct {
	Group    string `json:"group" db:"group" swaggerignore:"true"`
	Song     string `json:"song" db:"title" swaggerignore:"true"`
	Metadata `json:"metadata,omitempty" swaggerignore:"true"`
}

type Metadata struct {
	ReleaseDate *time.Time `json:"release_date" db:"release_date"` // При изменении: ISO 8601 или RFC 3339
	Lyrics      string     `json:"lyrics" db:"text"`
	Link        string     `json:"link" db:"link"`
}

type Lyrics struct {
	Lyrics string `json:"lyrics"`
}
