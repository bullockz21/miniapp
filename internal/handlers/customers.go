package handlers

import (
	"miniapp/internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCustomerById godoc
//
//	@Summary		Get Customer
//	@Description	Get customers data by id
//	@Tags			customer
//	@Produce		json
//	@Param			id	path		int	true	"Customer ID"
//	@Success		200	{object}	dto.CustomerDTO
//	@Failure		400	{object}	handlers.Error
//	@Failure		401	{object}	handlers.Error
//	@Router			/customers/:id [get]
func (h *Handler) GetCustomerById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		SendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
		return
	}
	customer, err := h.storage.LoadCustomer(ctx, int(id))
	if err != nil {
		SendError(c, http.StatusNotFound, Result{data: "FAIL: error get menu position", err: err})
	}
	c.JSON(http.StatusOK, customer)
	// SendSuccess(c, http.StatusOK, Result{data: customer, err: err})
}

// CreateCustomer godoc
//
//	@Summary		Create customer
//	@Description	Create customer
//	@Tags			customer
//	@Accept			json
//	@Produce		json
//	@Param			account	body		dto.CustomerDTO	true	"Customer create data"
//	@Success		200		{object}	handlers.Success
//	@Failure		400		{object}	handlers.Error
//	@Failure		401		{object}	handlers.Error
//	@Router			/customers [post]
func (h *Handler) CreateCustomer(c *gin.Context) {
	ctx := c.Request.Context()
	newCustomer := dto.CustomerDTO{}
	err := c.ShouldBindJSON(&newCustomer)
	if err != nil {
		SendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	id, err := h.storage.SaveNewCustomer(ctx, newCustomer)
	if err != nil {
		SendError(c, http.StatusNotFound, Result{data: "FAIL: error create customer", err: err})
		return
	}
	SendSuccess(c, http.StatusCreated, Result{data: id, err: nil})
}

// UpdateCustomer godoc
//
//	@Summary		Update customer
//	@Description	Update customer by ID
//	@Tags			customer
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Customer ID"
//	@Param			account	body		dto.WebUpdateCustomerDTO	true	"Customer create data"
//	@Success		200		{object}	handlers.Success
//	@Failure		400		{object}	handlers.Error
//	@Failure		401		{object}	handlers.Error
//	@Router			/customer/:id [patch]
func (h *Handler) UpdateCustomer(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		SendError(c, http.StatusBadRequest, Result{data: "incorrect id", err: err})
		return
	}
	UpdateCustomer := dto.CustomerDTO{}
	err = c.ShouldBindJSON(&UpdateCustomer)
	if err != nil {
		SendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	outId, err := h.storage.UpdateCustomer(ctx, UpdateCustomer, int(id))
	if err != nil {
		SendError(c, http.StatusNotFound, Result{data: "FAIL: error update customer", err: err})
		return
	}
	SendSuccess(c, http.StatusCreated, Result{data: outId, err: nil})
}
