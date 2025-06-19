package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	data any
	err  error
}

type Error struct {
	Status  string `json:"status" example:"status text"`
	Error   string `json:"error" example:"erros text"`
	Details string `json:"details" example:"details text"`
}

type Success struct {
	Status string `json:"status"`
	Data   Result `json:"data"`
}

func SendError(c *gin.Context, status int, res Result) {
	c.JSON(status, gin.H{
		"status":  status,
		"error":   res.data,
		"details": res.err.Error(),
	})
}

func SendSuccess(c *gin.Context, status int, res Result) {
	c.JSON(status, gin.H{
		"status": status,
		"data":   res.data,
	})
}

func handleContextError(c *gin.Context, ctx context.Context) {
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		SendError(c, http.StatusGatewayTimeout, Result{data: "request timeout", err: nil})
	} else {
		SendError(c, http.StatusRequestTimeout, Result{data: "request cancelled", err: nil})
	}
}
