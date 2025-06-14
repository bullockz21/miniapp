package handlers

import (
	"miniapp/internal/domain/user"
	"miniapp/internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsersList(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := h.storage.LoadAllUsers(ctx)

	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "FAIL: error get users list", err: err})
		return
	}
	sendSuccess(c, http.StatusOK, Result{data: users, err: nil})

}

func (h *Handler) GetUserById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
		return
	}

	user, err := h.storage.LoadUser(ctx, int(id))

	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "FAIL: error get user by id", err: err})
		return
	}
	sendSuccess(c, http.StatusOK, Result{data: user, err: nil})
}

func (h *Handler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	newUser := dto.UserDTO{}
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	newUser.PasswordHash, err = user.CreatePasswordHash(newUser.Password)
	newUser.Password = ""
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid hash", err: err})
		return
	}

	id, err := h.storage.SaveNewUser(ctx, newUser)

	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "FAIL: error create user", err: err})
		return
	}
	sendSuccess(c, http.StatusOK, Result{data: id, err: nil})
}

func (h *Handler) UpdateUserPasswordById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
		return
	}
	updateUser := dto.UserDTO{}
	err = c.ShouldBindJSON(&updateUser)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	updateUser.Id = int(id)
	updateUser.PasswordHash, err = user.CreatePasswordHash(updateUser.Password)
	updateUser.Password = ""
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid hash", err: err})
		return
	}

	outId, err := h.storage.UpdateUserPassword(ctx, updateUser)

	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "FAIL: error update user password", err: err})
		return
	}
	sendSuccess(c, http.StatusOK, Result{data: outId, err: nil})
}

func (h *Handler) UpdateUserEmailById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
		return
	}
	updateUser := dto.UserDTO{}
	err = c.ShouldBindJSON((&updateUser))
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	updateUser.Id = int(id)

	outId, err := h.storage.UpdateUserEmail(ctx, updateUser)

	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "FAIL: error update user email", err: err})
		return
	}
	sendSuccess(c, http.StatusOK, Result{data: outId, err: nil})
}

func (h *Handler) UpdateUserRoleIdById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
		return
	}
	updateUser := dto.UserDTO{}
	err = c.ShouldBindJSON((&updateUser))
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	updateUser.Id = int(id)

	outId, err := h.storage.UpdateUserRoleId(ctx, updateUser)

	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "FAIL: error update users role", err: err})
		return
	}
	sendSuccess(c, http.StatusOK, Result{data: outId, err: nil})

}

func (h *Handler) DeleteUserById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
		return
	}
	deleteUser := dto.UserDTO{}
	err = c.ShouldBindJSON((&deleteUser))
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	deleteUser.Id = int(id)

	err = h.storage.DeleteUser(ctx, deleteUser.Id)

	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "FAIL: error delete users", err: err})
		return
	}
	sendSuccess(c, http.StatusOK, Result{data: "User by DELETE", err: nil})
}
