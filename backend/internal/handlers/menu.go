package handlers

import (
	"miniapp/internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMenuList(c *gin.Context) {
	ctx := c.Request.Context()
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		menu, err := h.storage.LoadAllMenu(ctx)
		select {
		case resultChan <- Result{data: menu, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "FAIL: error get menu", err: res.err})
			return
		}
		sendSuccess(c, http.StatusOK, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
}

func (h *Handler) GetMenuById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
		return
	}
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		menu, err := h.storage.LoadMenuPosition(ctx, int(id))
		select {
		case resultChan <- Result{data: menu, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "FAIL: error get menu position", err: res.err})
			return
		}
		sendSuccess(c, http.StatusOK, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
}

func (h *Handler) CreateMenu(c *gin.Context) {
	ctx := c.Request.Context()
	newMenu := dto.MenuDTO{}
	err := c.ShouldBindJSON(&newMenu)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		id, err := h.storage.SaveNewMenuPosition(ctx, newMenu)
		select {
		case resultChan <- Result{data: id, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "FAIL: error create menu position", err: res.err})
			return
		}
		sendSuccess(c, http.StatusCreated, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
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
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		id, err := h.storage.UpdateMenuPosition(ctx, updateMenu, int(id))
		select {
		case resultChan <- Result{data: id, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "FAIL: error update menu position", err: res.err})
			return
		}
		sendSuccess(c, http.StatusCreated, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
}

func (h *Handler) DeleteMenuById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
		return
	}
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		err = h.storage.DeleteMenuPosition(ctx, int(id))
		select {
		case resultChan <- Result{data: nil, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "FAIL: error delete menu position", err: res.err})
			return
		}
		sendSuccess(c, http.StatusCreated, Result{data: "delete Menu position", err: nil})
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
}
