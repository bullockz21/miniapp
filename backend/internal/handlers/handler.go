package handlers

import (
	"context"
	"errors"
	"miniapp/internal/infrastructure/database/storage"
	"miniapp/pkg/cfg"
	"miniapp/pkg/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger  *logger.Logger
	storage storage.Storage
	cfg     cfg.Cfg
}

func NewHandler(logger *logger.Logger, storage storage.Storage) Handler {
	return Handler{
		logger:  logger,
		storage: storage,
		cfg:     *cfg.GetConfigEnv(),
	}
}

func (h *Handler) Register(r *gin.Engine) {
	// Default
	r.GET("/hello", h.HelloHandler)
	// User
	r.GET("/users", h.GetUsersList)
	r.GET("/users/:id", h.GetUserById)
	r.POST("/users", h.CreateUser)
	r.PATCH("/users/:id/password", h.UpdateUserPasswordById)
	r.PATCH("/users/:id/email", h.UpdateUserEmailById)
	r.PATCH("/users/:id/role_id", h.UpdateUserRoleIdById)
	r.DELETE("/users/:id", h.DeleteUserById)
	// Menu
	r.GET("/menu", h.GetMenuList)
	r.GET("/menu/:id", h.GetMenuById)
	r.POST("/menu", h.CreateMenu)
	r.PATCH("/menu/:id", h.UpdateMenuPosition)
	r.DELETE("menu/:id", h.DeleteMenuById)
	// Customer
	r.GET("/customer/:id", h.GetCustomerById)
	r.POST("/customer", h.CreateCustomer)
	r.PATCH("/customer/:id", h.UpdateCustomer)
}

func (h *Handler) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, string("Hello World!"))
}

var sem = make(chan struct{}, 50)

const (
	semTimeout     = 2 * time.Second
	requestTimeout = 10 * time.Second
)

func sendError(c *gin.Context, status int, msg string, err error) {
	c.JSON(status, gin.H{
		"error":   msg,
		"details": err.Error(),
	})
}

func sendSuccess(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func handleContextError(c *gin.Context, ctx context.Context) {
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		sendError(c, http.StatusGatewayTimeout, "request timeout", nil)
	} else {
		sendError(c, http.StatusRequestTimeout, "request cancelled", nil)
	}
}

func checkSemaphore(c *gin.Context, ctx context.Context) bool {
	select {
	case sem <- struct{}{}:
		return true
	case <-time.After(semTimeout):
		sendError(c, http.StatusTooManyRequests, "service busy", nil)
		return false
	case <-ctx.Done():
		handleContextError(c, ctx)
		return false
	}
}

func releaseSemaphore() {
	<-sem
}
