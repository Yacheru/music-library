package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"music-library/init/config"
	"music-library/internal/repository"
	"music-library/internal/server/http/client"
	"music-library/internal/server/http/handlers"
	"music-library/internal/server/http/middleware"
	"music-library/internal/service"
)

type Router struct {
	group   *gin.RouterGroup
	handler *handlers.Handler
}

func InitRoutesAndComponents(ctx context.Context, cfg *config.Config, group *gin.RouterGroup, db *sqlx.DB) *Router {
	repo := repository.NewRepository(db)
	httpClient := client.NewHTTPClient(cfg)
	serv := service.NewService(repo, httpClient)
	handler := handlers.NewHandler(serv)

	return &Router{
		group:   group,
		handler: handler,
	}
}

func (r *Router) Router() {
	{
		r.group.GET("/all", middleware.ValidateQuery(), middleware.ValidateFilter(), r.handler.GetAllSongs) // Получение данных библиотеки с фильтрацией по всем полям и пагинацией
		r.group.GET("/verse", middleware.ValidateQuery(), middleware.ValidateSong(), r.handler.GetVerse)    // Получение текста песни с пагинацией по куплетам
		r.group.DELETE("/", middleware.ValidateSong(), r.handler.DeleteSong)                                // Удаление песни
		r.group.PATCH("/", middleware.ValidateSong(), r.handler.EditSong)                                   // Изменение данных песни
		r.group.POST("/new", r.handler.StorageNewSong)                                                      // Добавление новой песни
	}
}
