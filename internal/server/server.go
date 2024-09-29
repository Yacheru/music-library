package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"music-library/init/config"
	"music-library/internal/repository/postgres"
	"music-library/internal/server/http/routes"
	"net/http"
	"time"

	"music-library/init/logger"
	"music-library/pkg/constants"
)

type Server struct {
	http *http.Server
}

func NewHTTPServer(ctx context.Context, cfg *config.Config) (*Server, error) {
	db, err := postgres.NewConnection(ctx, cfg)
	if err != nil {
		return nil, err
	}

	engine := setupGin(cfg)
	group := engine.Group(cfg.ApiEntry)
	routes.InitRoutesAndComponents(cfg, group, db).Router()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.ApiPort),
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{
		http: server,
	}, nil
}

func (s *Server) Run() {
	go func() {
		if err := s.http.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(err.Error(), constants.ServerCategory)
		}
	}()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}

func setupGin(cfg *config.Config) *gin.Engine {
	var mode = gin.ReleaseMode
	if cfg.ApiDebug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	engine := gin.Default()

	engine.Use(gin.Recovery())
	engine.Use(gin.LoggerWithFormatter(logger.HTTPLogger))

	return engine
}
