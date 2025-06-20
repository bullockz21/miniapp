package handlers

import (
	"fmt"
	"miniapp/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCustomerById godoc
//
//	@Summary		Get Customer
//	@Description	Get customers data by id
//	@Tags			customer
//	@Produce		json
//	@Param			id	path		int	true	"Customer ID"
//	@Success		200	{object}	dto.WebLoadCustomerDTO
//	@Failure		400	{object}	handlers.Error
//	@Failure		401	{object}	handlers.Error
//	@Router			/customer/{id} [get]
func (h *Handler) GetCustomerById(c *gin.Context) {

	idStr := c.Params.ByName("id")
	id := 0
	_, err := fmt.Sscanf(idStr, "%d", &id)

	if err != nil {
		SendError(c, http.StatusNotFound, Result{data: "incorrect id", err: err})
		return
	}
	customer, err := h.customer.Load(id)
	if err != nil {
		SendError(c, http.StatusNotFound, Result{data: "FAIL: error get customer", err: err})
		return
	}
	c.JSON(http.StatusOK, customer)
	// SendSuccess(c, http.StatusOK, Result{data: customer, err: err})
}

// GetCustomerList godoc
//
//	@Summary		Get Customer List
//	@Description	Get list of all customers
//	@Tags			customer
//	@Produce		json
//	@Success		200	{object}	[]dto.WebLoadCustomerDTO
//	@Failure		400	{object}	handlers.Error
//	@Failure		401	{object}	handlers.Error
//	@Router			/customer [get]
func (h *Handler) GetCustomerList(c *gin.Context) {

	customers, err := h.customer.LoadList()
	if err != nil {
		SendError(c, http.StatusNotFound, Result{data: "FAIL: error get customer list", err: err})
		return
	}
	c.JSON(http.StatusOK, customers)
	// SendSuccess(c, http.StatusOK, Result{data: customer, err: err})
}

// CreateCustomer godoc
//
//	@Summary		Create customer
//	@Description	Create customer
//	@Tags			customer
//	@Accept			json
//	@Produce		json
//	@Param			account	body		dto.WebCreateCustomerDTO	true	"Customer create data"
//	@Success		200		{object}	handlers.Success
//	@Failure		400		{object}	handlers.Error
//	@Failure		401		{object}	handlers.Error
//	@Router			/customer [post]
func (h *Handler) CreateCustomer(c *gin.Context) {
	newCustomer := dto.WebCreateCustomerDTO{}
	err := c.ShouldBindJSON(&newCustomer)
	if err != nil {
		SendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	id, err := h.customer.Create(newCustomer)
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
//	@Param			account	body		dto.WebUpdateCustomerDTO	true	"Customer update data"
//	@Success		200		{object}	handlers.Success
//	@Failure		400		{object}	handlers.Error
//	@Failure		401		{object}	handlers.Error
//	@Router			/customer/ [patch]
func (h *Handler) UpdateCustomer(c *gin.Context) {

	UpdateCustomer := dto.WebUpdateCustomerDTO{}
	err := c.ShouldBindJSON(&UpdateCustomer)
	if err != nil {
		SendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	outId, err := h.customer.Update(UpdateCustomer)
	if err != nil {
		SendError(c, http.StatusNotFound, Result{data: "FAIL: error update customer", err: err})
		return
	}
	SendSuccess(c, http.StatusCreated, Result{data: outId, err: nil})
}
