package http

import (
	"github.com/ZhanibekTau/go-jm-core/pkg/exception"
	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, exception *exception.AppException) {
	c.Set("exception", exception)
	c.Status(exception.Code)
}

func Response(c *gin.Context, status int, data any) {
	c.Set("data", data)
}
