package handlers

import (
	"context"
	"miniapp/internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMenuList(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), requestTimeout)
	defer cancel()
	if !checkSemaphore(c, ctx) {
		return
	}
	defer releaseSemaphore()
	type result struct {
		menu []dto.MenuDTO
		err  error
	}
	resultChan := make(chan result, 1)
	go func() {
		defer close(resultChan)
		menu, err := h.storage.LoadAllMenu(ctx)
		select {
		case resultChan <- result{menu: menu, err: err}:
		case <-ctx.Done():
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, "FAIL: error get menu", res.err)
			return
		}
		sendSuccess(c, http.StatusOK, res.menu)
	case <-ctx.Done():
		handleContextError(c, ctx)
	}
}

func (h *Handler) GetMenuById(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), requestTimeout)
	defer cancel()
	if !checkSemaphore(c, ctx) {
		return
	}
	defer releaseSemaphore()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, "incorrect id", err)
		return
	}
	type result struct {
		menu dto.MenuDTO
		err  error
	}
	resultChan := make(chan result, 1)
	go func() {
		defer close(resultChan)
		menu, err := h.storage.LoadMenuPosition(ctx, int(id))
		select {
		case resultChan <- result{menu: menu, err: err}:
		case <-ctx.Done():
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, "FAIL: error get menu position", res.err)
			return
		}
		sendSuccess(c, http.StatusOK, res.menu)
	case <-ctx.Done():
		handleContextError(c, ctx)
	}
}

func (h *Handler) CreateMenu(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), requestTimeout)
	defer cancel()
	if !checkSemaphore(c, ctx) {
		return
	}
	defer releaseSemaphore()
	newMenu := dto.MenuDTO{}
	err := c.ShouldBindJSON(&newMenu)
	if err != nil {
		sendError(c, http.StatusBadRequest, "invalid request body", err)
		return
	}
	type result struct {
		id  int
		err error
	}
	resultChan := make(chan result, 1)
	go func() {
		defer close(resultChan)
		id, err := h.storage.SaveNewMenuPosition(ctx, newMenu)
		select {
		case resultChan <- result{id: id, err: err}:
		case <-ctx.Done():
		}
	}()

	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, "FAIL: error create menu position", res.err)
			return
		}
		sendSuccess(c, http.StatusCreated, res.id)
	case <-ctx.Done():
		handleContextError(c, ctx)
	}
}

func (h *Handler) UpdateMenuPosition(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), requestTimeout)
	defer cancel()
	if !checkSemaphore(c, ctx) {
		return
	}
	defer releaseSemaphore()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, "incorrect id", err)
		return
	}
	updateMenu := dto.MenuDTO{}
	err = c.ShouldBindJSON(&updateMenu)
	if err != nil {
		sendError(c, http.StatusBadRequest, "invalid request body", err)
		return
	}
	type result struct {
		id  int
		err error
	}
	resultChan := make(chan result, 1)
	go func() {
		defer close(resultChan)
		id, err := h.storage.UpdateMenuPosition(ctx, updateMenu, int(id))
		select {
		case resultChan <- result{id: id, err: err}:
		case <-ctx.Done():
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, "FAIL: error update menu position", res.err)
			return
		}
		sendSuccess(c, http.StatusCreated, res.id)
	case <-ctx.Done():
		handleContextError(c, ctx)
	}
}

func (h *Handler) DeleteMenuById(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), requestTimeout)
	defer cancel()
	if !checkSemaphore(c, ctx) {
		return
	}
	defer releaseSemaphore()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, "incorrect id", err)
		return
	}
	type result struct {
		err error
	}
	resultChan := make(chan result, 1)
	go func() {
		defer close(resultChan)
		err = h.storage.DeleteMenuPosition(ctx, int(id))
		select {
		case resultChan <- result{err: err}:
		case <-ctx.Done():
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, "FAIL: error delete menu position", res.err)
			return
		}
		sendSuccess(c, http.StatusCreated, "delete Menu position")
	case <-ctx.Done():
		handleContextError(c, ctx)
	}
}
