package handlers

import (
	"miniapp/internal/infrastructure/database/storage"
	"miniapp/pkg/cfg"
	"miniapp/pkg/logger"
	"net/http"

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
	r.GET("/hello", h.HelloHandler)
	r.GET("/users", h.GetUsersList)
	r.GET("/users/:id", h.GetUserById)
	r.POST("/users", h.CreateUser)
	r.PATCH("/users/:id/password", h.UpdateUserPasswordById)
	r.PATCH("/users/:id/email", h.UpdateUserEmailById)
	r.PATCH("/users/:id/role_id", h.UpdateUserRoleIdById)
	r.DELETE("/users/:id", h.DeleteUserById)

	r.GET("/menu", h.GetMenuList)
	r.GET("/menu/:id", h.GetMenuById)
	r.POST("/menu", h.CreateMenu)
	r.PATCH("/menu/:id", h.UpdateMenuPosition)
	r.DELETE("menu/:id", h.DeleteMenuById)
}

func (h *Handler) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, string("Hello World!"))
}
