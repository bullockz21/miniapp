package handlers

import (
	"miniapp/internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetOrdersByOrderNum(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid id", err: err})
		return
	}
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		orders, err := h.storage.LoadOrdersByOrderNum(ctx, int(id))
		select {
		case resultChan <- Result{data: orders, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "FAIL: error get order position", err: res.err})
			return
		}
		sendSuccess(c, http.StatusCreated, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
}

func (h *Handler) CreateOrder(c *gin.Context) {
	ctx := c.Request.Context()
	newOrder := dto.OrderDTO{}
	err := c.ShouldBindJSON(&newOrder)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		id, err := h.storage.SaveNewOrder(ctx, newOrder)
		select {
		case resultChan <- Result{data: id, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "FAIL: error create order position", err: res.err})
			return
		}
		sendSuccess(c, http.StatusCreated, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
}

func (h *Handler) UpdateOrder(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid id", err: err})
		return
	}
	UpdateOrder := dto.OrderDTO{}
	err = c.ShouldBindJSON(&UpdateOrder)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		id, err := h.storage.UpdateOrder(ctx, UpdateOrder, int(id))
		select {
		case resultChan <- Result{data: id, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "FAIL: error update order position", err: res.err})
			return
		}
		sendSuccess(c, http.StatusCreated, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
}
