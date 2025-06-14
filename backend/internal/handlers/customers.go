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
	customer, err := h.storage.LoadCustomer(ctx, int(id))
	if err != nil {
		sendError(c, http.StatusNotFound, Result{data: "FAIL: error get menu position", err: err})
	}
	sendSuccess(c, http.StatusOK, Result{data: customer, err: err})
}

func (h *Handler) CreateCustomer(c *gin.Context) {
	ctx := c.Request.Context()
	newCustomer := dto.CustomerDTO{}
	err := c.ShouldBindJSON(&newCustomer)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	id, err := h.storage.SaveNewCustomer(ctx, newCustomer)
	if err != nil {
		sendError(c, http.StatusNotFound, Result{data: "FAIL: error create customer", err: err})
		return
	}
	sendSuccess(c, http.StatusCreated, Result{data: id, err: nil})
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
	outId, err := h.storage.UpdateCustomer(ctx, UpdateCustomer, int(id))
	if err != nil {
		sendError(c, http.StatusNotFound, Result{data: "FAIL: error update customer", err: err})
		return
	}
	sendSuccess(c, http.StatusCreated, Result{data: outId, err: nil})
}
