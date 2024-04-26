package jMiddleware

import (
	"fmt"
	"github.com/exgamer/go-rest-sdk/pkg/config/structures"
	"github.com/exgamer/go-rest-sdk/pkg/exception"
	"github.com/exgamer/go-rest-sdk/pkg/modules/j/jConstants"
	"github.com/exgamer/go-rest-sdk/pkg/modules/j/jLog"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func ResponseHandler(config *structures.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			sentry.CaptureException(err)
			logError(err.Error(), c, config)
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "data": err.Error()})

			return
		}

		appExceptionObject, exists := c.Get("exception")
		fmt.Printf("%+v\n", appExceptionObject)

		if !exists {
			data, _ := c.Get("data")
			logInfo("", c, config)
			c.JSON(http.StatusOK, gin.H{"success": true, "data": data})

			return
		}

		appException := exception.AppException{}
		mapstructure.Decode(appExceptionObject, &appException)
		sentry.CaptureException(appException.Error)
		fmt.Printf("%+v\n", appException)
		logError(appException.Error.Error(), c, config)
		c.JSON(appException.Code, gin.H{"success": false, "message": appException.Error.Error(), "details": appException.Context})
	}
}

func logInfo(message string, c *gin.Context, config *structures.AppConfig) {
	logResponse("INFO", message, c, config)
}

func logError(message string, c *gin.Context, config *structures.AppConfig) {
	logResponse("ERROR", message, c, config)
}

func logResponse(level string, message string, c *gin.Context, config *structures.AppConfig) {
	jLog.PrintJLog(level, config.Name, c.Request.Method, c.Request.RequestURI, c.Writer.Status(), c.GetHeader(jConstants.RequestIdHeaderName), message)
}
