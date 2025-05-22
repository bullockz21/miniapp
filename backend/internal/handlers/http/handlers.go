package handlers

import (
	"miniapp/internal/infrastructure/database/storage"
	"miniapp/pkg/cfg"
	"miniapp/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TEST
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Handler struct {
	logger  *logger.Logger
	storage storage.Storage
	cfg     cfg.Cfg
}

func NewHandler(logger *logger.Logger, storage storage.Storage) Handler {
	return Handler{
		logger:  logger,
		storage: storage,
		cfg:     *cfg.GetConfig(),
	}
}

func (h *Handler) Register(r *gin.Engine) {

	r.GET("/", h.GetHandlerTest)
	r.POST("/", h.PostHandlerTest)
	// TEST END
	r.GET("/login", h.LoginHandler)
	r.GET("/hello", h.HelloHandler)
}

func (h *Handler) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, string("Hello World!"))
}

func (h *Handler) GetHandlerTest(c *gin.Context) {
	u := User{
		Name: "Susan",
		Age:  21,
	}
	c.JSON(http.StatusOK, gin.H{
		"name": u.Name,
		"age":  u.Age,
	})
}

func (h *Handler) PostHandlerTest(c *gin.Context) {
	u := &User{
		Name: "",
		Age:  0,
	}
	if err := c.ShouldBindJSON(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": u.Name,
		"age":  u.Age,
	})

}

// TEST END

func (h *Handler) LoginHandler(c *gin.Context) {

}
