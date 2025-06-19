package handlers

import (
	"miniapp/internal/infrastructure/database/storage"
	"miniapp/pkg/cfg"
	"miniapp/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	logger  *logger.Logger
	storage storage.Storage
	cfg     cfg.Cfg
	sem     chan struct{}
}

func NewHandler(logger *logger.Logger, storage storage.Storage) Handler {
	return Handler{
		logger:  logger,
		storage: storage,
		cfg:     *cfg.GetConfigEnv(),
		sem:     make(chan struct{}, 50),
	}
}

func (h *Handler) Register(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Default
	r.GET("/hello", h.TimeoutAndSemoporeMiddleware(), h.HelloHandler)
	// User
	r.GET("/users", h.TimeoutAndSemoporeMiddleware(), h.GetUsersList)
	r.GET("/users/:id", h.TimeoutAndSemoporeMiddleware(), h.GetUserById)
	r.POST("/users", h.TimeoutAndSemoporeMiddleware(), h.CreateUser)
	r.PATCH("/users/:id/password", h.TimeoutAndSemoporeMiddleware(), h.UpdateUserPasswordById)
	r.PATCH("/users/:id/email", h.TimeoutAndSemoporeMiddleware(), h.UpdateUserEmailById)
	r.PATCH("/users/:id/role_id", h.TimeoutAndSemoporeMiddleware(), h.UpdateUserRoleIdById)
	r.DELETE("/users/:id", h.TimeoutAndSemoporeMiddleware(), h.DeleteUserById)
	// Menu
	r.GET("/menu", h.TimeoutAndSemoporeMiddleware(), h.GetMenuList)
	r.GET("/menu/:id", h.TimeoutAndSemoporeMiddleware(), h.GetMenuById)
	r.POST("/menu", h.TimeoutAndSemoporeMiddleware(), h.CreateMenu)
	r.PATCH("/menu/:id", h.TimeoutAndSemoporeMiddleware(), h.UpdateMenuPosition)
	r.DELETE("menu/:id", h.TimeoutAndSemoporeMiddleware(), h.DeleteMenuById)
	// Customer
	r.GET("/customer/:id", h.TimeoutAndSemoporeMiddleware(), h.GetCustomerById)
	r.POST("/customer", h.TimeoutAndSemoporeMiddleware(), h.CreateCustomer)
	r.PATCH("/customer/:id", h.TimeoutAndSemoporeMiddleware(), h.UpdateCustomer)
	// Orders
	r.GET("/order/:id", h.TimeoutAndSemoporeMiddleware(), h.GetOrdersByOrderId)
	r.GET("/order", h.TimeoutAndSemoporeMiddleware(), h.GetAllOrders)
	r.POST("/order", h.TimeoutAndSemoporeMiddleware(), h.CreateOrder)
	r.PATCH("/order/:id", h.TimeoutAndSemoporeMiddleware(), h.UpdateOrder)
}

func (h *Handler) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, string("Hello World!"))
}
