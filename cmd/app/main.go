package main

import (
	"context"
	"miniapp/internal/handlers"
	"miniapp/internal/infrastructure/database/db"
	"miniapp/internal/service"
	"miniapp/pkg/cfg"
	"miniapp/pkg/logger"
	"miniapp/pkg/postgresql"

	"miniapp/docs"

	"github.com/gin-gonic/gin"
)

func main() {

	docs.SwaggerInfo.Title = "Eda app API"
	docs.SwaggerInfo.Description = "Eda app API server"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:8080"
	docs.SwaggerInfo.BasePath = "/"

	logger := logger.GetLogger()
	cfg := cfg.GetConfigEnv()
	logger.Infoln("Connecting postgresql")
	psqlClient, err := postgresql.NewClient(context.Background(), 3, *cfg)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	logger.Infof("Postgresql is connected on %s:%s", cfg.Postgresql.Host, cfg.Postgresql.Port)
	err = psqlClient.Ping(context.Background())
	if err != nil {
		logger.Fatalln(err)
	}
	logger.Infoln("database pinged OK")

	storage := db.NewRepository(psqlClient, logger)

	custServ := service.NewCustomerService(storage)

	r := gin.New()

	handler := handlers.NewHandler(logger, custServ)
	handler.Register(r)
	r.Run(":8080")
}
