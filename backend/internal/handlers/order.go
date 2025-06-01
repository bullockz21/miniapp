package handlers

import (
	"context"
	"miniapp/internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetOrdersByOrderNum(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), requestTimeout)
	defer cancel()
	if !checkSemaphore(c, ctx) {
		return
	}
	defer releaseSemaphore()

	idStr := c.Params.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, "invalid id", err)
		return
	}

	type result struct {
		orders []dto.OrderDTO
		err    error
	}
	resultChan := make(chan result, 1)
	go func() {
		defer close(resultChan)
		orders, err := h.storage.LoadOrdersByOrderNum(ctx, int(id))
		select {
		case resultChan <- result{orders: orders, err: err}:
		case <-ctx.Done():
		}
	}()

	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, "FAIL: error get order position", res.err)
			return
		}
		sendSuccess(c, http.StatusCreated, res.orders)
	case <-ctx.Done():
		handleContextError(c, ctx)
	}
}

func (h *Handler) CreateOrder(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), requestTimeout)
	defer cancel()
	if !checkSemaphore(c, ctx) {
		return
	}
	defer releaseSemaphore()
	newOrder := dto.OrderDTO{}
	err := c.ShouldBindJSON(&newOrder)
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
		id, err := h.storage.SaveNewOrder(ctx, newOrder)
		select {
		case resultChan <- result{id: id, err: err}:
		case <-ctx.Done():
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, "FAIL: error create order position", res.err)
			return
		}
		sendSuccess(c, http.StatusCreated, res.id)
	case <-ctx.Done():
		handleContextError(c, ctx)
	}
}

func (h *Handler) UpdateOrder(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), requestTimeout)
	defer cancel()
	if !checkSemaphore(c, ctx) {
		return
	}
	defer releaseSemaphore()

	UpdateOrder := dto.OrderDTO{}
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, "invalid id", err)
		return
	}
	err = c.ShouldBindJSON(&UpdateOrder)
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
		id, err := h.storage.UpdateOrder(ctx, UpdateOrder, int(id))
		select {
		case resultChan <- result{id: id, err: err}:
		case <-ctx.Done():
		}
	}()

	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, "FAIL: error update order position", res.err)
			return
		}
		sendSuccess(c, http.StatusCreated, res.id)
	case <-ctx.Done():
		handleContextError(c, ctx)
	}
}
