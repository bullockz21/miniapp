package handlers

import (
	"context"
	"fmt"
	"miniapp/internal/domain/user"
	"miniapp/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsersList(c *gin.Context) {
	users, err := h.storage.LoadAllUsers(context.Background())

	if err != nil {
		c.JSON(http.StatusBadRequest, string(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUserById(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id := 0
	fmt.Sscanf(idStr, "%d", &id)
	user, err := h.storage.LoadUser(context.Background(), id)

	if err != nil {
		c.JSON(http.StatusBadRequest, string(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) CreateUser(c *gin.Context) {
	newUser := dto.UserDTO{}

	err := c.BindJSON((&newUser))

	if err != nil {
		c.JSON(http.StatusBadRequest, string(fmt.Sprintf("%v", err)))
		return
	}

	newUser.PasswordHash, err = user.CreatePasswordHash(newUser.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, string(fmt.Sprintf("%v", err)))
		return
	}

	newUser.Password = ""

	id, err := h.storage.SaveNewUser(context.Background(), newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, string(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, id)
}
