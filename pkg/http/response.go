package http

import (
	"github.com/gin-gonic/gin"
	"go-jm-core/pkg/exception"
)

func Error(c *gin.Context, exception *exception.AppException) {
	c.Set("exception", exception)
	c.Status(exception.Code)
}

func Response(c *gin.Context, status int, data any) {
	c.Set("data", data)
}
