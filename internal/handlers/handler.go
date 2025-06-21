package handlers

import (
	"fmt"
	"miniapp/internal/dto"
	"miniapp/internal/handlers/handler_repository"
	"miniapp/internal/service"
	"miniapp/internal/service/customer_repository"
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
	user     service.User
	sem      chan struct{}
}

func NewHandler(logger *logger.Logger, service service_repository.Service_repo) handler_repository.Handler {
	return &Handler{
		logger:   logger,
		customer: service.CustomerRepo,
	}
}

func (h *Handler) Register(r *gin.Engine) {

	r.POST("/auth", h.Authentification)

	main := r.Group("/")
	{
		main.GET("/customer/:id", h.GetCustomerById)
		main.GET("/customer", h.GetCustomerList)
		main.POST("/customer", h.CreateCustomer)
		main.PATCH("/customer", h.UpdateCustomer)
	}
	r.Use(gin.Logger())
	main.Use(h.TimeoutAndSemoporeMiddleware())
	main.Use(h.Authentification)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // swagger handler

	r.GET("/hello", h.Hello)

}

func (h *Handler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

// Authentification godoc
//
//	@Summary		Authentification
//	@Description	Authentification by login and password
//	@Tags			Authentification
//	@Accept			json
//	@Produce		json
//	@Param			account	body		dto.AuthDTO	true	"Auth data"
//	@Success		200		{object}	handlers.Success
//	@Failure		400		{object}	handlers.Error
//	@Failure		401		{object}	handlers.Error
//	@Router			/auth [post]
func (h *Handler) Authentification(c *gin.Context) {
	ctx := c.Request.Context()
	userData := dto.AuthDTO{}

	err := c.BindJSON(&userData)
	if err != nil {
		SendError(c, http.StatusBadRequest, Result{
			data: "invalid request body",
			err:  err,
		})
		return
	}

	token, err := h.user.AuthUser(ctx, userData)

	if err != nil {
		SendError(c, http.StatusUnauthorized, Result{
			data: "access denied, auth error",
			err:  err,
		})
		return
	}

	SendSuccess(c, http.StatusOK, Result{
		data: fmt.Sprintf("Bearer %s", token),
		err:  err,
	})
}
