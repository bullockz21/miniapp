package main

import (
	handlers "miniapp/internal/handlers/http"
	"miniapp/internal/handlers/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.MiddlewareTest())

	r.GET("/", handlers.GetHandlerTest)
	r.POST("/", handlers.PostHandlerTest)

	r.Run(":8080")
}
