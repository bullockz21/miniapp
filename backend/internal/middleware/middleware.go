package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiddlewareTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Hello from middleware")
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" || !VerifyToken(token) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			ctx.Redirect(http.StatusMovedPermanently, "/login")
		}
		ctx.Next()
	}
}

func VerifyToken(token string) bool {

	return false
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
