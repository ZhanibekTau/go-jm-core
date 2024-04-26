package exception

import (
	"github.com/go-errors/errors"
	"net/http"
	"runtime"
)

type AppException struct {
	Code    int
	Error   error
	Context string
}

func NewAppException(code int, err error) *AppException {
	return &AppException{code, err, CaptureStackTrace()}
}

func NewValidationAppException() *AppException {
	return &AppException{http.StatusUnprocessableEntity, errors.New("VALIDATION ERROR"), CaptureStackTrace()}
}

func CaptureStackTrace() string {
	buf := make([]byte, 1<<16)
	n := runtime.Stack(buf, false)

	return string(buf[:n])
}
