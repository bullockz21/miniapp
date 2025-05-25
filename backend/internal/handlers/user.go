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

func (h *Handler) UpdateUserPasswordById(c *gin.Context) {
	updateUser := dto.UserDTO{}

	idStr := c.Params.ByName("id")
	id := 0
	fmt.Sscanf(idStr, "%d", &id)

	err := c.BindJSON((&updateUser))

	if err != nil {
		c.JSON(http.StatusBadRequest, string(fmt.Sprintf("%v", err)))
		return
	}

	updateUser.Id = id

	updateUser.PasswordHash, err = user.CreatePasswordHash(updateUser.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, string(fmt.Sprintf("%v", err)))
		return
	}

	id, err = h.storage.UpdateUserPassword(context.Background(), updateUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, string(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) UpdateUserEmailById(c *gin.Context) {
	updateUser := dto.UserDTO{}

	idStr := c.Params.ByName("id")
	id := 0
	fmt.Sscanf(idStr, "%d", &id)

	err := c.BindJSON((&updateUser))

	fmt.Println(updateUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, string(fmt.Sprintf("%v", err)))
		return
	}

	updateUser.Id = id

	id, err = h.storage.UpdateUserEmail(context.Background(), updateUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, string(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) UpdateUserRoleIdById(c *gin.Context) {
	updateUser := dto.UserDTO{}

	idStr := c.Params.ByName("id")
	id := 0
	fmt.Sscanf(idStr, "%d", &id)

	err := c.BindJSON((&updateUser))

	if err != nil {
		c.JSON(http.StatusBadRequest, string(fmt.Sprintf("%v", err)))
		return
	}

	updateUser.Id = id

	id, err = h.storage.UpdateUserRoleId(context.Background(), updateUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, string(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) DeleteUserById(c *gin.Context) {
	deleteUser := dto.UserDTO{}

	idStr := c.Params.ByName("id")
	id := 0
	fmt.Sscanf(idStr, "%d", &id)

	err := c.BindJSON((&deleteUser))

	deleteUser.Id = id

	if err != nil {
		c.JSON(http.StatusBadRequest, string(fmt.Sprintf("%v", err)))
		return
	}

	err = h.storage.DeleteUser(context.Background(), deleteUser.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, string(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, id)
}
