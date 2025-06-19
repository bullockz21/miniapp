package handlers

// import (
// 	"miniapp/internal/dto"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// // GetMenuList godoc
// //
// //	@Summary		Get menu list
// //	@Description	Get menu list
// //	@Tags			menu
// //	@Produce		json
// //	@Success		200	{object}	[]dto.MenuDTO
// //	@Failure		400	{object}	handlers.Error
// //	@Failure		401	{object}	handlers.Error
// //	@Router			/menu [get]
// func (h *Handler) GetMenuList(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	menu, err := h.storage.LoadAllMenu(ctx)
// 	if err != nil {
// 		SendError(c, http.StatusNotFound, Result{data: "FAIL: error get menu", err: err})
// 		return
// 	}
// 	c.JSON(http.StatusOK, menu)
// 	// SendSuccess(c, http.StatusOK, Result{data: menu, err: nil})
// }

// // GetMenuById godoc
// //
// //	@Summary		Get menu position
// //	@Description	Get menu position by id
// //	@Tags			menu
// //	@Produce		json
// //	@Param			id	path		int	true	"Menu ID"
// //	@Success		200	{object}	dto.MenuDTO
// //	@Failure		400	{object}	handlers.Error
// //	@Failure		401	{object}	handlers.Error
// //	@Router			/menu/:id [get]
// func (h *Handler) GetMenuById(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
// 		return
// 	}
// 	menu, err := h.storage.LoadMenuPosition(ctx, int(id))
// 	if err != nil {
// 		SendError(c, http.StatusNotFound, Result{data: "FAIL: error get by id menu", err: err})
// 		return
// 	}
// 	c.JSON(http.StatusOK, menu)
// 	// SendSuccess(c, http.StatusOK, Result{data: menu, err: nil})
// }

// // CreateMenu godoc
// //
// //	@Summary		Create Menu position
// //	@Description	Create Menu position
// //	@Tags			menu
// //	@Accept			json
// //	@Produce		json
// //	@Param			account	body		dto.WebCreateMenuDTO	true	"Menu position data"
// //	@Success		200		{object}	dto.MenuDTO
// //	@Failure		400		{object}	handlers.Error
// //	@Failure		401		{object}	handlers.Error
// //	@Router			/menu [post]
// func (h *Handler) CreateMenu(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	newMenu := dto.MenuDTO{}
// 	err := c.ShouldBindJSON(&newMenu)
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
// 		return
// 	}
// 	id, err := h.storage.SaveNewMenuPosition(ctx, newMenu)
// 	if err != nil {
// 		SendError(c, http.StatusNotFound, Result{data: "FAIL: error create menu", err: err})
// 		return
// 	}
// 	SendSuccess(c, http.StatusOK, Result{data: id, err: nil})
// }

// // UpdateMenuPosition godoc
// //
// //	@Summary		Update menu position
// //	@Description	Update menu position by ID
// //	@Tags			menu
// //	@Accept			json
// //	@Produce		json
// //	@Param			id		path		int						true	"Menu position ID"
// //	@Param			account	body		dto.WebCreateMenuDTO	true	"Menu position data"
// //	@Success		200		{object}	dto.MenuDTO
// //	@Failure		400		{object}	handlers.Error
// //	@Failure		401		{object}	handlers.Error
// //	@Router			/menu/:id [patch]
// func (h *Handler) UpdateMenuPosition(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
// 		return
// 	}
// 	updateMenu := dto.MenuDTO{}
// 	err = c.ShouldBindJSON(&updateMenu)
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
// 		return
// 	}
// 	outId, err := h.storage.UpdateMenuPosition(ctx, updateMenu, int(id))
// 	if err != nil {
// 		SendError(c, http.StatusNotFound, Result{data: "FAIL: error update menu", err: err})
// 		return
// 	}
// 	SendSuccess(c, http.StatusOK, Result{data: outId, err: nil})
// }

// // DeleteMenuById godoc
// //
// //	@Summary		Delete menu position
// //	@Description	Delete menu position by ID
// //	@Tags			menu
// //	@Produce		json
// //	@Param			id	path		int	true	"Menu position ID"
// //	@Success		200	{object}	handlers.Success
// //	@Failure		400	{object}	handlers.Error
// //	@Failure		401	{object}	handlers.Error
// //	@Router			/menu/:id [delete]
// func (h *Handler) DeleteMenuById(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
// 		return
// 	}
// 	err = h.storage.DeleteMenuPosition(ctx, int(id))
// 	if err != nil {
// 		SendError(c, http.StatusNotFound, Result{data: "FAIL: error delete menu", err: err})
// 		return
// 	}
// 	SendSuccess(c, http.StatusOK, Result{data: "Position by DELETE", err: nil})
// }
