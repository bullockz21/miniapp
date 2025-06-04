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
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		users, err := h.storage.LoadAllUsers(ctx)
		select {
		case resultChan <- Result{data: users, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusBadRequest, Result{data: "FAIL: error get users list", err: res.err})
			return
		}
		sendSuccess(c, http.StatusOK, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
}

func (h *Handler) GetUserById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
		return
	}
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		user, err := h.storage.LoadUser(ctx, int(id))
		select {
		case resultChan <- Result{data: user, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusBadRequest, Result{data: "FAIL: error get user", err: res.err})
			return
		}
		sendSuccess(c, http.StatusOK, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
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
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid hash", err: err})
		return
	}
	newUser.Password = ""
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		id, err := h.storage.SaveNewUser(ctx, newUser)
		select {
		case resultChan <- Result{data: id, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "FAIL: error create user", err: res.err})
			return
		}
		sendSuccess(c, http.StatusCreated, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
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
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid hash", err: err})
		return
	}
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		outId, err := h.storage.UpdateUserPassword(ctx, updateUser)
		select {
		case resultChan <- Result{data: outId, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "FAIL: error update password", err: res.err})
			return
		}
		sendSuccess(c, http.StatusCreated, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
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
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		outId, err := h.storage.UpdateUserEmail(ctx, updateUser)
		select {
		case resultChan <- Result{data: outId, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "FAIL: error update email", err: res.err})
			return
		}
		sendSuccess(c, http.StatusCreated, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
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
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		outId, err := h.storage.UpdateUserRoleId(ctx, updateUser)
		select {
		case resultChan <- Result{data: outId, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "FAIL: error update role", err: res.err})
			return
		}
		sendSuccess(c, http.StatusCreated, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}

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
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		err := h.storage.DeleteUser(ctx, deleteUser.Id)
		select {
		case resultChan <- Result{data: nil, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "FAIL: error delete user", err: res.err})
			return
		}
		sendSuccess(c, http.StatusCreated, Result{data: "User by delete", err: nil})
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
}
