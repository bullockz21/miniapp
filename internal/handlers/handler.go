package handlers

import (
	"miniapp/internal/handlers/handler_repository"
	"miniapp/internal/service/customer/customer_repository"
	"miniapp/internal/service/service_repository"
	"miniapp/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	logger   *logger.Logger
	customer customer_repository.Customer
	sem      chan struct{}
}

func NewHandler(logger *logger.Logger, service service_repository.Service_repo) handler_repository.Handler {
	return &Handler{
		logger:   logger,
		customer: service.CustomerRepo,
	}
}

func (h *Handler) Register(r *gin.Engine) {
	r.Use(gin.Logger())
	main := r.Group("/")
	{
		main.GET("/customer/:id", h.GetCustomerById)
		main.GET("/customer", h.GetCustomerList)
		main.POST("/customer", h.CreateCustomer)
		main.PATCH("/customer", h.UpdateCustomer)
	}
	main.Use(h.TimeoutAndSemoporeMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // swagger handler

	r.GET("/hello", h.Hello)

}

func (h *Handler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
