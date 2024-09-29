package middleware

import (
	"github.com/gin-gonic/gin"
	"music-library/internal/server/http/handlers"
	"net/http"
)

func ValidateSong() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		title := ctx.Query("title")
		link := ctx.Query("link")

		if title == "" && link == "" {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "title or link required")
			return
		}

		ctx.Next()
	}
}
