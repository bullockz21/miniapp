package handlers

import (
	"miniapp/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// import (
// 	"miniapp/internal/domain/user"
// 	"miniapp/internal/dto"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// func (h *Handler) GetUsersList(c *gin.Context) {
// 	ctx := c.Request.Context()

// 	users, err := h.storage.LoadAllUsers(ctx)

// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "FAIL: error get users list", err: err})
// 		return
// 	}
// 	SendSuccess(c, http.StatusOK, Result{data: users, err: nil})

// }

// func (h *Handler) GetUserById(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
// 		return
// 	}

// 	user, err := h.storage.LoadUser(ctx, int(id))

// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "FAIL: error get user by id", err: err})
// 		return
// 	}
// 	SendSuccess(c, http.StatusOK, Result{data: user, err: nil})
// }

// CreateCustomer godoc
//
//	@Summary		Create user
//	@Description	Create user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			account	body		dto.WebCreateUserDTO	true	"User create data"
//	@Success		200		{object}	handlers.Success
//	@Failure		400		{object}	handlers.Error
//	@Failure		401		{object}	handlers.Error
//	@Router			/user [post]
func (h *Handler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	newUser := dto.WebCreateUserDTO{}
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		SendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}

	id, err := h.user.Create(ctx, newUser)

	if err != nil {
		SendError(c, http.StatusBadRequest, Result{data: "FAIL: error create user", err: err})
		return
	}
	SendSuccess(c, http.StatusOK, Result{data: id, err: nil})
}

// func (h *Handler) UpdateUserPasswordById(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
// 		return
// 	}
// 	updateUser := dto.UserDTO{}
// 	err = c.ShouldBindJSON(&updateUser)
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
// 		return
// 	}
// 	updateUser.Id = int(id)
// 	updateUser.PasswordHash, err = user.CreatePasswordHash(updateUser.Password)
// 	updateUser.Password = ""
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "invalid hash", err: err})
// 		return
// 	}

// 	outId, err := h.storage.UpdateUserPassword(ctx, updateUser)

// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "FAIL: error update user password", err: err})
// 		return
// 	}
// 	SendSuccess(c, http.StatusOK, Result{data: outId, err: nil})
// }

// func (h *Handler) UpdateUserEmailById(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
// 		return
// 	}
// 	updateUser := dto.UserDTO{}
// 	err = c.ShouldBindJSON((&updateUser))
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
// 		return
// 	}
// 	updateUser.Id = int(id)

// 	outId, err := h.storage.UpdateUserEmail(ctx, updateUser)

// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "FAIL: error update user email", err: err})
// 		return
// 	}
// 	SendSuccess(c, http.StatusOK, Result{data: outId, err: nil})
// }

// func (h *Handler) UpdateUserRoleIdById(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
// 		return
// 	}
// 	updateUser := dto.UserDTO{}
// 	err = c.ShouldBindJSON((&updateUser))
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
// 		return
// 	}
// 	updateUser.Id = int(id)

// 	outId, err := h.storage.UpdateUserRoleId(ctx, updateUser)

// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "FAIL: error update users role", err: err})
// 		return
// 	}
// 	SendSuccess(c, http.StatusOK, Result{data: outId, err: nil})

// }

// func (h *Handler) DeleteUserById(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
// 		return
// 	}
// 	deleteUser := dto.UserDTO{}
// 	err = c.ShouldBindJSON((&deleteUser))
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
// 		return
// 	}
// 	deleteUser.Id = int(id)

// 	err = h.storage.DeleteUser(ctx, deleteUser.Id)

// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "FAIL: error delete users", err: err})
// 		return
// 	}
// 	SendSuccess(c, http.StatusOK, Result{data: "User by DELETE", err: nil})
// }
