package handlers

import (
	"miniapp/internal/handlers/handler_repository"
	"miniapp/internal/service/customer/customer_repository"
	"miniapp/internal/service/service_repository"
	"miniapp/pkg/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	logger          *logger.Logger
	customerService customer_repository.Customer
}

func NewHandler(logger *logger.Logger, service service_repository.Service_repo) handler_repository.Handler {
	return &Handler{
		logger:          logger,
		customerService: service.CustomerRepo,
	}
}

func (h *Handler) Register(r *gin.Engine) {
	r.Use(gin.Logger())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // swagger handler

	r.GET("/customer/:id", h.GetCustomerById)
	r.POST("/customer", h.CreateCustomer)
	r.PATCH("/customer/:id", h.UpdateCustomer)
}
