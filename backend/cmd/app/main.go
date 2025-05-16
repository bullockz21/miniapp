package main

import (
	"miniapp/internal/middleware"
	handlers "miniapp/internal/presentation/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.MiddlewareTest())

	r.GET("/", handlers.GetHandlerTest)
	r.POST("/", handlers.PostHandlerTest)

	r.Run(":8080")
}
