package main

import (
	handlers "miniapp/internal/handlers/http"
	"miniapp/internal/handlers/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.AuthMiddleware())
	// TEST
	r.Use(middleware.MiddlewareTest())
	r.GET("/", handlers.GetHandlerTest)
	r.POST("/", handlers.PostHandlerTest)
	// TEST END
	r.GET("/login", handlers.LoginHandler)

	r.Run(":8080")
}
