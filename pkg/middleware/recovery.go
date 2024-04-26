package middleware

import (
	"github.com/ZhanibekTau/go-jm-core/pkg/exception"
	httpResponse "github.com/ZhanibekTau/go-jm-core/pkg/http"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

func Recovery() gin.HandlerFunc {
	return RecoveryWithWriter(gin.DefaultErrorWriter)
}

func RecoveryWithWriter(out io.Writer) gin.HandlerFunc {
	var logger *log.Logger

	if out != nil {
		logger = log.New(out, "\n\n\x1b[31m", log.LstdFlags)
	}

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if logger != nil {
					httprequest, _ := httputil.DumpRequest(c.Request, false)
					goErr := errors.Wrap(err, 3)
					reset := string([]byte{27, 91, 48, 109})
					logger.Printf("[Recovery] panic recovered:\n\n%s%s\n\n%s%s", httprequest, goErr.Error(), goErr.Stack(), reset)
				}

				sentry.CaptureException(errors.New(err))
				httpResponse.Error(c, exception.NewAppException(http.StatusInternalServerError, errors.New(err)))
			}
		}()
		c.Next()
	}
}
