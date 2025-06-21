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
}

func NewHandler(logger *logger.Logger, service service_repository.Service_repo) handler_repository.Handler {
	return &Handler{
		logger:   logger,
		customer: service.CustomerRepo,
	}
}

func (h *Handler) Register(r *gin.Engine) {
	r.Use(gin.Logger())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // swagger handler

	r.GET("/hello", h.Hello)

	r.GET("/customer/:id", h.GetCustomerById)
	r.GET("/customer", h.GetCustomerList)
	r.POST("/customer", h.CreateCustomer)
	r.PATCH("/customer", h.UpdateCustomer)
}

func (h *Handler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
