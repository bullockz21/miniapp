package handlers

import (
	"context"
	"miniapp/internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCustomerById(c *gin.Context) {
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
		result dto.CustomerDTO
		err    error
	}
	resultChan := make(chan result, 1)
	go func() {
		defer close(resultChan)
		customer, err := h.storage.LoadCustomer(ctx, int(id))
		select {
		case resultChan <- result{result: customer, err: err}:
		case <-ctx.Done():
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, "FAIL: error get menu position", res.err)
			return
		}
		sendSuccess(c, http.StatusOK, res.result)
	case <-ctx.Done():
		handleContextError(c, ctx)
	}
}

func (h *Handler) CreateCustomer(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), requestTimeout)
	defer cancel()
	if !checkSemaphore(c, ctx) {
		return
	}
	defer releaseSemaphore()
	type result struct {
		id  int
		err error
	}
	newCustomer := dto.CustomerDTO{}
	err := c.ShouldBindJSON(&newCustomer)
	if err != nil {
		sendError(c, http.StatusBadRequest, "invalid request body", err)
		return
	}
	resultChan := make(chan result, 1)
	go func() {
		defer close(resultChan)
		id, err := h.storage.SaveNewCustomer(ctx, newCustomer)
		select {
		case resultChan <- result{id: id, err: err}:
		case <-ctx.Done():
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, "FAIL: error create customer", res.err)
			return
		}
		sendSuccess(c, http.StatusCreated, res.id)
	case <-ctx.Done():
		handleContextError(c, ctx)
	}
}

func (h *Handler) UpdateCustomer(c *gin.Context) {
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
	UpdateCustomer := dto.CustomerDTO{}
	err = c.ShouldBindJSON(&UpdateCustomer)
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
		outId, err := h.storage.UpdateCustomer(ctx, UpdateCustomer, int(id))
		select {
		case resultChan <- result{id: outId, err: err}:
		case <-ctx.Done():
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, "FAIL: error update customer", res.err)
			return
		}
		sendSuccess(c, http.StatusOK, res.id)
	case <-ctx.Done():
		handleContextError(c, ctx)
	}
}
