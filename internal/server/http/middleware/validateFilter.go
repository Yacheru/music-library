package middleware

import (
	"github.com/gin-gonic/gin"
	"music-library/internal/entities"
	"music-library/internal/server/http/handlers"
	"net/http"
	"strings"
)

func ValidateFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// &filter=group: SomeGroup
		// &filter=song: SomeSong
		// &filter=lyrics: SomeLyrics
		// &filter=link: SomeLink
		// &filter=release_date: ASC || DESC

		filter := ctx.Query("filter")

		if filter == "" {
			ctx.Next()
			return
		}

		var parts []string

		if strings.Count(filter, ": ") == 1 {
			parts = strings.Split(filter, ": ")
		} else {
			parts = strings.Split(filter, ":")
		}

		if len(parts) != 2 {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "Invalid filter. Example: &filter=group: SomeGroup")
			return
		}

		if parts[1] == "" {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "Filter value not provided")
			return
		}

		var entityFilter = new(entities.Filter)
		switch parts[0] {
		case "group":
			entityFilter.Group = &parts[1]
		case "song":
			entityFilter.Song = &parts[1]
		case "lyrics":
			entityFilter.Lyrics = &parts[1]
		case "link":
			entityFilter.Link = &parts[1]
		case "release_date":
			entityFilter.ReleaseDate = &parts[1]

			if parts[1] != "ASC" && parts[1] != "DESC" {
				handlers.NewErrorResponse(ctx, http.StatusBadRequest, "Release_date filter support ASC or DESC")
				return
			}
		default:
			ctx.Next()
			return
		}

		ctx.Set("filter", entityFilter)

		ctx.Next()
	}
}
