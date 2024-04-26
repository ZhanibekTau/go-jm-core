package gin

import (
	"fmt"
	"github.com/ZhanibekTau/go-jm-core/pkg/config/structures"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	timeout "github.com/vearne/gin-timeout"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"net/http"
	"time"
)

func InitRouter(appConfig *structures.AppConfig) *gin.Engine {
	if appConfig.AppEnv == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "404 page not found"})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"code": "METHOD_NOT_ALLOWED", "message": "405 method not allowed"})
	})

	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: appConfig.SentryDsn,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	router.Use(sentrygin.New(sentrygin.Options{}))

	router.Use(gin.Logger())
	router.Use(timeout.Timeout(timeout.WithTimeout(time.Duration(appConfig.HandlerTimeout) * time.Second)))
	router.Use(gin.CustomRecovery(ErrorHandler))

	return router
}

func ErrorHandler(c *gin.Context, err any) {
	goErr := errors.Wrap(err, 2)

	details := make([]string, len(goErr.StackFrames()))

	for _, frame := range goErr.StackFrames() {
		details = append(details, frame.String())
	}

	sentry.CaptureException(goErr)

	c.JSON(http.StatusInternalServerError, gin.H{"success": false, "data": goErr.Error(), "details": details})
}
