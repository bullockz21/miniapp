package handlers

import (
	"miniapp/internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCustomerById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
		return
	}
	resultChan := make(chan Result)
	go func() {
		defer close(resultChan)
		customer, err := h.storage.LoadCustomer(ctx, int(id))
		select {
		case resultChan <- Result{data: customer, err: err}:
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

func (h *Handler) CreateCustomer(c *gin.Context) {
	ctx := c.Request.Context()
	newCustomer := dto.CustomerDTO{}
	err := c.ShouldBindJSON(&newCustomer)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	resultChan := make(chan Result)
	go func() {
		defer close(resultChan)
		id, err := h.storage.SaveNewCustomer(ctx, newCustomer)
		select {
		case resultChan <- Result{data: id, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "FAIL: error create customer", err: res.err})
			return
		}
		sendSuccess(c, http.StatusCreated, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
}

func (h *Handler) UpdateCustomer(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
		return
	}
	UpdateCustomer := dto.CustomerDTO{}
	err = c.ShouldBindJSON(&UpdateCustomer)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	resultChan := make(chan Result)
	go func() {
		defer close(resultChan)
		outId, err := h.storage.UpdateCustomer(ctx, UpdateCustomer, int(id))
		select {
		case resultChan <- Result{data: outId, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "FAIL: error update customer", err: res.err})
			return
		}
		sendSuccess(c, http.StatusOK, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
}
