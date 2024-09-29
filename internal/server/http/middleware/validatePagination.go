package middleware

import (
	"github.com/gin-gonic/gin"
	"music-library/internal/server/http/handlers"
	"net/http"
	"strconv"
)

func ValidateQuery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		offset, err := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
		if err != nil || offset < 0 {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "offset must be a positive integer")
			return
		}

		limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
		if err != nil || limit < 0 {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "limit must be a positive integer")
			return
		}

		ctx.Next()
	}
}
