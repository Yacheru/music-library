package main

import (
	"context"
	"music-library/init/config"
	"music-library/init/logger"
	"music-library/internal/server"
	"music-library/pkg/constants"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	cfg := &config.ServerConfig
	if err := config.InitConfig(); err != nil {
		cancel()
	}

	logger.InitLogger(cfg.ApiDebug)

	serv, err := server.NewHTTPServer(ctx, cfg)
	if err != nil {
		cancel()
	}

	if serv != nil {
		serv.Run()
	}
	logger.Info("service running", constants.MainCategory)

	<-ctx.Done()

	logger.Info("service shutdown", constants.MainCategory)

	if serv != nil {
		err := serv.Shutdown(ctx)
		if err != nil {
			logger.Error(err.Error(), constants.MainCategory)
		}
	}

	logger.Info("service exited", constants.MainCategory)
}
