package handlers

import (
	"miniapp/internal/dto"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// func (h *Handler) GetMenuList(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	type result struct {
// 		menu []dto.MenuDTO
// 		err  error
// 	}
// 	resultChan := make(chan result, 1)
// 	defer close(resultChan)
// 	select {
// 	case sem <- struct{}{}:
// 		defer func() { <-sem }()
// 	case <-time.After(3 * time.Second):
// 		c.JSON(http.StatusGatewayTimeout, gin.H{
// 			"error": "request timeout",
// 		})
// 		return
// 	}
// 	go func() {
// 		menu, err := h.storage.LoadAllMenu(ctx)
// 		resultChan <- result{menu: menu, err: err}
// 	}()
// 	select {
// 	case res := <-resultChan:
// 		if res.err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error":   "FAIL: error get menu",
// 				"details": res.err.Error(),
// 			})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"status": "OK",
// 			"data":   res.menu,
// 		})
// 	case <-time.After(10 * time.Second):
// 		c.JSON(http.StatusGatewayTimeout, gin.H{
// 			"error": "request timeout",
// 		})
// 	case <-ctx.Done():
// 		c.JSON(http.StatusRequestTimeout, gin.H{
// 			"error": "request cancelled",
// 		})
// 	}
// }

func (h *Handler) GetOrdersByOrderNum(c *gin.Context) {
	ctx := c.Request.Context()
	type result struct {
		orders []dto.OrderDTO
		err    error
	}
	resultChan := make(chan result, 1)
	defer close(resultChan)
	select {
	case sem <- struct{}{}:
		defer func() { <-sem }()
	case <-time.After(3 * time.Second):
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"error": "request timeout",
		})
		return
	}
	go func() {
		idStr := c.Params.ByName("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			resultChan <- result{orders: nil, err: err}
			return
		}
		orders, err := h.storage.LoadOrdersByOrderNum(ctx, int(id))
		resultChan <- result{orders: orders, err: err}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "FAIL: error get menu position",
				"details": res.err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"data":   res.orders,
		})
	case <-time.After(10 * time.Second):
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"error": "request timeout",
		})
	case <-ctx.Done():
		c.JSON(http.StatusRequestTimeout, gin.H{
			"error": "request cancelled",
		})
	}
}

func (h *Handler) CreateOrder(c *gin.Context) {
	ctx := c.Request.Context()
	type result struct {
		id  int
		err error
	}
	resultChan := make(chan result, 1)
	defer close(resultChan)
	select {
	case sem <- struct{}{}:
		defer func() { <-sem }()
	case <-time.After(3 * time.Second):
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"error": "request timeout",
		})
		return
	}
	go func() {
		newOrder := dto.OrderDTO{}
		err := c.ShouldBindJSON(&newOrder)
		if err != nil {
			resultChan <- result{id: 0, err: err}
			return
		}
		id, err := h.storage.SaveNewOrder(ctx, newOrder)
		resultChan <- result{id: id, err: err}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "FAIL: error create customer position",
				"details": res.err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"data":   res.id,
		})
	case <-time.After(10 * time.Second):
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"error": "request timeout",
		})
	case <-ctx.Done():
		c.JSON(http.StatusRequestTimeout, gin.H{
			"error": "request cancelled",
		})
	}
}

func (h *Handler) UpdateOrder(c *gin.Context) {
	ctx := c.Request.Context()
	type result struct {
		id  int
		err error
	}
	resultChan := make(chan result, 1)
	defer close(resultChan)
	select {
	case sem <- struct{}{}:
		defer func() { <-sem }()
	case <-time.After(3 * time.Second):
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"error": "request timeout",
		})
		return
	}
	go func() {
		UpdateOrder := dto.OrderDTO{}
		idStr := c.Params.ByName("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			resultChan <- result{id: 0, err: err}
			return
		}
		err = c.ShouldBindJSON(&UpdateOrder)
		if err != nil {
			resultChan <- result{id: 0, err: err}
			return
		}
		// updateMenu.Id = int(id)
		outId, err := h.storage.UpdateOrder(ctx, UpdateOrder, int(id))
		resultChan <- result{id: outId, err: err}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "FAIL: error update customer position",
				"details": res.err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"data":   res.id,
		})
	case <-time.After(10 * time.Second):
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"error": "request timeout",
		})
	case <-ctx.Done():
		c.JSON(http.StatusRequestTimeout, gin.H{
			"error": "request cancelled",
		})
	}
}

// func (h *Handler) DeleteMenuById(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	type result struct {
// 		err error
// 	}
// 	resultChan := make(chan result, 1)
// 	defer close(resultChan)
// 	select {
// 	case sem <- struct{}{}:
// 		defer func() { <-sem }()
// 	case <-time.After(3 * time.Second):
// 		c.JSON(http.StatusGatewayTimeout, gin.H{
// 			"error": "request timeout",
// 		})
// 		return
// 	}
// 	go func() {
// 		idStr := c.Params.ByName("id")
// 		id, err := strconv.ParseInt(idStr, 10, 64)
// 		if err != nil {
// 			resultChan <- result{err: err}
// 			return
// 		}
// 		err = h.storage.DeleteMenuPosition(ctx, int(id))
// 		resultChan <- result{err: err}
// 	}()
// 	select {
// 	case res := <-resultChan:
// 		if res.err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error":   "FAIL: error delete menu position",
// 				"details": res.err.Error(),
// 			})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"status": "OK",
// 			"data":   "delete Menu position",
// 		})
// 	case <-time.After(10 * time.Second):
// 		c.JSON(http.StatusGatewayTimeout, gin.H{
// 			"error": "request timeout",
// 		})
// 	case <-ctx.Done():
// 		c.JSON(http.StatusRequestTimeout, gin.H{
// 			"error": "request cancelled",
// 		})
// 	}
// }
