package postgres

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"music-library/internal/entities"
	"music-library/pkg/constants"
)

type MusicPostgres struct {
	db *sqlx.DB
}

func NewMusicPostgres(db *sqlx.DB) *MusicPostgres {
	return &MusicPostgres{db: db}
}

func (m *MusicPostgres) EditSong(ctx *gin.Context, title, link string, song *entities.Song) (*entities.Song, error) {
	var entitySong = new(entities.Song)

	query := `
		UPDATE songs
		SET 
			"group" = CASE WHEN $1 != '' THEN $1 ELSE "group" END,
			title = CASE WHEN $2 != '' THEN $2 ELSE title END,
			text = CASE WHEN $3 != '' THEN $3 ELSE text END,
			link = CASE WHEN $4 != '' THEN $4 ELSE link END,
			release_date = CASE WHEN $5::TIMESTAMP IS NOT NULL THEN $5 ELSE release_date END
		WHERE title = $6 OR link = $7
		RETURNING "group", title, text, link, release_date
	`
	err := m.db.GetContext(ctx, entitySong, query, song.Group, song.Song, song.Lyrics, song.Link, song.ReleaseDate, title, link)
	if err != nil {
		return nil, err
	}

	return entitySong, nil
}

func (m *MusicPostgres) DeleteSong(ctx *gin.Context, title, link string) error {
	query := `
		DELETE FROM songs WHERE title=$1 OR link=$2;
	`
	rows, err := m.db.ExecContext(ctx, query, title, link)
	if err != nil {
		return err
	}
	count, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return constants.SongNotFoundError
	}

	return nil
}

func (m *MusicPostgres) GetVerse(ctx *gin.Context, title, link string) (string, error) {
	var verses string

	query := `
		SELECT text FROM songs
		WHERE title = $1 OR link = $2
	`
	err := m.db.GetContext(ctx.Request.Context(), &verses, query, title, link)
	if err != nil {
		return "", err
	}

	return verses, nil
}

func (m *MusicPostgres) GetAllSongs(ctx *gin.Context, limit, offset int, filter *entities.Filter) ([]*entities.Song, error) {
	var songs []*entities.Song
	var query string
	var args []interface{}

	query = `SELECT "group", title, text, link, release_date FROM songs `

	if filter != nil {
		if filter.Song != nil {
			query += `WHERE title ILIKE '%' || $1 || '%' LIMIT $2 OFFSET $3`
			args = append(args, *filter.Song, limit, offset)
		} else if filter.Link != nil {
			query += `WHERE link ILIKE '%' || $1 || '%' LIMIT $2 OFFSET $3`
			args = append(args, *filter.Link, limit, offset)
		} else if filter.Group != nil {
			query += `WHERE "group" ILIKE '%' || $1 || '%' LIMIT $2 OFFSET $3`
			args = append(args, *filter.Group, limit, offset)
		} else if filter.Lyrics != nil {
			query += `WHERE text ILIKE '%' || $1 || '%' LIMIT $2 OFFSET $3`
			args = append(args, *filter.Lyrics, limit, offset)
		} else if filter.ReleaseDate != nil {
			query += fmt.Sprintf(`ORDER BY release_date %s LIMIT $1 OFFSET $2`, *filter.ReleaseDate)
			args = append(args, limit, offset)
		}
	} else {
		query += `LIMIT $1 OFFSET $2`
		args = append(args, limit, offset)
	}

	fmt.Println(query)

	err := m.db.SelectContext(ctx.Request.Context(), &songs, query, args...)
	if err != nil {
		return nil, err
	}

	return songs, nil
}

func (m *MusicPostgres) StorageNewSong(ctx *gin.Context, song *entities.Song) error {
	query := `
		INSERT INTO songs ("group", title, text, link, release_date) 
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := m.db.ExecContext(ctx.Request.Context(), query, song.Group, song.Song, song.Lyrics, song.Link, song.ReleaseDate)
	if err != nil {
		return err
	}

	return nil
}
