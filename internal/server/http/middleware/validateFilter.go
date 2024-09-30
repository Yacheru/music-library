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

		f := ctx.Query("filter")
		if f == "" {
			ctx.Next()
			return
		}

		parts := strings.SplitN(f, ":", 2)
		if len(parts) != 2 {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "Invalid filter format. Example: &filter=group: SomeGroup")
			return
		}

		filter, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		if value == "" {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "Filter value not provided")
			return
		}

		var entityFilter = new(entities.Filter)
		switch filter {
		case "group":
			entityFilter.Group = &value
		case "song":
			entityFilter.Song = &value
		case "lyrics":
			entityFilter.Lyrics = &value
		case "link":
			entityFilter.Link = &value
		case "release_date":
			if value != "ASC" && value != "DESC" {
				handlers.NewErrorResponse(ctx, http.StatusBadRequest, "Release_date filter supports only ASC or DESC")
				return
			}

			entityFilter.ReleaseDate = &value
		default:
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "Unsupported filter type")
			return
		}

		ctx.Set("filter", entityFilter)

		ctx.Next()
	}
}
