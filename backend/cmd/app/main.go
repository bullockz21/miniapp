package main

import (
	"context"
	"miniapp/internal/handlers"
	"miniapp/internal/infrastructure/database/db"
	"miniapp/pkg/cfg"
	"miniapp/pkg/logger"
	"miniapp/pkg/postgresql"

	"github.com/gin-gonic/gin"
)

func main() {

	logger := logger.GetLogger()

	cfg := cfg.GetConfig()

	logger.Infoln("Connecting postgresql")

	psqlClient, err := postgresql.NewClient(context.Background(), 3, *cfg)
	if err != nil {
		logger.Fatalf("%v", err)
	}

	logger.Infof("Postgresql is connected on %s:%s", cfg.Postgresql.Host, cfg.Postgresql.Port)

	repository := db.NewRepository(psqlClient, logger)

	r := gin.Default()

	// r.Use(middleware.AuthMiddleware())
	// r.Use(middleware.MiddlewareTest())

	handler := handlers.NewHandler(logger, repository)
	handler.Register(r)

	r.Run(":8181")
}
