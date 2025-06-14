package handlers

import (
	"miniapp/internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMenuList(c *gin.Context) {
	ctx := c.Request.Context()
	menu, err := h.storage.LoadAllMenu(ctx)
	if err != nil {
		sendError(c, http.StatusNotFound, Result{data: "FAIL: error get menu", err: err})
		return
	}
	sendSuccess(c, http.StatusOK, Result{data: menu, err: nil})

}

func (h *Handler) GetMenuById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
		return
	}
	menu, err := h.storage.LoadMenuPosition(ctx, int(id))
	if err != nil {
		sendError(c, http.StatusNotFound, Result{data: "FAIL: error get by id menu", err: err})
		return
	}
	sendSuccess(c, http.StatusOK, Result{data: menu, err: nil})
}

func (h *Handler) CreateMenu(c *gin.Context) {
	ctx := c.Request.Context()
	newMenu := dto.MenuDTO{}
	err := c.ShouldBindJSON(&newMenu)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	id, err := h.storage.SaveNewMenuPosition(ctx, newMenu)
	if err != nil {
		sendError(c, http.StatusNotFound, Result{data: "FAIL: error create menu", err: err})
		return
	}
	sendSuccess(c, http.StatusOK, Result{data: id, err: nil})
}

func (h *Handler) UpdateMenuPosition(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
		return
	}
	updateMenu := dto.MenuDTO{}
	err = c.ShouldBindJSON(&updateMenu)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	outId, err := h.storage.UpdateMenuPosition(ctx, updateMenu, int(id))
	if err != nil {
		sendError(c, http.StatusNotFound, Result{data: "FAIL: error update menu", err: err})
		return
	}
	sendSuccess(c, http.StatusOK, Result{data: outId, err: nil})
}

func (h *Handler) DeleteMenuById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
		return
	}
	err = h.storage.DeleteMenuPosition(ctx, int(id))
	if err != nil {
		sendError(c, http.StatusNotFound, Result{data: "FAIL: error delete menu", err: err})
		return
	}
	sendSuccess(c, http.StatusOK, Result{data: "Position by DELETE", err: nil})
}
