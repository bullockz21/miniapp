package handler_repository

import "github.com/gin-gonic/gin"

type Handler interface {
	Register(r *gin.Engine)
}
