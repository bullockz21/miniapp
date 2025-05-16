package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func MiddlewareTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Hello from middleware")
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
