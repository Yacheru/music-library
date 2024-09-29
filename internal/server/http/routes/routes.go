package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"music-library/internal/repository"
	"music-library/internal/server/http/handlers"
	"music-library/internal/service"
)

type Router struct {
	group    *gin.RouterGroup
	handler *handlers.Handler
}

func InitRoutesAndComponents(ctx context.Context, group *gin.RouterGroup, db *sqlx.DB) *Router {
	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	handler := handlers.NewHandler(serv)

	return &Router{
		group: group,
		handler: handler,
	}
}

func (r *Router) Router() {
	{
		r.handler.
	}
}
