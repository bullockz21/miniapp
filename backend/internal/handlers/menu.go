package handlers

import (
	"miniapp/internal/dto"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMenuList(c *gin.Context) {
	ctx := c.Request.Context()
	type result struct {
		menu []dto.MenuDTO
		err  error
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
		menu, err := h.storage.LoadAllMenu(ctx)
		resultChan <- result{menu: menu, err: err}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "FAIL: error get menu",
				"details": res.err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"data":   res.menu,
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

func (h *Handler) GetMenuById(c *gin.Context) {
	ctx := c.Request.Context()
	type result struct {
		menu dto.MenuDTO
		err  error
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
			resultChan <- result{menu: dto.MenuDTO{}, err: err}
			return
		}
		menu, err := h.storage.LoadMenuPosition(ctx, int(id))
		resultChan <- result{menu: menu, err: err}
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
			"data":   res.menu,
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

func (h *Handler) CreateMenu(c *gin.Context) {
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
		newMenu := dto.MenuDTO{}
		err := c.ShouldBindJSON(&newMenu)
		if err != nil {
			resultChan <- result{id: 0, err: err}
			return
		}
		id, err := h.storage.SaveNewMenuPosition(ctx, newMenu)
		resultChan <- result{id: id, err: err}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "FAIL: error create menu position",
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

func (h *Handler) UpdateMenuPosition(c *gin.Context) {
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
		updateMenu := dto.MenuDTO{}
		idStr := c.Params.ByName("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			resultChan <- result{id: 0, err: err}
			return
		}
		err = c.ShouldBindJSON(&updateMenu)
		if err != nil {
			resultChan <- result{id: 0, err: err}
			return
		}
		// updateMenu.Id = int(id)
		outId, err := h.storage.UpdateMenuPosition(ctx, updateMenu, int(id))
		resultChan <- result{id: outId, err: err}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "FAIL: error update menu position",
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

func (h *Handler) DeleteMenuById(c *gin.Context) {
	ctx := c.Request.Context()
	type result struct {
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
		idStr := c.Params.ByName("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			resultChan <- result{err: err}
			return
		}
		err = h.storage.DeleteMenuPosition(ctx, int(id))
		resultChan <- result{err: err}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "FAIL: error delete menu position",
				"details": res.err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"data":   "delete Menu position",
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
