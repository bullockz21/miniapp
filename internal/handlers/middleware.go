package handlers

import (
	"context"
	"miniapp/internal/service"
	"miniapp/pkg/cfg"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	semTimeout     = 2 * time.Second
	requestTimeout = 10 * time.Second
)

func (h *Handler) TimeoutAndSemoporeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), requestTimeout)
		defer cancel()
		if !h.checkSemaphore(c, ctx) {
			c.Abort()
			return
		}
		defer h.releaseSemaphore()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func (h *Handler) checkSemaphore(c *gin.Context, ctx context.Context) bool {
	select {
	case h.sem <- struct{}{}:
		return true
	case <-time.After(semTimeout):
		SendError(c, http.StatusTooManyRequests, Result{data: "service busy", err: nil})
		return false
	case <-ctx.Done():
		handleContextError(c, ctx)
		return false
	}
}

func (h *Handler) releaseSemaphore() {
	<-h.sem
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": "false",
				"status":  "Unauthorized access",
				"msg":     "empty or incorrect token",
			})
			return
		}

		hArr := strings.Split(h, "Bearer ")

		if hArr[1] == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": "false",
				"status":  "Unauthorized access",
				"msg":     "empty or incorrect token",
			})
			return
		}
		_, err := service.ParseToken(hArr[1], cfg.GetConfigEnv().JwtSecretKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": "false",
				"status":  "Unauthorized access",
				"msg":     "empty or incorrect token",
			})
			return
		}
		c.Next()
	}
}
