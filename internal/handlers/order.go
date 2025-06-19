package handlers

// import (
// 	"miniapp/internal/dto"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// func (h *Handler) GetOrdersByOrderId(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "invalid id", err: err})
// 		return
// 	}
// 	orders, err := h.storage.LoadOrdersByOrderNum(ctx, int(id))
// 	if err != nil {
// 		SendError(c, http.StatusNotFound, Result{data: "FAIL: error get order position", err: err})
// 		return
// 	}
// 	SendSuccess(c, http.StatusCreated, Result{data: orders, err: nil})
// }

// func (h *Handler) GetAllOrders(c *gin.Context) {
// 	ctx := c.Request.Context()

// 	orders, err := h.storage.LoadOrders(ctx)

// 	if err != nil {
// 		SendError(c, http.StatusNotFound, Result{data: "FAIL: error get all order", err: err})
// 		return
// 	}
// 	SendSuccess(c, http.StatusCreated, Result{data: orders, err: nil})
// }

// func (h *Handler) CreateOrder(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	newOrder := []dto.OrderDTO{}
// 	err := c.ShouldBindJSON(&newOrder)
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
// 		return
// 	}
// 	type Ord struct {
// 		Id    int
// 		Count int
// 	}

// 	id := 0
// 	count := 0
// 	for _, order := range newOrder {
// 		id, err = h.storage.SaveNewOrder(ctx, order)
// 		if err != nil {
// 			return
// 		}
// 		count++
// 	}

// 	if err != nil {
// 		SendError(c, http.StatusNotFound, Result{data: "FAIL: error create order", err: err})
// 		return
// 	}
// 	SendSuccess(c, http.StatusCreated, Result{data: Ord{Id: id, Count: count}, err: nil})
// }

// func (h *Handler) UpdateOrder(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	idStr := c.Params.ByName("id")
// 	id, err := strconv.ParseInt(idStr, 10, 64)
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "invalid id", err: err})
// 		return
// 	}
// 	UpdateOrder := dto.OrderDTO{}
// 	err = c.ShouldBindJSON(&UpdateOrder)
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
// 		return
// 	}
// 	outId, err := h.storage.UpdateOrder(ctx, UpdateOrder, int(id))
// 	if err != nil {
// 		SendError(c, http.StatusNotFound, Result{data: "FAIL: error update order", err: err})
// 		return
// 	}
// 	SendSuccess(c, http.StatusCreated, Result{data: outId, err: nil})
// }
