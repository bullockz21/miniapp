package main

import (
	"context"
	"log"
	"miniapp/internal/handlers"
	"miniapp/internal/infrastructure/database/db"
	"miniapp/internal/service/service_repository"
	"miniapp/pkg/cfg"
	"miniapp/pkg/logger"
	"miniapp/pkg/postgresql"
	"os"
	"os/signal"
	"syscall"

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

	r := gin.New()

	service := service_repository.NewServiceRepository(storage)

	handler := handlers.NewHandler(logger, service)
	handler.Register(r)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("\nInterrupt signal received. Exiting...")
		psqlClient.Close()
		os.Exit(0)
	}()

	r.Run(":8080")
}
